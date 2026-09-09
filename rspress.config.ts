import * as path from 'node:path';
import { defineConfig } from '@rspress/core';
import { pluginGiscus } from 'rspress-plugin-giscus';

export default defineConfig({
  root: path.join(__dirname, 'docs'),
  lang: 'zh',
  plugins: [
    pluginGiscus({
      // 评论区载体：difyz9/ytb2bili-docs（已在 GitHub 开启 Discussions）
      // categoryId 取自 giscus.app 生成的配置；theme 如需暗色适配可改为 dark_tritanopia
      repo: 'difyz9/ytb2bili-docs',
      repoId: 'R_kgDORQaFxg',
      category: 'General',
      categoryId: 'DIC_kwDORQaFxs4DFMMa',
      theme: 'light_tritanopia',
    }),
  ],
  title: 'ytb2bili-go',
  description: 'YouTube → Bilibili 视频搬运流水线工具（搜索、下载、转录、翻译、配音、投稿、字幕上传）',
  themeConfig: {
    socialLinks: [
      {
        icon: 'github',
        mode: 'link',
        content: 'https://github.com/difyz9/ytb2bili-cli',
      },
    ],
  },
});
