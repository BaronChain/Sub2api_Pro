---
title: 令牌分组介绍
icon: layer-group
order: 2
author: FreeCode Team
date: 2026-01-01
category:
  - 模型分组介绍
---

# 令牌分组介绍

FreeCode 按上游平台与用途划分了多个令牌分组。创建令牌时选择对应分组，令牌即可调用该分组下的模型。常见分组示意如下（以控制台“模型广场”实时显示为准）：

| 分组 | 适用客户端 / 用途 |
| --- | --- |
| `Default` | 默认分组，按账户默认倍率计费 |
| `CC` | Claude Code 使用的 Claude 系列模型 |
| `Codex` | Codex 使用的 OpenAI 系列模型 |
| `Gemini` | Gemini CLI 使用的 Gemini 系列模型 |
| `Image` | 绘图模型分组（Nano Banana、GPT-Image 等） |
| `*-sale` | 折扣分组，享受更低倍率 |
| `*-officially` | 官方直连分组 |

::: tip 选择建议
- 配置 **Claude Code** → 选择 `CC` 相关分组。
- 配置 **Codex** → 选择 `Codex` 相关分组。
- 配置 **Gemini CLI** → 选择 `Gemini` 相关分组。
- 使用 **绘图模型** → 选择 `Image` 相关分组。

分组的可用模型与倍率请以 [模型广场](./1-intro.md) 实时显示为准。
:::

::: warning 分组选错的后果
如果令牌分组与客户端不匹配，配置后调用时会提示“模型不存在”或无法调用。请务必对照客户端选择正确分组。
:::
