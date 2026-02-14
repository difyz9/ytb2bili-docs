根据分享链接获取单个作品数据/Get single video data by sharing link


data.aweme_detail.video.download_addr.url_list


---

{
    "code": 200,
    "request_id": "c9813f75-3cf5-4180-84ac-f20867e91986",
    "message": "Request successful. This request will incur a charge.",
    "message_zh": "请求成功，本次请求将被计费。",
    "support": "Discord: https://discord.gg/aMEAS8Xsvz",
    "time": "2026-01-29 05:26:32",
    "time_stamp": 1769693192,
    "time_zone": "America/Los_Angeles",
    "docs": "https://api.tikhub.io/#/Douyin-App-V3-API/fetch_one_video_by_share_url_api_v1_douyin_app_v3_fetch_one_video_by_share_url_get",
    "cache_message": "This request will be cached. You can access the cached result directly using the URL below, valid for 24 hours. Accessing the cache will not incur additional charges.",
    "cache_message_zh": "本次请求将被缓存，你可以使用下面的 URL 直接访问缓存结果，有效期为 24 小时，访问缓存不会产生额外费用。",
    "cache_url": "https://cache.tikhub.io/api/v1/cache/public/c9813f75-3cf5-4180-84ac-f20867e91986?sign=04b326be4c19804df4d1f5ac4778a9c36cdcf75340aada0dcc48b50f82519fbc",
    "router": "/api/v1/douyin/app/v3/fetch_one_video_by_share_url",
    "params": {
        "share_url": " https://v.douyin.com/p7-T74NL4h4/"
    },
    "data": {
        "aweme_detail": {
            "entertainment_video_type": 3,
            "is_use_music": false,
            "item_aigc_follow_shot": 1,
            "collect_stat": 0,
            "dislike_dimension_list_v2": [
                {
                    "icon": "https://p3.douyinpic.com/aweme-server-static-resource/dislike_objective.png~tplv-obj.image",
                    "text": "我不想看",
                    "entitys": [
                        {
                            "dimension_id": 1,
                            "pre_text": "作者",
                            "sub_text": "一发没",
                            "server_extra": "{\"author_id\":2543612830166135}",
                            "en_name": "author"
                        },
                        {
                            "pre_text": "音乐",
                            "sub_text": "@一发没创作的原声",
                            "server_extra": "{\"music_id\":7600676840910375726}",
                            "en_name": "music",
                            "dimension_id": 2
                        }
                    ]
                },
                {
                    "icon": "https://p3.douyinpic.com/aweme-server-static-resource/dislike_subjective.png~tplv-obj.image",
                    "text": "反馈内容问题",
                    "entitys": [
                        {
                            "dimension_id": 12,
                            "pre_text": "相似过多",
                            "server_extra": "{\"subjective_id\":1001}",
                            "en_name": "repeat"
                        },
                        {
                            "server_extra": "{\"subjective_id\":1002}",
                            "en_name": "ad",
                            "dimension_id": 12,
                            "pre_text": "广告营销"
                        },
                        {
                            "dimension_id": 12,
                            "pre_text": "恶心恐怖",
                            "server_extra": "{\"subjective_id\":1003}",
                            "en_name": "disgust"
                        },
                        {
                            "pre_text": "色情低俗",
                            "server_extra": "{\"subjective_id\":1004}",
                            "en_name": "porn",
                            "dimension_id": 12
                        },
                        {
                            "en_name": "fake",
                            "dimension_id": 12,
                            "pre_text": "虚假夸张",
                            "server_extra": "{\"subjective_id\":1005}"
                        }
                    ]
                }
            ],
            "prevent_download": false,
            "douplus_user_type": 0,
            "activity_video_type": -1,
            "relation_labels": null,
            "story_ttl": 0,
            "aweme_listen_struct": {
                "trace_info": "{\"copyright_not_speech\":\"false\",\"copyright_reason\":\"not_intercept\",\"copyright_tag_hit\":\"0\",\"copyright_use_aed_default\":\"false\",\"copyright_use_tag_default\":\"false\",\"cp_ab\":\"true\",\"desc\":\"\",\"duration_over\":\"true\",\"hit_high_risk\":\"true\",\"media_type\":\"4\",\"reason\":\"hit_ab_0\",\"show\":\"false\"}"
            },
            "enable_comment_sticker_rec": false,
            "boost_status": 0,
            "geofencing_regions": null,
            "trends_event_track": "{}",
            "caption": "",
            "friend_recommend_info": {
                "friend_recommend_source": 26,
                "disable_friend_recommend_guide_label": false
            },
            "comment_gid": 7600676564674478720,
            "douyin_pc_video_extra_seo": "",
            "show_follow_button": {},
            "danmaku_control": {
                "post_privilege_level": 0,
                "pass_through_params": "{\"has_danmaku\":1}",
                "smart_mode_decision": 0,
                "enable_danmaku": true,
                "is_post_denied": false,
                "skip_danmaku": false,
                "last_danmaku_offset": 603661,
                "post_denied_reason": "",
                "activities": [
                    {
                        "id": 1224,
                        "type": 1
                    }
                ],
                "first_danmaku_offset": 17605,
                "danmaku_cnt": 0
            },
            "text_extra": [
                {
                    "end": 0,
                    "type": 5,
                    "search_text": "疯狂动物城2",
                    "start": 0,
                    "search_query_id": "6595537046654162180"
                }
            ],
            "component_info_v2": "{\"desc_lines_limit\":0,\"hide_marquee\":false}",
            "without_watermark": false,
            "entertainment_product_info": {
                "market_info": {
                    "limit_free": {
                        "in_free": false
                    }
                }
            },
            "rate": 12,
            "packed_clips": null,
            "guide_scene_info": {},
            "ent_log_extra": {
                "log_extra": "{\"global_log_extra\":null,\"page_log_extra\":null,\"aweme_log_extra\":{\"ce_current_group_id\":\"7600676564674478720\",\"author_id\":\"2543612830166135\"},\"extra_log_extra\":null}"
            },
            "item_warn_notification": {
                "content": "",
                "show": false,
                "type": 0
            },
            "create_time": 1769683575,
            "geofencing": [],
            "is_fantasy": false,
            "uniqid_position": null,
            "disable_relation_bar": 0,
            "is_share_post": false,
            "aweme_type_tags": "",
            "distribute_circle": {
                "distribute_type": 0,
                "is_campus": false,
                "campus_block_interaction": false
            },
            "is_image_beat": false,
            "is_moment_story": 0,
            "is_collects_selected": 0,
            "aweme_id": "7600676564674478720",
            "origin_comment_ids": null,
            "libfinsert_task_id": "",
            "component_control": {},
            "vr_type": 0,
            "interaction_stickers": null,
            "preview_title": "疯狂动物城2（国语） ",
            "xigua_base_info": {
                "status": 0,
                "star_altar_order_id": 0,
                "star_altar_type": 0,
                "item_id": 0
            },
            "item_share": 0,
            "author": {
                "download_setting": -1,
                "new_story_cover": null,
                "region": "CN",
                "sec_uid": "MS4wLjABAAAAPwTniQa_UZ9Z3ICiHGPnVsruQ_MUJHdAtTrCelTIS0NV_IPTM4Qf8_uyc8wqXAMX",
                "neiguang_shield": 0,
                "is_binded_weibo": false,
                "enterprise_verify_reason": "",
                "share_qrcode_uri": "",
                "constellation": 0,
                "hide_search": false,
                "awemehts_greet_info": "",
                "youtube_expire_time": 0,
                "favoriting_count": 0,
                "interest_tags": null,
                "is_blocked_v2": false,
                "room_id": 0,
                "shield_comment_notice": 0,
                "school_name": "",
                "data_label_list": null,
                "follow_status": 0,
                "special_follow_status": 0,
                "custom_verify": "",
                "card_entries": null,
                "location": "",
                "total_favorited": 2948,
                "weibo_schema": "",
                "card_entries_not_display": null,
                "cha_list": null,
                "language": "zh-Hans",
                "user_canceled": false,
                "duet_setting": 0,
                "weibo_name": "",
                "follower_list_secondary_information_struct": null,
                "is_star": false,
                "special_lock": 1,
                "is_gov_media_vip": false,
                "story_interactive": 0,
                "geofencing": [],
                "cv_level": "",
                "has_facebook_token": false,
                "apple_account": 0,
                "birthday": "",
                "unique_id_modify_time": 1769693192,
                "aweme_hotsoon_auth": 1,
                "follower_count": 0,
                "max_follower_count": 0,
                "youtube_channel_id": "",
                "user_permissions": null,
                "signature": "",
                "google_account": "",
                "show_image_bubble": false,
                "personal_tag_list": null,
                "ad_cover_url": null,
                "can_set_geofencing": null,
                "user_not_show": 1,
                "ins_id": "",
                "mate_add_permission": 0,
                "account_region": "",
                "school_id": "",
                "fb_expire_time": 0,
                "is_verified": true,
                "authority_status": 0,
                "avatar_larger": {
                    "uri": "1080x1080/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa",
                    "width": 720,
                    "url_list": [
                        "https://p3.douyinpic.com/aweme/1080x1080/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.webp?from=327834062",
                        "https://p9.douyinpic.com/aweme/1080x1080/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.webp?from=327834062",
                        "https://p6.douyinpic.com/aweme/1080x1080/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.webp?from=327834062",
                        "https://p3.douyinpic.com/aweme/1080x1080/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.jpeg?from=327834062"
                    ],
                    "height": 720
                },
                "live_agreement_time": 0,
                "is_cf": 0,
                "platform_sync_info": null,
                "ban_user_functions": [],
                "comment_setting": 0,
                "item_list": null,
                "shield_follow_notice": 0,
                "story_ttl": 0,
                "need_recommend": 0,
                "user_mode": 0,
                "accept_private_policy": false,
                "school_poi_id": "",
                "tw_expire_time": 0,
                "live_commerce": false,
                "avatar_thumb": {
                    "uri": "100x100/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa",
                    "url_list": [
                        "https://p3.douyinpic.com/aweme/100x100/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.webp?from=327834062",
                        "https://p6.douyinpic.com/aweme/100x100/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.webp?from=327834062",
                        "https://p9.douyinpic.com/aweme/100x100/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.webp?from=327834062",
                        "https://p3.douyinpic.com/aweme/100x100/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.jpeg?from=327834062"
                    ],
                    "height": 720,
                    "width": 720
                },
                "is_discipline_member": false,
                "text_extra": null,
                "signature_display_lines": 5,
                "avatar_medium": {
                    "uri": "720x720/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa",
                    "url_list": [
                        "https://p3.douyinpic.com/aweme/720x720/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.webp?from=327834062",
                        "https://p9.douyinpic.com/aweme/720x720/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.webp?from=327834062",
                        "https://p6.douyinpic.com/aweme/720x720/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.webp?from=327834062",
                        "https://p3.douyinpic.com/aweme/720x720/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.jpeg?from=327834062"
                    ],
                    "height": 720,
                    "width": 720
                },
                "school_type": 0,
                "followers_detail": null,
                "ky_only_predict": 0,
                "close_friend_type": 0,
                "user_not_see": 0,
                "avatar_300x300": {
                    "height": 720,
                    "uri": "300x300/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa",
                    "url_list": [
                        "https://p3.douyinpic.com/img/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa~c5_300x300.webp?from=327834062",
                        "https://p6.douyinpic.com/img/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa~c5_300x300.webp?from=327834062",
                        "https://p9.douyinpic.com/img/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa~c5_300x300.webp?from=327834062",
                        "https://p3.douyinpic.com/img/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa~c5_300x300.jpeg?from=327834062"
                    ],
                    "width": 720
                },
                "user_rate": 1,
                "aweme_control": {
                    "can_share": true,
                    "can_comment": true,
                    "can_show_comment": true,
                    "can_forward": true
                },
                "signature_extra": null,
                "offline_info_list": null,
                "is_phone_binded": false,
                "status": 1,
                "special_people_labels": null,
                "shield_digg_notice": 0,
                "youtube_channel_title": "",
                "following_count": 0,
                "user_tags": null,
                "uid": "2543612830166135",
                "is_block": false,
                "reflow_page_gid": 0,
                "card_sort_priority": null,
                "sync_to_toutiao": 0,
                "gender": 0,
                "avatar_168x168": {
                    "height": 720,
                    "width": 720,
                    "uri": "168x168/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa",
                    "url_list": [
                        "https://p3.douyinpic.com/img/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa~c5_168x168.webp?from=327834062",
                        "https://p6.douyinpic.com/img/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa~c5_168x168.webp?from=327834062",
                        "https://p9.douyinpic.com/img/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa~c5_168x168.webp?from=327834062",
                        "https://p3.douyinpic.com/img/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa~c5_168x168.jpeg?from=327834062"
                    ]
                },
                "relative_users": null,
                "cover_url": [
                    {
                        "url_list": [
                            "https://p3-sign.douyinpic.com/obj/c8510002be9a3a61aad2?lk3s=138a59ce&x-expires=1770901200&x-signature=QchLii%2BUD8XvyHzPDW6ei9vI86s%3D&from=327834062"
                        ],
                        "height": 720,
                        "width": 720,
                        "uri": "c8510002be9a3a61aad2"
                    }
                ],
                "show_nearby_active": false,
                "hide_location": false,
                "endorsement_info_list": null,
                "short_id": "97681640691",
                "follower_request_status": 0,
                "risk_notice_text": "",
                "is_blocking_v2": false,
                "twitter_id": "",
                "live_verify": 0,
                "type_label": null,
                "unique_id": "97681640691",
                "with_dou_entry": false,
                "prevent_download": false,
                "display_info": null,
                "is_ad_fake": false,
                "with_fusion_shop_entry": false,
                "follower_status": 0,
                "with_shop_entry": false,
                "contacts_status": 2,
                "has_twitter_token": false,
                "need_points": null,
                "search_impr": {
                    "entity_id": "2543612830166135"
                },
                "is_not_show": false,
                "with_commerce_entry": false,
                "has_unread_story": false,
                "has_insights": false,
                "enable_nearby_visible": true,
                "weibo_verify": "",
                "link_item_list": null,
                "share_info": {
                    "share_desc": "",
                    "share_title_myself": "",
                    "share_qrcode_url": {
                        "width": 720,
                        "url_list": [],
                        "height": 720,
                        "uri": ""
                    },
                    "share_desc_info": "",
                    "share_url": "",
                    "share_weibo_desc": "",
                    "share_title": "",
                    "share_title_other": ""
                },
                "avatar_uri": "aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa",
                "commerce_user_level": 0,
                "nickname": "一发没",
                "story_open": false,
                "im_role_ids": null,
                "homepage_bottom_toast": null,
                "is_mix_user": false,
                "reflow_page_uid": 0,
                "school_category": 0,
                "stitch_setting": 0,
                "white_cover_url": null,
                "story25_comment": 0,
                "video_icon": {
                    "uri": "",
                    "url_list": [],
                    "height": 720,
                    "width": 720
                },
                "live_status": 0,
                "download_prompt_ts": 0,
                "user_period": 0,
                "twitter_name": "",
                "has_youtube_token": false,
                "verification_type": 1,
                "user_age": -1,
                "create_time": 0,
                "aweme_hotsoon_auth_relation": 1,
                "contrail_list": null,
                "bind_phone": "",
                "live_agreement": 0,
                "has_orders": false,
                "secret": 0,
                "weibo_url": "",
                "verify_info": "",
                "cf_list": null,
                "has_email": false,
                "live_high_value": 0,
                "story_count": 0,
                "aweme_count": 2,
                "react_setting": 0,
                "comment_filter_status": 0
            },
            "image_infos": null,
            "is_top": 0,
            "region": "",
            "group_id": "7591537633856064795",
            "music": {
                "avatar_thumb": {
                    "width": 720,
                    "uri": "100x100/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa",
                    "url_list": [
                        "https://p3.douyinpic.com/aweme/100x100/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.webp?from=327834062",
                        "https://p6.douyinpic.com/aweme/100x100/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.webp?from=327834062",
                        "https://p9.douyinpic.com/aweme/100x100/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.webp?from=327834062",
                        "https://p3.douyinpic.com/aweme/100x100/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.jpeg?from=327834062"
                    ],
                    "height": 720
                },
                "music_chart_ranks": null,
                "artists": [],
                "offline_desc": "",
                "is_audio_url_with_cookie": false,
                "owner_nickname": "一发没",
                "preview_start_time": 0,
                "is_del_video": false,
                "avatar_large": {
                    "width": 720,
                    "uri": "1080x1080/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa",
                    "url_list": [
                        "https://p3.douyinpic.com/aweme/1080x1080/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.webp?from=327834062",
                        "https://p6.douyinpic.com/aweme/1080x1080/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.webp?from=327834062",
                        "https://p9.douyinpic.com/aweme/1080x1080/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.webp?from=327834062",
                        "https://p3.douyinpic.com/aweme/1080x1080/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.jpeg?from=327834062"
                    ],
                    "height": 720
                },
                "is_original_sound": true,
                "musician_user_infos": null,
                "position": null,
                "user_count": 0,
                "is_matched_metadata": false,
                "pgc_music_type": 2,
                "redirect": false,
                "mute_share": false,
                "collect_stat": 0,
                "author_deleted": false,
                "prevent_download": false,
                "author_position": null,
                "album": "",
                "source_platform": 23,
                "search_impr": {
                    "entity_id": "7600676840910375726"
                },
                "shoot_duration": 6451,
                "prevent_item_download_status": 0,
                "owner_handle": "97681640691",
                "status": 1,
                "music_status": 1,
                "show_origin_clip": false,
                "music_collect_count": 0,
                "cover_medium": {
                    "height": 720,
                    "width": 720,
                    "uri": "720x720/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa",
                    "url_list": [
                        "https://p3.douyinpic.com/aweme/720x720/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.webp?from=327834062",
                        "https://p6.douyinpic.com/aweme/720x720/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.webp?from=327834062",
                        "https://p9.douyinpic.com/aweme/720x720/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.webp?from=327834062",
                        "https://p3.douyinpic.com/aweme/720x720/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.jpeg?from=327834062"
                    ]
                },
                "cover_large": {
                    "url_list": [
                        "https://p3.douyinpic.com/aweme/1080x1080/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.webp?from=327834062",
                        "https://p6.douyinpic.com/aweme/1080x1080/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.webp?from=327834062",
                        "https://p9.douyinpic.com/aweme/1080x1080/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.webp?from=327834062",
                        "https://p3.douyinpic.com/aweme/1080x1080/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.jpeg?from=327834062"
                    ],
                    "uri": "1080x1080/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa",
                    "height": 720,
                    "width": 720
                },
                "is_restricted": false,
                "can_background_play": true,
                "id_str": "7600676840910375726",
                "reason_type": 0,
                "cover_thumb": {
                    "height": 720,
                    "uri": "100x100/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa",
                    "width": 720,
                    "url_list": [
                        "https://p3.douyinpic.com/aweme/100x100/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.webp?from=327834062",
                        "https://p6.douyinpic.com/aweme/100x100/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.webp?from=327834062",
                        "https://p9.douyinpic.com/aweme/100x100/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.webp?from=327834062",
                        "https://p3.douyinpic.com/aweme/100x100/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.jpeg?from=327834062"
                    ]
                },
                "tag_list": null,
                "id": 7600676840910375726,
                "avatar_medium": {
                    "height": 720,
                    "uri": "720x720/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa",
                    "url_list": [
                        "https://p3.douyinpic.com/aweme/720x720/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.webp?from=327834062",
                        "https://p6.douyinpic.com/aweme/720x720/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.webp?from=327834062",
                        "https://p9.douyinpic.com/aweme/720x720/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.webp?from=327834062",
                        "https://p3.douyinpic.com/aweme/720x720/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.jpeg?from=327834062"
                    ],
                    "width": 720
                },
                "lyric_short_position": null,
                "video_duration": 6451,
                "preview_end_time": 0,
                "is_pgc": false,
                "external_song_info": [],
                "sec_uid": "MS4wLjABAAAAPwTniQa_UZ9Z3ICiHGPnVsruQ_MUJHdAtTrCelTIS0NV_IPTM4Qf8_uyc8wqXAMX",
                "matched_pgc_sound": {
                    "author": "Disney & Shakira",
                    "title": "Zoo - From \"Zootopia 2\"",
                    "mixed_title": "",
                    "mixed_author": "",
                    "cover_medium": {
                        "width": 720,
                        "uri": "tos-cn-v-2774c002/oAiZX6EwyDlPABBlBgfIihCQAOEswAWCBAeoBt",
                        "url_list": [
                            "https://p3.douyinpic.com/aweme/200x200/tos-cn-v-2774c002/oAiZX6EwyDlPABBlBgfIihCQAOEswAWCBAeoBt.jpeg",
                            "https://p6.douyinpic.com/aweme/200x200/tos-cn-v-2774c002/oAiZX6EwyDlPABBlBgfIihCQAOEswAWCBAeoBt.jpeg",
                            "https://p9.douyinpic.com/aweme/200x200/tos-cn-v-2774c002/oAiZX6EwyDlPABBlBgfIihCQAOEswAWCBAeoBt.jpeg"
                        ],
                        "height": 720
                    }
                },
                "music_cover_atmosphere_color_value": "",
                "extra":"{\"aed_music_score\":0.71,\"aed_singing_score\":0.7,\"aggregate_exempt_conf\":[],\"beats\":{},\"cover_colors\":null,\"douyin_beats_info\":{},\"dsp_switch\":0,\"extract_item_id\":7600676564674478720,\"has_edited\":0,\"hit_high_follow_auto\":false,\"hit_high_follow_extend\":false,\"hit_high_follow_original\":false,\"hotsoon_review_time\":-1,\"is_aed_music\":1,\"is_high_follow_user\":false,\"is_red\":0,\"is_subsidy_exp\":false,\"mini_luna\":true,\"music_display_mapping_title\":\"Zoo - From \\\"Zootopia 2\\\" - Disney \& Shakira\",\"music_label_id\":null,\"music_specific_copyright\":true,\"music_tagging\":{\"AEDs\":[\"Female\",\"Vocal\"],\"Genres\":[\"Pop\"],\"Instruments\":null,\"Languages\":[\"English\"],\"Moods\":[\"Dynamic\"],\"SingingVersions\":[\"Seed\"],\"Themes\":[\"Dance\"]},\"review_unshelve_reason\":0,\"reviewed\":0,\"schedule_search_time\":0,\"with_aed_model\":1}",
                "end_time": 0,
                "start_time": 0,
                "is_video_self_see": false,
                "dmv_auto_show": false,
                "author": "一发没",
                "schema_url": "",
                "title": "@一发没创作的原声",
                "mid": "7600676840910375726",
                "is_commerce_music": false,
                "duration": 6451,
                "audition_duration": 6451,
                "author_status": 1,
                "owner_id": "2543612830166135",
                "binded_challenge_id": 0,
                "song": {
                    "id": 7554129057526810640,
                    "id_str": "7554129057526810640",
                    "artists": null
                },
                "dsp_status": 10,
                "play_url": {
                    "url_list": [
                        "https://sf5-hl-cdn-tos.douyinstatic.com/obj/ies-music/7600676858417793830.mp3",
                        "https://sf5-hl-ali-cdn-tos.douyinstatic.com/obj/ies-music/7600676858417793830.mp3"
                    ],
                    "url_key": "7600676840910375726",
                    "height": 720,
                    "width": 720,
                    "uri": "https://sf5-hl-cdn-tos.douyinstatic.com/obj/ies-music/7600676858417793830.mp3"
                },
                "is_original": false,
                "unshelve_countries": null,
                "cover_hd": {
                    "url_list": [
                        "https://p3.douyinpic.com/aweme/1080x1080/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.webp?from=327834062",
                        "https://p6.douyinpic.com/aweme/1080x1080/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.webp?from=327834062",
                        "https://p9.douyinpic.com/aweme/1080x1080/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.webp?from=327834062",
                        "https://p3.douyinpic.com/aweme/1080x1080/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa.jpeg?from=327834062"
                    ],
                    "uri": "1080x1080/aweme-avatar/tos-cn-i-0813c001_oQRwKEgIEqDAcDoFDpCeBpMNwAzA9AAAE428fa",
                    "height": 720,
                    "width": 720
                },
                "artist_user_infos": null
            },
            "need_vs_entry": true,
            "social_tag_list": null,
            "img_bitrate": null,
            "user_digged": 0,
            "original": 0,
            "have_dashboard": false,
            "challenge_position": null,
            "label_top_text": null,
            "is_24_story": 0,
            "is_25_story": 0,
            "origin_text_extra": [],
            "is_hash_tag": 1,
            "is_duet_sing": false,
            "video": {
                "horizontal_type": 1,
                "ratio": "720p",
                "duration": 6451549,
                "format": "mp4",
                "bit_rate": [
                    {
                        "HDR_bit": "",
                        "play_addr": {
                            "height": 720,
                            "width": 1720,
                            "file_cs": "c:0-4976074-c8e6",
                            "data_size": 752145758,
                            "file_hash": "da9f8c238097b54f59364ac5aa30e835",
                            "uri": "v02f52g10003d5tgc9vog65he2bdsmvg",
                            "url_list": [
                                "https://v5-dy-o-abtest.zjcdn.com/e925365269079b1ca870c6ddd8f71e5c/697b874b/video/tos/cn/tos-cn-v-694b40/oQSSebD8ICEENVivfqfAExTF2w7n4ftYKE4TIP/?a=1128&ch=26&cr=13&dr=0&lr=all&cd=0%7C0%7C0%7C&cv=1&br=910&bt=910&cs=2&ds=3&ft=tKyhWzoTd95~~iSYsiI9B4yoFIKrHgSFS4v7p0HCnzXnZeMdxRyqkZ&mime_type=video_mp4&qs=15&rc=aTg4Ojo8ZDdnPDc4ZDw6ZUBpanBnNWs5cjxqODYzNDVpM0AvM18vYTQ2XmMxLTVhYDI2YSNzcF9iMmRzYHFhLS1kMi9zcw%3D%3D&btag=c0010e000b8009&cdn_type=1&cquery=105e_104O_104U_100b_104i&dy_q=1769693192&feature_id=eceb2a8579548afb597413997a0403f2&l=2026012921263252B4EA0EC35D84FA6928&pdp=online_xy&pwid=1&req_cdn_type=r",
                                "https://v5-hl-zenl-ov.zjcdn.com/f4ff7b097f60af41558b095c4d509b45/697b874b/video/tos/cn/tos-cn-v-694b40/oQSSebD8ICEENVivfqfAExTF2w7n4ftYKE4TIP/?a=1128&ch=26&cr=13&dr=0&lr=all&cd=0%7C0%7C0%7C&cv=1&br=910&bt=910&cs=2&ds=3&ft=tKyhWzoTd95~~iSYsiI9B4yoFIKrHgSFS4v7p0HCnzXnZeMdxRyqkZ&mime_type=video_mp4&qs=15&rc=aTg4Ojo8ZDdnPDc4ZDw6ZUBpanBnNWs5cjxqODYzNDVpM0AvM18vYTQ2XmMxLTVhYDI2YSNzcF9iMmRzYHFhLS1kMi9zcw%3D%3D&btag=c0010e000b8009&cdn_type=1&cquery=104i_105e_104O_104U_100b&dy_q=1769693192&feature_id=eceb2a8579548afb597413997a0403f2&l=2026012921263252B4EA0EC35D84FA6928&pdp=online_xy&pwid=1&req_cdn_type=r",
                                "https://api-play.amemv.com/aweme/v1/play/?video_id=v02f52g10003d5tgc9vog65he2bdsmvg&line=0&file_id=9b57c9da850449fda49a79d727b5652f&sign=da9f8c238097b54f59364ac5aa30e835&is_play_url=1&source=PackSourceEnum_AWEME_DETAIL",
                                "https://api.amemv.com/aweme/v1/play/?video_id=v02f52g10003d5tgc9vog65he2bdsmvg&line=1&file_id=9b57c9da850449fda49a79d727b5652f&sign=da9f8c238097b54f59364ac5aa30e835&is_play_url=1&source=PackSourceEnum_AWEME_DETAIL"
                            ],
                            "url_key": "v02f52g10003d5tgc9vog65he2bdsmvg_bytevc1_720p_932669"
                        },
                        "video_extra": "{\"PktOffsetMap\":\"[{\\\"time\\\": 1, \\\"offset\\\": 5232363}, {\\\"time\\\": 2, \\\"offset\\\": 5450403}, {\\\"time\\\": 3, \\\"offset\\\": 5774198}, {\\\"time\\\": 4, \\\"offset\\\": 6115903}, {\\\"time\\\": 5, \\\"offset\\\": 6405128}, {\\\"time\\\": 10, \\\"offset\\\": 7142043}]\",\"format\":\"mp4\",\"definition\":\"720p\",\"quality\":\"adapt_lowest\",\"file_id\":\"9b57c9da850449fda49a79d727b5652f\",\"applog_map\":{\"feature_id\":\"eceb2a8579548afb597413997a0403f2\"},\"audio_metrics\":null,\"audio_score\":null}",
                        "format": "mp4",
                        "gear_name": "comet_bvc1_s1_r1_adapt_lowest_720_1",
                        "bit_rate": 932669,
                        "is_h265": 1,
                        "HDR_type": "",
                        "quality_type": 15,
                        "is_bytevc1": 1,
                        "FPS": 24
                    }
                ],
                "audio": {},
                "bit_rate_audio": null,
                "has_watermark": true,
                "is_source_HDR": 0,
                "cover": {
                    "url_list": [
                        "https://p3-sign.douyinpic.com/tos-cn-i-0813c001/oE7MEC2DaBAjBQAUAYi3Ae0ABATggC0vIVfT7i~tplv-dy-resize-walign-adapt-aq:720:q75.webp?lk3s=138a59ce&x-expires=1770901200&x-signature=Fo3IQL94vxBPto30xP9Q2ozg61s%3D&from=327834062&s=PackSourceEnum_AWEME_DETAIL&se=false&sc=cover&biz_tag=aweme_video&l=2026012921263252B4EA0EC35D84FA6928",
                        "https://p3-sign.douyinpic.com/tos-cn-i-0813c001/oE7MEC2DaBAjBQAUAYi3Ae0ABATggC0vIVfT7i~tplv-dy-resize-walign-adapt-aq:720:q75.jpeg?lk3s=138a59ce&x-expires=1770901200&x-signature=MPt9pWdu%2BlbhhwrLuITok9SIIQc%3D&from=327834062&s=PackSourceEnum_AWEME_DETAIL&se=false&sc=cover&biz_tag=aweme_video&l=2026012921263252B4EA0EC35D84FA6928"
                    ],
                    "uri": "tos-cn-i-0813c001/oE7MEC2DaBAjBQAUAYi3Ae0ABATggC0vIVfT7i",
                    "height": 301,
                    "width": 720
                },
                "use_static_cover": true,
                "origin_cover": {
                    "uri": "tos-cn-p-694b40/okEQtiEIICASn2efSBFI77F8fyY2dqf4DEwTAi",
                    "url_list": [
                        "https://p3-sign.douyinpic.com/tos-cn-p-694b40/okEQtiEIICASn2efSBFI77F8fyY2dqf4DEwTAi~tplv-dy-360p.webp?lk3s=138a59ce&x-expires=1770901200&x-signature=9y9tF3afxkz03IbJPxYbP%2FY8t3w%3D&from=327834062&s=PackSourceEnum_AWEME_DETAIL&se=false&sc=origin_cover&biz_tag=aweme_video&l=2026012921263252B4EA0EC35D84FA6928",
                        "https://p3-sign.douyinpic.com/tos-cn-p-694b40/okEQtiEIICASn2efSBFI77F8fyY2dqf4DEwTAi~tplv-dy-360p.jpeg?lk3s=138a59ce&x-expires=1770901200&x-signature=xhETX49Fe5MGmVU0r1AMRWvaNGk%3D&from=327834062&s=PackSourceEnum_AWEME_DETAIL&se=false&sc=origin_cover&biz_tag=aweme_video&l=2026012921263252B4EA0EC35D84FA6928"
                    ],
                    "height": 360,
                    "width": 860
                },
                "cover_original_scale": {
                    "uri": "tos-cn-i-0813c001/oE7MEC2DaBAjBQAUAYi3Ae0ABATggC0vIVfT7i",
                    "url_list": [
                        "https://p3-sign.douyinpic.com/tos-cn-i-0813c001/oE7MEC2DaBAjBQAUAYi3Ae0ABATggC0vIVfT7i~tplv-dy-360p.webp?lk3s=138a59ce&x-expires=1770901200&x-signature=DA53ZRahSBCd7g8Um6k0GJKqfkE%3D&from=327834062&s=PackSourceEnum_AWEME_DETAIL&se=false&sc=origin_cover&biz_tag=aweme_video&l=2026012921263252B4EA0EC35D84FA6928",
                        "https://p3-sign.douyinpic.com/tos-cn-i-0813c001/oE7MEC2DaBAjBQAUAYi3Ae0ABATggC0vIVfT7i~tplv-dy-360p.jpeg?lk3s=138a59ce&x-expires=1770901200&x-signature=3glRweZV9vBWa9X%2B0Q1MIxweCC4%3D&from=327834062&s=PackSourceEnum_AWEME_DETAIL&se=false&sc=origin_cover&biz_tag=aweme_video&l=2026012921263252B4EA0EC35D84FA6928"
                    ],
                    "height": 360,
                    "width": 860
                },
                "meta": "{\"format\":\"mp4\",\"frame_select_info\":\"{\\\"top_frame_score\\\":8.828560596663243,\\\"bottom_frame_score\\\":7.162414529280742,\\\"color_diff_score\\\":12.418807385321154,\\\"black_score\\\":0.09523809523809523}\",\"gear_vqm\":\"{\\\"1080p_720p\\\":-1,\\\"720p_540p\\\":-1}\",\"hrids\":\"140022\",\"hv_text_ocr_info\":\"{\\\"width\\\":640,\\\"height\\\":268,\\\"union_text_rect\\\":{\\\"left_top\\\":{\\\"x\\\":0,\\\"y\\\":65},\\\"right_down\\\":{\\\"x\\\":638,\\\"y\\\":266}}}\",\"is_spatial_video\":\"0\",\"isad\":\"0\",\"item_info_video_tags_v2_level_1\":\"607\",\"item_info_video_tags_v2_level_2\":\"60701\",\"item_info_video_tags_v2_level_3\":\"6070106\",\"loudness\":\"-16.1\",\"peak\":\"0.71614\",\"play_time_prob_dist\":\"[900,0.3575,9059.48]\",\"qprf\":\"0.7634180784225464\",\"sdgs\":\"[\\\"comet_bvc1_s1_r1_adapt_lowest_720_1\\\"]\",\"sr_potential\":\"{\\\"v1.0\\\":{\\\"score\\\":60.596}}\",\"sr_score\":\"0.000\",\"strategy_tokens\":\"[\\\"online\\\",\\\"s_comet01_1229\\\",\\\"s_comet02_1229\\\"]\",\"vqs_origin\":\"67.32\"}",
                "download_addr": {
                    "url_list": [
                        "https://v5-dy-o-abtest.zjcdn.com/52483c1fb6fc8730705cf8e4094b5492/697b874b/video/tos/cn/tos-cn-v-694b40/oIS9TFbFJmgeE5wqQPC2NCfPATLeUIelEAtGxI/?a=1128&ch=26&cr=13&dr=0&lr=all&cd=0%7C0%7C0%7C&cv=1&br=1688&bt=1688&cs=0&ds=3&ft=tKyhWzoTd95~~iSYsiI9B4yoFIKrHgSFS4v7p0HCnzXnZeMdxRyqkZ&mime_type=video_mp4&qs=0&rc=NWhoNzk4N2ZkaGc5Ozs6N0BpanBnNWs5cjxqODYzNDVpM0AuNmI0Ml42X14xMjBhNDA1YSNzcF9iMmRzYHFhLS1kMi9zcw%3D%3D&btag=c0010e000b8009&cdn_type=1&cquery=100b_104i_105e_104O_104U&dy_q=1769693192&feature_id=dc6e471fd69cf67451806384e04f2b47&l=2026012921263252B4EA0EC35D84FA6928&pdp=online_xy&pwid=1&req_cdn_type=r",
                        "https://v5-hl-zenl-ov.zjcdn.com/b8d9119af54e4dbd38d0c153f5951a8a/697b874b/video/tos/cn/tos-cn-v-694b40/oIS9TFbFJmgeE5wqQPC2NCfPATLeUIelEAtGxI/?a=1128&ch=26&cr=13&dr=0&lr=all&cd=0%7C0%7C0%7C&cv=1&br=1688&bt=1688&cs=0&ds=3&ft=tKyhWzoTd95~~iSYsiI9B4yoFIKrHgSFS4v7p0HCnzXnZeMdxRyqkZ&mime_type=video_mp4&qs=0&rc=NWhoNzk4N2ZkaGc5Ozs6N0BpanBnNWs5cjxqODYzNDVpM0AuNmI0Ml42X14xMjBhNDA1YSNzcF9iMmRzYHFhLS1kMi9zcw%3D%3D&btag=c0010e000b8009&cdn_type=1&cquery=104O_104U_100b_104i_105e&dy_q=1769693192&feature_id=dc6e471fd69cf67451806384e04f2b47&l=2026012921263252B4EA0EC35D84FA6928&pdp=online_xy&pwid=1&req_cdn_type=r",
                        "https://api-play.amemv.com/aweme/v1/play/?video_id=v02f52g10003d5tgc9vog65he2bdsmvg&line=0&ratio=720p&watermark=1&media_type=4&vr_type=0&improve_bitrate=0&biz_sign=NRPI1lC4AlkgB1YZ_o3Jx6YaRSRKvePEceaWstk3LWuwTcZpAZwDks4E_6LT9VaaOzLa7CA0Ou-V1V2BgVkbe98Q7EbHDGbNHaUZIZuaRMUGyO7s5wbVW7et0I2xs1hA&logo_name=aweme_search_suffix&source=PackSourceEnum_AWEME_DETAIL",
                        "https://api.amemv.com/aweme/v1/play/?video_id=v02f52g10003d5tgc9vog65he2bdsmvg&line=1&ratio=720p&watermark=1&media_type=4&vr_type=0&improve_bitrate=0&biz_sign=NRPI1lC4AlkgB1YZ_o3Jx6YaRSRKvePEceaWstk3LWuwTcZpAZwDks4E_6LT9VaaOzLa7CA0Ou-V1V2BgVkbe98Q7EbHDGbNHaUZIZuaRMUGyO7s5wbVW7et0I2xs1hA&logo_name=aweme_search_suffix&source=PackSourceEnum_AWEME_DETAIL"
                    ],
                    "file_cs": "c:0-4919522-b0fc",
                    "uri": "v02f52g10003d5tgc9vog65he2bdsmvg",
                    "width": 720,
                    "data_size": 1395332813,
                    "height": 720
                },
                "cdn_url_expired": 1769703243,
                "is_h265": 0,
                "play_addr_h264": {
                    "url_key": "v02f52g10003d5tgc9vog65he2bdsmvg_h264_720p_1303625",
                    "width": 1720,
                    "file_hash": "ad2f348983404cae91be6df71b50964f",
                    "file_cs": "c:0-4919522-b0fc",
                    "uri": "v02f52g10003d5tgc9vog65he2bdsmvg",
                    "url_list": [
                        "https://v5-dy-o-abtest.zjcdn.com/cd7c9c7a1ba12e2456e510445238abe6/697b874b/video/tos/cn/tos-cn-v-694b40/oQWPxFfeQIbtu7qCgDSq7GnE9I8lI5meAfwLRT/?a=1128&ch=26&cr=13&dr=0&lr=all&cd=0%7C0%7C0%7C&cv=1&br=1273&bt=1273&cs=0&ds=3&ft=tKyhWzoTd95~~iSYsiI9B4yoFIKrHgSFS4v7p0HCnzXnZeMdxRyqkZ&mime_type=video_mp4&qs=0&rc=ZjtoODU0PGk5OGQ0NTM6ZkBpanBnNWs5cjxqODYzNDVpM0A0LV4yYWNiXzAxLzZfNDBeYSNzcF9iMmRzYHFhLS1kMi9zcw%3D%3D&btag=c0010e000b8009&cdn_type=1&cquery=104U_100b_104i_105e_104O&dy_q=1769693192&feature_id=fea919893f650a8c49286568590446ef&l=2026012921263252B4EA0EC35D84FA6928&pdp=online_xy&pwid=1&req_cdn_type=r",
                        "https://v5-hl-zenl-ov.zjcdn.com/e17d5573516e493b3f7f2d0c4903d533/697b874b/video/tos/cn/tos-cn-v-694b40/oQWPxFfeQIbtu7qCgDSq7GnE9I8lI5meAfwLRT/?a=1128&ch=26&cr=13&dr=0&lr=all&cd=0%7C0%7C0%7C&cv=1&br=1273&bt=1273&cs=0&ds=3&ft=tKyhWzoTd95~~iSYsiI9B4yoFIKrHgSFS4v7p0HCnzXnZeMdxRyqkZ&mime_type=video_mp4&qs=0&rc=ZjtoODU0PGk5OGQ0NTM6ZkBpanBnNWs5cjxqODYzNDVpM0A0LV4yYWNiXzAxLzZfNDBeYSNzcF9iMmRzYHFhLS1kMi9zcw%3D%3D&btag=c0010e000b8009&cdn_type=1&cquery=105e_104O_104U_100b_104i&dy_q=1769693192&feature_id=fea919893f650a8c49286568590446ef&l=2026012921263252B4EA0EC35D84FA6928&pdp=online_xy&pwid=1&req_cdn_type=r",
                        "https://api-play.amemv.com/aweme/v1/play/?video_id=v02f52g10003d5tgc9vog65he2bdsmvg&line=0&file_id=ac3770b2912a4536bfe9f1d255ae087c&sign=ad2f348983404cae91be6df71b50964f&is_play_url=1&source=PackSourceEnum_AWEME_DETAIL",
                        "https://api.amemv.com/aweme/v1/play/?video_id=v02f52g10003d5tgc9vog65he2bdsmvg&line=1&file_id=ac3770b2912a4536bfe9f1d255ae087c&sign=ad2f348983404cae91be6df71b50964f&is_play_url=1&source=PackSourceEnum_AWEME_DETAIL"
                    ],
                    "height": 720,
                    "data_size": 1051300289
                },
                "dynamic_cover": {
                    "width": 720,
                    "uri": "tos-cn-i-0813c001/oE7MEC2DaBAjBQAUAYi3Ae0ABATggC0vIVfT7i",
                    "height": 720,
                    "url_list": [
                        "https://p3-sign.douyinpic.com/obj/tos-cn-i-0813c001/oE7MEC2DaBAjBQAUAYi3Ae0ABATggC0vIVfT7i?lk3s=138a59ce&x-expires=1770901200&x-signature=BNMaWXIbNJ9jHRBkdg3%2B0Hhmp0I%3D&from=327834062_large&s=PackSourceEnum_AWEME_DETAIL&se=false&sc=dynamic_cover&biz_tag=aweme_video&l=2026012921263252B4EA0EC35D84FA6928"
                    ]
                },
                "animated_cover": {
                    "uri": "tos-cn-i-0813c001/oE7MEC2DaBAjBQAUAYi3Ae0ABATggC0vIVfT7i",
                    "url_list": [
                        "https://p3-sign.douyinpic.com/obj/tos-cn-i-0813c001/oE7MEC2DaBAjBQAUAYi3Ae0ABATggC0vIVfT7i?lk3s=138a59ce&x-expires=1770901200&x-signature=BNMaWXIbNJ9jHRBkdg3%2B0Hhmp0I%3D&from=327834062_large&s=PackSourceEnum_AWEME_DETAIL&se=false&sc=dynamic_cover&biz_tag=aweme_video&l=2026012921263252B4EA0EC35D84FA6928"
                    ]
                },
                "play_addr_265": {
                    "height": 720,
                    "width": 1720,
                    "data_size": 752145758,
                    "file_hash": "da9f8c238097b54f59364ac5aa30e835",
                    "uri": "v02f52g10003d5tgc9vog65he2bdsmvg",
                    "url_key": "v02f52g10003d5tgc9vog65he2bdsmvg_bytevc1_720p_932669",
                    "file_cs": "c:0-4976074-c8e6",
                    "url_list": [
                        "https://v5-dy-o-abtest.zjcdn.com/e925365269079b1ca870c6ddd8f71e5c/697b874b/video/tos/cn/tos-cn-v-694b40/oQSSebD8ICEENVivfqfAExTF2w7n4ftYKE4TIP/?a=1128&ch=26&cr=13&dr=0&lr=all&cd=0%7C0%7C0%7C&cv=1&br=910&bt=910&cs=2&ds=3&ft=tKyhWzoTd95~~iSYsiI9B4yoFIKrHgSFS4v7p0HCnzXnZeMdxRyqkZ&mime_type=video_mp4&qs=15&rc=aTg4Ojo8ZDdnPDc4ZDw6ZUBpanBnNWs5cjxqODYzNDVpM0AvM18vYTQ2XmMxLTVhYDI2YSNzcF9iMmRzYHFhLS1kMi9zcw%3D%3D&btag=c0010e000b8009&cdn_type=1&cquery=105e_104O_104U_100b_104i&dy_q=1769693192&feature_id=eceb2a8579548afb597413997a0403f2&l=2026012921263252B4EA0EC35D84FA6928&pdp=online_xy&pwid=1&req_cdn_type=r",
                        "https://v5-hl-zenl-ov.zjcdn.com/f4ff7b097f60af41558b095c4d509b45/697b874b/video/tos/cn/tos-cn-v-694b40/oQSSebD8ICEENVivfqfAExTF2w7n4ftYKE4TIP/?a=1128&ch=26&cr=13&dr=0&lr=all&cd=0%7C0%7C0%7C&cv=1&br=910&bt=910&cs=2&ds=3&ft=tKyhWzoTd95~~iSYsiI9B4yoFIKrHgSFS4v7p0HCnzXnZeMdxRyqkZ&mime_type=video_mp4&qs=15&rc=aTg4Ojo8ZDdnPDc4ZDw6ZUBpanBnNWs5cjxqODYzNDVpM0AvM18vYTQ2XmMxLTVhYDI2YSNzcF9iMmRzYHFhLS1kMi9zcw%3D%3D&btag=c0010e000b8009&cdn_type=1&cquery=104i_105e_104O_104U_100b&dy_q=1769693192&feature_id=eceb2a8579548afb597413997a0403f2&l=2026012921263252B4EA0EC35D84FA6928&pdp=online_xy&pwid=1&req_cdn_type=r",
                        "https://api-play.amemv.com/aweme/v1/play/?video_id=v02f52g10003d5tgc9vog65he2bdsmvg&line=0&file_id=9b57c9da850449fda49a79d727b5652f&sign=da9f8c238097b54f59364ac5aa30e835&is_play_url=1&source=PackSourceEnum_AWEME_DETAIL",
                        "https://api.amemv.com/aweme/v1/play/?video_id=v02f52g10003d5tgc9vog65he2bdsmvg&line=1&file_id=9b57c9da850449fda49a79d727b5652f&sign=da9f8c238097b54f59364ac5aa30e835&is_play_url=1&source=PackSourceEnum_AWEME_DETAIL"
                    ]
                },
                "tags": null,
                "width": 1720,
                "height": 720,
                "is_bytevc1": 0,
                "big_thumbs": [
                    {
                        "img_x_len": 10,
                        "img_y_len": 129,
                        "interval": 5,
                        "uris": [],
                        "img_num": 1290,
                        "img_url": "https://p26-sign.douyinpic.com/tos-cn-p-694b40/oUThtCn2firItQ8SfwIFEY8eIESfnGX7tADBq4~tplv-noop.image?cquery=104i_105e_104O_104U_100b&dy_q=1769693192&l=2026012921263252B4EA0EC35D84FA6928&x-expires=1769703243&x-signature=J8B5mWWSDlJUYU2oxgbe%2B0OWVeI%3D",
                        "img_x_size": 240,
                        "img_y_size": 100,
                        "img_urls": [],
                        "uri": "tos-cn-p-694b40/oUThtCn2firItQ8SfwIFEY8eIESfnGX7tADBq4",
                        "duration": 6451.4585,
                        "fext": "jpg"
                    }
                ],
                "play_addr": {
                    "uri": "v02f52g10003d5tgc9vog65he2bdsmvg",
                    "url_list": [
                        "https://v5-dy-o-abtest.zjcdn.com/cd7c9c7a1ba12e2456e510445238abe6/697b874b/video/tos/cn/tos-cn-v-694b40/oQWPxFfeQIbtu7qCgDSq7GnE9I8lI5meAfwLRT/?a=1128&ch=26&cr=13&dr=0&lr=all&cd=0%7C0%7C0%7C&cv=1&br=1273&bt=1273&cs=0&ds=3&ft=tKyhWzoTd95~~iSYsiI9B4yoFIKrHgSFS4v7p0HCnzXnZeMdxRyqkZ&mime_type=video_mp4&qs=0&rc=ZjtoODU0PGk5OGQ0NTM6ZkBpanBnNWs5cjxqODYzNDVpM0A0LV4yYWNiXzAxLzZfNDBeYSNzcF9iMmRzYHFhLS1kMi9zcw%3D%3D&btag=c0010e000b8009&cdn_type=1&cquery=104U_100b_104i_105e_104O&dy_q=1769693192&feature_id=fea919893f650a8c49286568590446ef&l=2026012921263252B4EA0EC35D84FA6928&pdp=online_xy&pwid=1&req_cdn_type=r",
                        "https://v5-hl-zenl-ov.zjcdn.com/e17d5573516e493b3f7f2d0c4903d533/697b874b/video/tos/cn/tos-cn-v-694b40/oQWPxFfeQIbtu7qCgDSq7GnE9I8lI5meAfwLRT/?a=1128&ch=26&cr=13&dr=0&lr=all&cd=0%7C0%7C0%7C&cv=1&br=1273&bt=1273&cs=0&ds=3&ft=tKyhWzoTd95~~iSYsiI9B4yoFIKrHgSFS4v7p0HCnzXnZeMdxRyqkZ&mime_type=video_mp4&qs=0&rc=ZjtoODU0PGk5OGQ0NTM6ZkBpanBnNWs5cjxqODYzNDVpM0A0LV4yYWNiXzAxLzZfNDBeYSNzcF9iMmRzYHFhLS1kMi9zcw%3D%3D&btag=c0010e000b8009&cdn_type=1&cquery=105e_104O_104U_100b_104i&dy_q=1769693192&feature_id=fea919893f650a8c49286568590446ef&l=2026012921263252B4EA0EC35D84FA6928&pdp=online_xy&pwid=1&req_cdn_type=r",
                        "https://api-play.amemv.com/aweme/v1/play/?video_id=v02f52g10003d5tgc9vog65he2bdsmvg&line=0&file_id=ac3770b2912a4536bfe9f1d255ae087c&sign=ad2f348983404cae91be6df71b50964f&is_play_url=1&source=PackSourceEnum_AWEME_DETAIL",
                        "https://api.amemv.com/aweme/v1/play/?video_id=v02f52g10003d5tgc9vog65he2bdsmvg&line=1&file_id=ac3770b2912a4536bfe9f1d255ae087c&sign=ad2f348983404cae91be6df71b50964f&is_play_url=1&source=PackSourceEnum_AWEME_DETAIL"
                    ],
                    "height": 720,
                    "data_size": 1051300289,
                    "file_cs": "c:0-4919522-b0fc",
                    "url_key": "v02f52g10003d5tgc9vog65he2bdsmvg_h264_720p_1303625",
                    "width": 1720,
                    "file_hash": "ad2f348983404cae91be6df71b50964f"
                },
                "is_callback": true,
                "need_set_token": false
            },
            "mark_largely_following": false,
            "video_text": [],
            "common_bar_info":"[{\"id\":\"related_search_anchor\",\"view_type\":3,\"business_type\":1,\"schema\":\"\",\"separator_dot\":{\"show\":1},\"main_content\":{\"text\":\"相关搜索\"},\"tail_icon\":{\"url\":\"\",\"show\":1,\"icon_url\":{\"uri\":\"tos-cn-i-3a9ajr9h4x/ic_s_s_arrowright_outlined_12.png\",\"url_list\":[\"https://p3-sign.douyinpic.com/tos-cn-i-3a9ajr9h4x/ic_s_s_arrowright_outlined_12.png~noop.png?lk3s=138a59ce\&x-expires=1770901200\&x-signature=KqZ2vfBfiB8gZ4sqiWzoNUJgogA%3D\&from=327834062\",\"https://p3-sign.douyinpic.com/tos-cn-i-3a9ajr9h4x/ic_s_s_arrowright_outlined_12.png~noop.jpeg?lk3s=138a59ce\&x-expires=1770901200\&x-signature=QNBy6G7aYzwMkIMoowOz3eolH9A%3D\&from=327834062\"]}},\"bar_container\":{\"height\":40,\"background_color\":\"#57292929\"},\"extra\":\"{\\\"suggest_words\\\":{\\\"words\\\":[{\\\"word\\\":\\\"疯狂动物城2\\\",\\\"word_id\\\":\\\"6595537046654162180\\\",\\\"info\\\":\\\"{\\\\\\\"ecpm_boost_tag\\\\\\\":false,\\\\\\\"start_ts\\\\\\\":1}\\\",\\\"schema\\\":\\\"aweme://search?needBack2Origin=1\\\&type=general\\\&boost_ecom_word=0\\\&client_engine_extra={\\\\\\\"boost_ecom_word\\\\\\\":0,\\\\\\\"qrec_for_search\\\\\\\":\\\\\\\"{}\\\\\\\"}\\\&keyword=%E7%96%AF%E7%8B%82%E5%8A%A8%E7%89%A9%E5%9F%8E2\\\"},{\\\"word\\\":\\\"兔子警官朱迪与狐狸尼克\\\",\\\"word_id\\\":\\\"6651174348876944654\\\",\\\"info\\\":\\\"{\\\\\\\"ecpm_boost_tag\\\\\\\":false,\\\\\\\"start_ts\\\\\\\":2}\\\",\\\"schema\\\":\\\"aweme://search?client_engine_extra={\\\\\\\"boost_ecom_word\\\\\\\":0,\\\\\\\"qrec_for_search\\\\\\\":\\\\\\\"{}\\\\\\\"}\\\&keyword=%E5%85%94%E5%AD%90%E8%AD%A6%E5%AE%98%E6%9C%B1%E8%BF%AA%E4%B8%8E%E7%8B%90%E7%8B%B8%E5%B0%BC%E5%85%8B\\\&needBack2Origin=1\\\&type=general\\\&boost_ecom_word=0\\\"}],\\\"bar_type\\\":1,\\\"ui_info\\\":{\\\"text\\\":[{\\\"text\\\":\\\"相关搜索\\\",\\\"color\\\":\\\"#FFFFFFFF\\\",\\\"is_bold\\\":1,\\\"type\\\":1}]},\\\"extra_info\\\":{\\\"gold_priority_show\\\":1,\\\"is_life_intent\\\":1,\\\"multi_intention\\\":5,\\\"resp_from\\\":\\\"normal\\\"}}}\",\"tracer_info\":\"{\\\"main_structure_req_path\\\":\\\"/aweme/v1/multi/aweme/detail/\\\"}\",\"name\":\"相关搜索\",\"desc\":\"\",\"show_strategy\":{\"display_identity_status\":3},\"priority\":53500}]",
            "aweme_control": {
                "can_forward": true,
                "can_share": true,
                "can_comment": true,
                "can_show_comment": true
            },
            "entertainment_video_paid_way": {
                "paid_ways": [],
                "paid_type": 0,
                "enable_use_new_ent_data": true
            },
            "video_tag": [
                {
                    "tag_id": 2014,
                    "tag_name": "二次元",
                    "level": 1
                },
                {
                    "tag_id": 2014002,
                    "tag_name": "二次元内容",
                    "level": 2
                },
                {
                    "tag_id": 2014002001,
                    "tag_name": "动漫IP",
                    "level": 3
                }
            ],
            "shoot_way": "",
            "share_url": "https://www.iesdouyin.com/share/video/7600676564674478720/?region=US&mid=7600676840910375726&u_code=-1&did=MS4wLjABAAAACHr-q1RAppEvOQ7rBnckdpl31xTk2fWrd5RKxN--gy7KLaFepUBr0M0_xUV_yduf&iid=MS4wLjABAAAA6JLDyzPypdMiuicVMAm3_bFAtlJhNNRAl5gjJXUgkLrUp9XUDDaEp-tWO_lhxsPM&with_sec_did=1&video_share_track_ver=&titleType=title&share_sign=r5iLYYWEP5q062vYm31bH51pgD3pTqeMZVpM.9KLESg-&share_version=270500&ts=1769693192&from_aid=1128&from_ssr=1&share_track_info=%7B%22link_description_type%22%3A%22%22%7D",
            "publish_plus_alienation": {
                "alienation_type": 0
            },
            "feed_comment_config": {
                "audio_comment_permission": 0,
                "input_config_text": "有爱评论，说点儿好听的",
                "image_challenge": {
                    "inner_publish_text": "我也发一张"
                },
                "common_flags": "{\"item_author_nickname\":\"一发没\",\"need_personal_rec\":\"1\",\"video_labels_v2_tag1\":\"二次元\",\"video_labels_v2_tag2\":\"动漫剧场\"}",
                "author_audit_status": 0
            },
            "impression_data": {
                "group_id_list_c": [],
                "group_id_list_d": [],
                "group_id_list_a": [],
                "group_id_list_b": [
                    7590816633599560372,
                    7592567101590112884
                ],
                "similar_id_list_a": null,
                "similar_id_list_b": null
            },
            "risk_infos": {
                "content": "",
                "risk_sink": false,
                "warn": false,
                "type": 0,
                "vote": false
            },
            "preview_video_status": 0,
            "incentive_item_type": 0,
            "flash_mob_trends": 0,
            "is_preview": 0,
            "poi_patch_info": {
                "item_patch_poi_prompt_mark": 0,
                "extra": ""
            },
            "city": "360100",
            "follow_shoot_clip_info": {
                "clip_from_user": 7600676840910375726,
                "clip_video_all": 7600676840910375726
            },
            "admire_auth": {
                "exit_admire_in_aweme_post": 0,
                "is_show_admire_button": 0,
                "admire_button": 0,
                "is_iron_fans_in_aweme_post": 0,
                "is_click_admire_icon_recently": 0,
                "author_can_admire": 0,
                "is_show_admire_tab": 0,
                "is_admire": 0,
                "is_fifty_admire_author_stable_fans": 0
            },
            "is_moment_history": 0,
            "long_video": null,
            "video_control": {
                "allow_download": false,
                "timer_status": 1,
                "share_ignore_visibility": false,
                "download_info": {
                    "level": 1,
                    "fail_info": {
                        "code": 290002,
                        "reason": "isolation_content",
                        "msg": "视频暂时无法保存，链接已复制"
                    }
                },
                "duet_info": {
                    "level": 2,
                    "fail_info": {
                        "code": 100014,
                        "reason": "video_duration"
                    }
                },
                "allow_record": true,
                "draft_progress_bar": 1,
                "prevent_download_type": 2,
                "download_ignore_visibility": false,
                "show_progress_bar": 1,
                "share_grayed": false,
                "timer_info": {
                    "public_time": 1769683575,
                    "timer_status": 0
                },
                "share_type": 0,
                "allow_music": false,
                "allow_stitch": true,
                "allow_share": true,
                "duet_ignore_visibility": false,
                "allow_dynamic_wallpaper": false,
                "allow_douplus": true,
                "disable_record_reason": "",
                "allow_duet": true,
                "allow_react": true
            },
            "has_vs_entry": false,
            "cmt_swt": false,
            "misc_info": "{\"common_business_mob\":\"{}\",\"is_teen_video\":0}",
            "image_crop_ctrl": 0,
            "is_first_video": false,
            "visual_search_info": {
                "is_show_img_entrance": false,
                "is_high_recall_ecom": false,
                "is_ecom_img": false,
                "is_high_accuracy_ecom": false
            },
            "is_ads": false,
            "ecom_comment_atmosphere_type": 0,
            "status": {
                "is_private": false,
                "download_status": 0,
                "private_status": 0,
                "aweme_id": "7600676564674478720",
                "with_fusion_goods": false,
                "allow_comment": true,
                "part_see": 0,
                "video_hide_search": 0,
                "dont_share_status": 0,
                "allow_friend_recommend_guide": false,
                "is_delete": false,
                "allow_friend_recommend": false,
                "enable_soft_delete": 0,
                "with_goods": false,
                "aweme_edit_info": {
                    "edit_status": 0,
                    "has_modified_all": false,
                    "button_status": 2,
                    "button_toast": ""
                },
                "review_result": {
                    "review_status": 0
                },
                "self_see": false,
                "allow_share": true,
                "reviewed": 0,
                "listen_video_status": 2,
                "not_allow_soft_del_reason": "ab",
                "is_prohibited": false,
                "allow_self_recommend_to_friend": true,
                "in_reviewing": false
            },
            "user_recommend_status": 1,
            "is_from_ad_auth": false,
            "press_panel_info": "[{\"interactive\":[\"2_story\",\"2_friend\"]},{\"feedback\":[\"rr_feedback\",\"dislike\",\"ignore\",\"block\",\"unfollow\",\"sever\",\"dislike_collect\"]},{\"control\":[\"speed\",\"auth\",\"delete\",\"save\",\"collect\",\"reward\",\"bg_play\",\"duet\",\"together\"]}]",
            "is_in_scope": false,
            "is_pgcshow": false,
            "distance": "",
            "item_title": "疯狂动物城2（国语）",
            "aweme_type": 4,
            "duration": 6451549,
            "comment_permission_info": {
                "toast_guide": false,
                "comment_permission_status": 0,
                "can_comment": true,
                "item_detail_entry": true,
                "press_entry": true
            },
            "cha_list": null,
            "statistics": {
                "lose_comment_count": 0,
                "digg_count": 181,
                "exposure_count": 0,
                "collect_count": 159,
                "admire_count": 0,
                "lose_count": 0,
                "live_watch_count": 0,
                "download_count": 0,
                "aweme_id": "7600676564674478720",
                "digest": "",
                "share_count": 97,
                "play_count": 0,
                "forward_count": 0,
                "comment_count": 12,
                "whatsapp_share_count": 0
            },
            "hybrid_label": null,
            "bodydance_score": 0,
            "commerce_config_data": null,
            "product_genre_info": {
                "material_genre_sub_type_set": [
                    4
                ],
                "special_info": {
                    "recommend_group_name": 0
                },
                "product_genre_type": 2
            },
            "desc": "疯狂动物城2（国语） ",
            "nickname_position": null,
            "game_tag_info": {
                "is_game": false
            },
            "duet_aggregate_in_music_tab": false,
            "is_karaoke": false,
            "is_story": 0,
            "item_duet": 0,
            "item_react": 0,
            "media_type": 4,
            "chapter_list": null,
            "comment_list": null,
            "series_paid_info": {
                "series_paid_status": 0,
                "item_price": 0
            },
            "position": null,
            "guide_btn_type": 0,
            "personal_page_botton_diagnose_style": 0,
            "xigua_task": {
                "is_xigua_task": false
            },
            "image_list": null,
            "commerce_info": {
                "ad_type": 0,
                "is_ad": false
            },
            "dislike_dimension_list": null,
            "suggest_words": {
                "suggest_words": [
                    {
                        "words": [
                            {
                                "word": "疯狂动物城2",
                                "word_id": "6595537046654162180",
                                "info": "{\"End\":0,\"Start\":0,\"ecpm_boost_tag\":false,\"log_pb\":\"\",\"qrec_for_search\":\"{}\"}"
                            }
                        ],
                        "scene": "detail_inbox_rex",
                        "icon_url": "",
                        "hint_text": "",
                        "extra_info": "{\"is_life_intent\":1,\"resp_from\":\"normal\"}"
                    },
                    {
                        "words": [
                            {
                                "word": "疯狂动物城2",
                                "word_id": "6595537046654162180",
                                "info": "{\"End\":0,\"Start\":0,\"ecpm_boost_tag\":false,\"log_pb\":\"\",\"qrec_for_search\":\"{}\"}"
                            }
                        ],
                        "scene": "comment_top_rec",
                        "icon_url": "",
                        "hint_text": "大家都在搜：",
                        "extra_info": "{\"is_life_intent\":1,\"resp_from\":\"normal\"}"
                    }
                ]
            },
            "main_arch_common": "{\"music_detail_fail_reason\":\"other\",\"music_detail_fail_type\":14,\"music_detail_fail_toast\":\"该声音不可用\"}",
            "author_user_id": 2543612830166135,
            "item_stitch": 0,
            "nearby_level": 0,
            "image_comment": {},
            "share_rec_extra": "",
            "sort_label": "",
            "with_promotional_music": false,
            "series_basic_info": {},
            "should_open_ad_report": false,
            "cf_assets_type": 0,
            "desc_language": "un",
            "horizontal_type": 1,
            "sec_item_id": "MS4wLjAAAAAA4uGpj-c6firmlZlkggwAy0zCBDhZNJIEs9RROaYWdrSNCAHS1hN7FR-9jz49hP6P",
            "is_life_item": false,
            "can_cache_to_local": true,
            "cf_recheck_ts": 0,
            "item_comment_settings": 0,
            "report_action": false,
            "origin_duet_resource_uri": "",
            "distribute_type": 1,
            "author_mask_tag": 0,
            "is_new_text_mode": 0,
            "video_game_data_channel_config": {},
            "promotions": [],
            "is_vr": false,
            "collection_corner_mark": 0,
            "share_info": {
                "share_desc": "在抖音，记录美好生活",
                "share_link_desc": "2.05 复制打开抖音，看看【一发没的作品】疯狂动物城2（国语）   %s 03/16 yGv:/ o@Q.xS ",
                "share_desc_info": "#在抖音，记录美好生活#疯狂动物城2（国语） ",
                "share_weibo_desc": "#在抖音，记录美好生活#疯狂动物城2（国语） ",
                "bool_persist": 0,
                "share_quote": "",
                "share_signature_desc": "",
                "share_title": "疯狂动物城2（国语） ",
                "share_signature_url": "",
                "share_title_other": "",
                "share_url": "https://www.iesdouyin.com/share/video/7600676564674478720/?region=US&mid=7600676840910375726&u_code=-1&did=MS4wLjABAAAACHr-q1RAppEvOQ7rBnckdpl31xTk2fWrd5RKxN--gy7KLaFepUBr0M0_xUV_yduf&iid=MS4wLjABAAAA6JLDyzPypdMiuicVMAm3_bFAtlJhNNRAl5gjJXUgkLrUp9XUDDaEp-tWO_lhxsPM&with_sec_did=1&video_share_track_ver=&titleType=title&share_sign=r5iLYYWEP5q062vYm31bH51pgD3pTqeMZVpM.9KLESg-&share_version=270500&ts=1769693192&from_aid=1128&from_ssr=1&share_track_info=%7B%22link_description_type%22%3A%22%22%7D",
                "share_title_myself": ""
            },
            "video_share_edit_status": 0,
            "is_relieve": false,
            "image_album_music_info": {
                "begin_time": -1,
                "end_time": -1,
                "volume": -1
            },
            "anchors": null,
            "cover_labels": null,
            "photo_search_entrance": {
                "ecom_type": 0
            },
            "aweme_acl": {
                "download_mask_panel": {
                    "show_type": 0,
                    "code": 1
                }
            },
            "images": null,
            "authentication_token": "MS4wLjAAAAAAJAjFlZtk6BVDlkHvXfWi6Uez64xjVvSbemu1JI8qHX1TJpvhCCB7YkJesENePtZz6_ACj9LTCMtmsg6TThfWi6m3Xfi93KlL4Mk5rFfgDyxMmIv6HJlbRlGbVHrM03AtF4JsmJSHJ4YAdo1HJAjXFtgFST_UWfGjMXVuAO1UxEC9JyuVTZp1VZJQ78FxSQ4QYZJk-2MhblbtXZPL3eINXsC-KQPDN16rIEfzNduDSvx-si1rUBzMxuScDHZvvEmk",
            "original_images": null,
            "comment_words_recommend": {},
            "fall_card_struct": {
                "recommend_reason_v2": "[]"
            },
            "poi_biz": {},
            "video_labels": null,
            "enable_decorated_emoji": true,
            "entertainment_recommend_info": "{\"recommend_free_count\":0,\"test_info\":\"\",\"recommend_playlet_trigger_reqid\":\"\",\"recommend_playlet_trigger_cid\":\"\",\"recommend_playlet_trigger_reqid_type\":\"\",\"playlet_continue_play_config\":null,\"roi3_mode\":0,\"predict_identity\":\"\"}"
        }
    }
}