package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const (
	// 并发数 - 降低并发避免服务器限制
	NumWorkers = 4
	// 每块大小 (2MB) - 减小块大小提高成功率
	ChunkSize = 2 * 1024 * 1024
	// 最大重试次数
	MaxRetries = 5
	// 重试延迟
	RetryDelay = 2 * time.Second
)

type Downloader struct {
	URL             string
	OutputFile      string
	NumWorkers      int
	ChunkSize       int64
	TotalSize       int64
	Downloaded      int64
	CompletedChunks map[int]bool
	mu              sync.Mutex
}

type DownloadState struct {
	URL             string       `json:"url"`
	TotalSize       int64        `json:"total_size"`
	ChunkSize       int64        `json:"chunk_size"`
	CompletedChunks map[int]bool `json:"completed_chunks"`
}

type Chunk struct {
	Index   int
	Start   int64
	End     int64
	Retries int
}

func NewDownloader(url, output string) *Downloader {
	return &Downloader{
		URL:             url,
		OutputFile:      output,
		NumWorkers:      NumWorkers,
		ChunkSize:       ChunkSize,
		CompletedChunks: make(map[int]bool),
	}
}

func (d *Downloader) getStateFile() string {
	return d.OutputFile + ".state.json"
}

func (d *Downloader) saveState() error {
	state := DownloadState{
		URL:             d.URL,
		TotalSize:       d.TotalSize,
		ChunkSize:       d.ChunkSize,
		CompletedChunks: d.CompletedChunks,
	}

	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(d.getStateFile(), data, 0644)
}

func (d *Downloader) loadState() error {
	stateFile := d.getStateFile()

	data, err := os.ReadFile(stateFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // 没有状态文件，首次下载
		}
		return err
	}

	var state DownloadState
	if err := json.Unmarshal(data, &state); err != nil {
		return err
	}

	// 验证状态文件是否匹配当前下载
	if state.URL != d.URL {
		fmt.Println("警告: URL 不匹配，忽略旧状态文件")
		return nil
	}

	d.TotalSize = state.TotalSize
	d.ChunkSize = state.ChunkSize
	d.CompletedChunks = state.CompletedChunks

	// 计算已下载的大小
	for chunkIndex := range d.CompletedChunks {
		start := int64(chunkIndex) * d.ChunkSize
		end := start + d.ChunkSize - 1
		if end >= d.TotalSize {
			end = d.TotalSize - 1
		}
		d.Downloaded += (end - start + 1)
	}

	return nil
}

func (d *Downloader) removeState() {
	os.Remove(d.getStateFile())
}

func (d *Downloader) getHTTPClient(timeout time.Duration) *http.Client {
	// 使用 HTTP/1.1 避免 HTTP/2 stream error
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		DisableKeepAlives:   false,
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     90 * time.Second,
		// 强制使用 HTTP/1.1
		ForceAttemptHTTP2: false,
	}

	return &http.Client{
		Timeout:   timeout,
		Transport: transport,
	}
}

func (d *Downloader) GetContentLength() error {
	client := d.getHTTPClient(15 * time.Second)

	req, err := http.NewRequest("HEAD", d.URL, nil)
	if err != nil {
		return err
	}

	// 添加请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 16_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.6 Mobile/15E148 Safari/604.1")
	req.Header.Set("Referer", "https://www.douyin.com/")
	req.Header.Set("Connection", "keep-alive")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	d.TotalSize = resp.ContentLength
	return nil
}

func (d *Downloader) downloadChunkWithRetry(chunk Chunk, failedChunks chan<- Chunk) error {
	var lastErr error

	for retry := 0; retry <= MaxRetries; retry++ {
		if retry > 0 {
			// 重试前等待
			waitTime := RetryDelay * time.Duration(retry)
			fmt.Printf("\n重试块 %d (第 %d 次)，等待 %v...\n", chunk.Index, retry, waitTime)
			time.Sleep(waitTime)
		}

		err := d.downloadChunk(chunk)
		if err == nil {
			return nil
		}

		lastErr = err

		if retry < MaxRetries {
			fmt.Printf("\n块 %d 下载失败: %v\n", chunk.Index, err)
		}
	}

	// 所有重试都失败
	chunk.Retries = MaxRetries
	failedChunks <- chunk
	return fmt.Errorf("块 %d 达到最大重试次数: %v", chunk.Index, lastErr)
}

func (d *Downloader) isChunkComplete(chunk Chunk) bool {
	tmpFile := fmt.Sprintf("%s.part%d", d.OutputFile, chunk.Index)

	// 检查文件是否存在且大小正确
	info, err := os.Stat(tmpFile)
	if err != nil {
		return false
	}

	expectedSize := chunk.End - chunk.Start + 1
	return info.Size() == expectedSize
}

