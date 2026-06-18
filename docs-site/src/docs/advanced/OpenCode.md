---
title: OpenCode
icon: code
order: 3
author: FreeCode Team
date: 2026-01-01
category:
  - 第三方接入
---

# OpenCode

## 项目介绍

OpenCode 是一个开源的 AI 编程助手，可在终端、IDE 或桌面环境中辅助编写、调试和改进代码。

核心特色：

- 原生终端 / TUI 支持，适合命令行开发者
- 自动加载正确的语言服务器（LSP）以提升上下文理解
- 支持多会话并行与会话链接共享
- 支持 75+ 模型提供商，包括本地模型

## 环境配置

打开终端，运行以下命令全局安装 OpenCode：

```bash
npm install -g opencode-ai
```

安装完成后，在终端输入 `opencode`，若出现界面则安装成功。

## 通过 CC-Switch 配置

参考 [CC-Switch 通用步骤](../ccswitch/1-common.md) 下载并安装 CC-Switch，打开软件后：

1. 上方配置项选择到 **OpenCode**，然后点击 **添加供应商**。
2. 新增自定义供应商，端点填写 `https://freecode.codes`。
3. 在 **供应商标识** 中填写分组名称，例如 `FreeCode-Codex`。
4. 在 **接口格式** 中选择合适的格式：
   - Claude 系列模型：**Anthropic**
   - Codex 系列模型：**OpenAI**
   - Gemini 系列模型：**Google (Gemini)**
5. 在 **API Key** 中填入 [创建 API 令牌](../register/4-token.md) 中创建的 Key。
6. 在 **额外选项** 中配置键值对 `{"setCacheKey":true}`。
7. 在 **模型配置** 中配置该分组下正确的模型名（可在 [令牌分组介绍](../token/2-group.md) 查询）。例如 Codex 分组：
   - 模型 ID：`gpt-5.2` 显示名称：`gpt-5.2`
   - 模型 ID：`gpt-5.2-codex` 显示名称：`gpt-5.2-codex`
8. 全部配置好后，点击右下角 **添加**，并在界面中选择刚配置好的渠道点击添加。

## 验证配置

重新打开终端，运行 `opencode`，输入 `/models` 命令，观察刚才配置的渠道是否存在，存在则配置成功。

开始愉快的对话吧~
