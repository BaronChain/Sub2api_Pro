---
title: OpenAI 兼容接口
icon: robot
order: 1
author: FreeCode Team
---

# OpenAI 兼容接口

OpenAI 兼容接口适合 Codex、OpenAI SDK、Cherry Studio、Continue、Cline、OpenCode 等支持 OpenAI API 的客户端。

## Base URL

```text
https://freecode.codes/v1
```

如果客户端要求填写完整接口地址，请在 Base URL 后追加具体路径，例如 `/chat/completions` 或 `/responses`。

## 认证方式

```http
Authorization: Bearer 你的API令牌
```

请使用 Codex/OpenAI 相关分组的 API 令牌。若使用了 Claude 或 Gemini 分组令牌，可能会出现模型不存在或 403。

## 常见接口

| 接口 | 路径 | 说明 |
|------|------|------|
| Models | `/v1/models` | 获取当前令牌可见的模型列表 |
| Chat Completions | `/v1/chat/completions` | 传统 OpenAI 聊天补全接口 |
| Responses | `/v1/responses` | 新版 Responses API，Codex 常用 |
| Embeddings | `/v1/embeddings` | 向量接口，是否可用取决于分组和模型 |
| Images | `/v1/images/generations`、`/v1/images/edits` | 图片生成与编辑 |

## curl 示例

```bash
curl https://freecode.codes/v1/chat/completions \
  -H "Authorization: Bearer 你的API令牌" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "你的模型名",
    "messages": [
      {"role": "user", "content": "你好"}
    ]
  }'
```

## Codex 注意事项

Codex 通常需要使用 Responses API。配置示例请看 [Codex 配置](../cli/3-codex.md)。如果遇到连接失败，请确认：

1. `base_url` 是否为 `https://freecode.codes/v1`；
2. `wire_api` 是否按教程设置；
3. API Key 是否来自 Codex 分组；
4. 本地代理或网络工具没有拦截长连接。
