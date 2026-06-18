---
title: 用户侧排查清单
icon: screwdriver-wrench
order: 6
author: FreeCode Team
---

# 用户侧排查清单

当 Claude Code、Codex、Gemini CLI 或其他客户端无法调用时，可以按本页从上到下排查。

## 1. 确认 API 令牌

- Key 是否完整复制，没有多余空格或换行；
- 令牌是否处于启用状态；
- 分组是否匹配客户端；
- 是否设置了 IP 白名单、模型限制或过期时间；
- 是否达到令牌额度或周期限制。

## 2. 确认 Base URL

常见 Base URL：

| 场景 | Base URL |
|------|----------|
| Claude / Anthropic 类客户端 | `https://freecode.codes` |
| OpenAI / Codex 兼容客户端 | `https://freecode.codes/v1` |
| Gemini CLI | `https://freecode.codes` |
| Images API | `https://freecode.codes/v1` |

如果客户端要求填写完整接口地址，请按客户端说明补上 `/chat/completions`、`/responses` 或其他路径。

## 3. 确认模型名

模型名必须和模型广场或客户端文档一致。常见错误包括：

- 多写或少写版本号；
- 把 Claude 模型填到 OpenAI 客户端里；
- 把 Gemini 模型填到 Codex 配置里；
- 客户端缓存旧模型列表。

## 4. 根据状态码判断

| 状态码 | 常见原因 |
|--------|----------|
| 401 | API Key 错误、过期、禁用或认证头写错 |
| 403 | 无权限、分组不匹配、模型受限或策略限制 |
| 404 | 路径错误、Base URL 错误或模型名错误 |
| 429 | 并发、RPM、周期额度或上游限流 |
| 5xx | 上游异常、渠道不可用或服务临时故障 |

## 5. 提交问题时提供什么

为了更快定位，请提供：

- 使用的客户端名称和版本；
- 使用的 Base URL；
- 模型名；
- 报错截图或完整错误文本；
- 大致请求时间；
- API Key 只提供前后几位，不要发完整 Key。