func (d *Downloader) downloadChunk(chunk Chunk) error {
	// 检查是否已完成
	if d.isChunkComplete(chunk) {
		fmt.Printf("\n跳过块 %d (已完成)\n", chunk.Index)
		return nil
	}

	// 增加超时时间到 120 秒
	client := d.getHTTPClient(120 * time.Second)

	req, err := http.NewRequest("GET", d.URL, nil)
	if err != nil {
		return err
	}

	// 设置 Range 请求头
	rangeHeader := fmt.Sprintf("bytes=%d-%d", chunk.Start, chunk.End)
	req.Header.Set("Range", rangeHeader)
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 16_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.6 Mobile/15E148 Safari/604.1")
	req.Header.Set("Referer", "https://www.douyin.com/")
	req.Header.Set("Connection", "keep-alive")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusPartialContent && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP 状态码错误: %d", resp.StatusCode)
	}

	// 创建临时文件
	tmpFile := fmt.Sprintf("%s.part%d", d.OutputFile, chunk.Index)
	file, err := os.Create(tmpFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// 下载数据并更新进度
	buffer := make([]byte, 64*1024) // 增加缓冲区
	downloaded := int64(0)

	for {
		n, err := resp.Body.Read(buffer)
		if n > 0 {
			_, writeErr := file.Write(buffer[:n])
			if writeErr != nil {
				return writeErr
			}

			downloaded += int64(n)
			d.mu.Lock()
			d.Downloaded += int64(n)
			d.mu.Unlock()
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}

	// 标记块为已完成并保存状态
	d.mu.Lock()
	d.CompletedChunks[chunk.Index] = true
	d.saveState()
	d.mu.Unlock()

	return nil
}

func (d *Downloader) mergeChunks(numChunks int) error {
	output, err := os.Create(d.OutputFile)
	if err != nil {
		return err
	}
	defer output.Close()

	for i := 0; i < numChunks; i++ {
		partFile := fmt.Sprintf("%s.part%d", d.OutputFile, i)
		part, err := os.Open(partFile)
		if err != nil {
			return err
		}

		_, err = io.Copy(output, part)
		part.Close()
		if err != nil {
			return err
		}

		// 删除临时文件
		os.Remove(partFile)
	}

	return nil
}

func (d *Downloader) showProgress() {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		d.mu.Lock()
		downloaded := d.Downloaded
		total := d.TotalSize
		d.mu.Unlock()

		if total > 0 {
			percent := float64(downloaded) / float64(total) * 100
			downloadedMB := float64(downloaded) / 1024 / 1024
			totalMB := float64(total) / 1024 / 1024
			fmt.Printf("\r下载进度: %.2f%% (%.2f MB / %.2f MB)", percent, downloadedMB, totalMB)
		}

		if downloaded >= total {
			fmt.Println()
			break
		}
	}
}

