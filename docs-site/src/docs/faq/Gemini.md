---
title: Gemini
icon: robot
order: 3
author: FreeCode Team
date: 2026-01-01
category:
  - 常见问题
---

# Gemini 相关问题

## Gemini CLI 使用难题与建议

### 现状说明

Gemini CLI 目前存在多种使用问题，例如可能无法正常调用模型、无法粘贴图片。因此通常不建议将 Gemini-3 接入 Gemini CLI。

### 更推荐的方式

- 优先使用 **Roo Code / Cline** 等第三方 VSCode 插件。
- 如必须使用 Gemini CLI，建议使用更稳定的 Gemini 分组渠道。

::: tip 提示
每个分组支持的模型可在 [令牌分组介绍](../token/2-group.md) 查看，避免配置时出现“无可用渠道”或“模型不存在”问题。在 Roo Code 等第三方使用时，请选取 **OpenAI Response** 请求格式。
:::

## 如何在 Cline 中使用 Gemini-3

### 1. 创建 Gemini 分组令牌

按照 [创建 API 令牌](../register/4-token.md) 创建 gemini 分组的令牌。

### 2. 安装 Cline 插件

1. 打开 VSCode（要求 1.80.0+）。
2. 单击左侧 **扩展** 图标（或按 `Ctrl+Shift+X` / `Cmd+Shift+X`）。
3. 搜索 `Cline`，找到后单击 **安装**。

### 3. 打开 Cline 界面

- 方式一：单击左侧边栏的 Cline 图标。
- 方式二：按 `Ctrl+Shift+P` / `Cmd+Shift+P`，输入 `Cline: Open` 回车。

### 4. 首次配置

单击 **API Configuration** 按钮，填写：

| 配置项 | 推荐值 | 说明 |
| --- | --- | --- |
| API Provider | `OpenAI-compatible` | 推荐选择此项，支持更多模型 |
| Base URL | `https://freecode.codes/v1` | FreeCode 的兼容端点 |
| API Key | `sk-******` | 你的 FreeCode gemini 分组 API Key |
| Model ID | `gemini-3-pro-preview` | 推荐使用的模型 |

::: warning 安全提醒
请妥善保管你的 API Key，不要在群聊或公开截图中泄露。
:::

### 5. 完成配置

单击右上角 **Done**。如果之前使用过 Cline，请单击右上角 ⚙️ 设置按钮进入配置界面。
