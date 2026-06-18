---
title: Gemini 兼容接口
icon: robot
order: 3
author: FreeCode Team
---

# Gemini 兼容接口

Gemini CLI 或支持 Gemini 自定义端点的客户端，可以通过 FreeCode 的 Gemini 兼容入口接入。

## Base URL

```text
https://freecode.codes
```

Gemini CLI 常见 `.env` 配置：

```env
GOOGLE_GEMINI_BASE_URL=https://freecode.codes
GEMINI_API_KEY=你的Gemini分组API令牌
```

完整步骤请看 [Gemini 配置](../cli/4-gemini.md)。

## 令牌分组

请使用 Gemini 相关分组的 API 令牌。若使用 Codex、CC 或图片专用分组，可能会出现模型不存在、鉴权失败或调用不可用。

## 第三方客户端

如果客户端只支持 OpenAI 兼容接口，也可以尝试在客户端中选择 OpenAI Compatible，并填写：

```text
https://freecode.codes/v1
```

但这种方式是否可用取决于客户端对 Gemini 模型和响应格式的支持。优先使用客户端原生 Gemini 配置。

## 常见问题

### Gemini CLI 不稳定或无法配置

Gemini CLI 本身对自定义端点和版本变化较敏感。如果配置困难，可以考虑：

- 使用 [CC-Switch](../ccswitch/4-gemini.md)；
- 使用支持自定义 API 的 IDE 插件或客户端；
- 在 [FAQ - Gemini](../faq/Gemini.md) 中查看替代方案。

### Cline 中如何使用 Gemini

请参考 [Gemini FAQ](../faq/Gemini.md) 中的 Cline 配置说明，重点是选择兼容 Provider、填写正确 Base URL 和 Gemini 分组 API Key。
