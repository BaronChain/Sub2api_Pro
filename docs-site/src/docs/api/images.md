---
title: 图片生成与编辑接口
icon: image
order: 4
author: FreeCode Team
---

# 图片生成与编辑接口

FreeCode 支持通过 OpenAI Images API 风格调用部分图片模型。具体可用模型和分组以模型广场为准。

## Base URL

```text
https://freecode.codes/v1
```

## 常见路径

| 能力 | 路径 |
|------|------|
| 文生图 | `/v1/images/generations` |
| 图片编辑 / 图生图 | `/v1/images/edits` |

## 文生图示例

```bash
curl --location 'https://freecode.codes/v1/images/generations' \
  --header 'Authorization: Bearer 你的图片分组令牌' \
  --header 'Content-Type: application/json' \
  --data '{
    "model": "gpt-image-2",
    "prompt": "一只在月球上喝咖啡的猫",
    "size": "1024x1024"
  }'
```

## 图片请求注意事项

- 图片生成通常比文本对话耗时更久；
- 如果客户端支持关闭流式输出，绘图场景建议关闭；
- 本机代理、公司网关或浏览器插件可能限制长连接，导致请求超时；
- 建议把 `freecode.codes` 加入代理直连或白名单规则；
- 返回 URL 或 Base64 取决于模型和请求参数。

更多模型示例请看：

- [Nano Banana2 Pro 绘图教程](../paint/Banana.md)
- [GPT-Image-2 绘图教程](../paint/GPTImage.md)