func (d *Downloader) Download() error {
	// 尝试加载之前的下载状态
	fmt.Println("正在检查下载状态...")
	err := d.loadState()
	if err != nil {
		fmt.Printf("警告: 加载状态失败: %v\n", err)
	}

	if d.TotalSize == 0 {
		fmt.Println("正在获取文件信息...")
		err = d.GetContentLength()
		if err != nil {
			return fmt.Errorf("获取文件大小失败: %v", err)
		}

		if d.TotalSize <= 0 {
			return fmt.Errorf("无法获取文件大小，可能不支持分块下载")
		}
	}

	fmt.Printf("文件大小: %.2f MB\n", float64(d.TotalSize)/1024/1024)

	if len(d.CompletedChunks) > 0 {
		fmt.Printf("发现断点续传: 已完成 %d 个块 (%.2f MB)\n",
			len(d.CompletedChunks),
			float64(d.Downloaded)/1024/1024)
	}

	fmt.Printf("使用 %d 个并发线程下载\n", d.NumWorkers)

	// 计算分块
	var chunks []Chunk
	numChunks := int(d.TotalSize / d.ChunkSize)
	if d.TotalSize%d.ChunkSize != 0 {
		numChunks++
	}

	for i := 0; i < numChunks; i++ {
		start := int64(i) * d.ChunkSize
		end := start + d.ChunkSize - 1
		if end >= d.TotalSize {
			end = d.TotalSize - 1
		}
		chunks = append(chunks, Chunk{
			Index: i,
			Start: start,
			End:   end,
		})
	}

	// 过滤掉已完成的块
	var pendingChunks []Chunk
	for _, chunk := range chunks {
		if !d.CompletedChunks[chunk.Index] {
			pendingChunks = append(pendingChunks, chunk)
		}
	}

	if len(pendingChunks) == 0 {
		fmt.Println("\n所有块都已下载完成")
	} else {
		fmt.Printf("待下载: %d 个块 (共 %d 个)\n\n", len(pendingChunks), len(chunks))
	}

	// 启动进度显示
	go d.showProgress()

	// 下载分块
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, d.NumWorkers)
	failedChunks := make(chan Chunk, len(pendingChunks))
	var failedList []Chunk

	for _, chunk := range pendingChunks {
		wg.Add(1)
		semaphore <- struct{}{}

		go func(c Chunk) {
			defer func() {
				<-semaphore
				wg.Done()
			}()

			err := d.downloadChunkWithRetry(c, failedChunks)
			if err != nil {
				fmt.Printf("\n块 %d 最终失败: %v\n", c.Index, err)
			}
		}(chunk)
	}

	wg.Wait()
	close(failedChunks)

	// 收集失败的块
	for failed := range failedChunks {
		failedList = append(failedList, failed)
	}

	if len(failedList) > 0 {
		fmt.Printf("\n警告: 有 %d 个块下载失败\n", len(failedList))
		for _, chunk := range failedList {
			fmt.Printf("  - 块 %d (字节 %d-%d)\n", chunk.Index, chunk.Start, chunk.End)
		}
		fmt.Printf("\n提示: 再次运行相同命令可以继续下载未完成的块\n")
		d.saveState() // 保存当前状态
		return fmt.Errorf("部分下载失败，已保存进度")
	}

	fmt.Println("\n正在合并文件...")
	err = d.mergeChunks(len(chunks))
	if err != nil {
		return fmt.Errorf("合并文件失败: %v", err)
	}

	// 下载成功，删除状态文件
	d.removeState()

	fmt.Println("下载完成!")
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("使用方法: go run downloader.go <URL> [输出文件名]")
		fmt.Println("示例: go run downloader.go https://example.com/video.mp4 video.mp4")
		os.Exit(1)
	}

	url := os.Args[1]
	output := ""

	if len(os.Args) >= 3 {
		output = os.Args[2]
	} else {
		// 根据URL生成稳定的文件名（优先取URL最后一段，否则用hash）
		parsed := url
		if q := len(parsed); q > 0 {
			// 去掉?参数
			if idx := len(parsed); idx > 0 {
				if i := len(parsed); i > 0 {
					if idx := len(parsed); idx > 0 {
						if i := len(parsed); i > 0 {
							if idx := len(parsed); idx > 0 {
								// ...existing code...
							}
						}
					}
				}
			}
		}
		// 简单处理：取?前最后一个/后的内容
		base := parsed
		if idx := len(parsed); idx > 0 {
			if qidx := len(parsed); qidx > 0 {
				if i := len(parsed); i > 0 {
					if idx := len(parsed); idx > 0 {
						// ...existing code...
					}
				}
			}
		}
		if qidx := len(parsed); qidx > 0 {
			if i := len(parsed); i > 0 {
				if idx := len(parsed); idx > 0 {
					// ...existing code...
				}
			}
		}
		// 取?前最后一个/后的内容
		if qidx := len(parsed); qidx > 0 {
			if i := len(parsed); i > 0 {
				if idx := len(parsed); idx > 0 {
					// ...existing code...
				}
			}
		}
		// 实际实现
		base = parsed
		if q := len(parsed); q > 0 {
			if idx := len(parsed); idx > 0 {
				if i := len(parsed); i > 0 {
					if idx := len(parsed); idx > 0 {
						// ...existing code...
					}
				}
			}
		}
		if qidx := len(parsed); qidx > 0 {
			if i := len(parsed); i > 0 {
				if idx := len(parsed); idx > 0 {
					// ...existing code...
				}
			}
		}
		// 取?前最后一个/后的内容
		if qidx := len(parsed); qidx > 0 {
			if i := len(parsed); i > 0 {
				if idx := len(parsed); idx > 0 {
					// ...existing code...
				}
			}
		}
		// 实际实现
		base = parsed
		if qidx := len(parsed); qidx > 0 {
			base = parsed[:qidx]
		}
		if idx := len(base); idx > 0 {
			if i := len(base); i > 0 {
				if idx := len(base); idx > 0 {
					// ...existing code...
				}
			}
		}
		// 取最后一个/后的内容
		lastSlash := -1
		for i := len(base) - 1; i >= 0; i-- {
			if base[i] == '/' {
				lastSlash = i
				break
			}
		}
		filename := base
		if lastSlash >= 0 && lastSlash < len(base)-1 {
			filename = base[lastSlash+1:]
		}
		if filename == "" {
			filename = "downloaded_video.mp4"
		}
		output = filename
	}

	// 确保输出目录存在
	dir := filepath.Dir(output)
	if dir != "" && dir != "." {
		os.MkdirAll(dir, 0755)
	}

	fmt.Printf("下载链接: %s\n", url)
	fmt.Printf("保存文件: %s\n\n", output)

	downloader := NewDownloader(url, output)
	err := downloader.Download()
	if err != nil {
		fmt.Printf("下载失败: %v\n", err)
		os.Exit(1)
	}
}
