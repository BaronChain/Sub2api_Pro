---
title: 环境检查（通用步骤）
icon: gears
order: 1
author: FreeCode Team
date: 2026-01-01
category:
  - CLI 配置教程
---

# 环境检查（通用步骤）

## （1）确认 Node.js 环境已安装

在 Windows 或 macOS 终端输入以下命令：

```bash
npm list -g --depth-0
```

正常情况下命令会顺利执行（没有任何内容也没关系）。如果提示“命令未找到”，则说明你没有安装 Node.js，需要先安装运行 Claude Code、Codex、Gemini 所需的环境。

安装完成后，重新执行上述命令，如果不再提示“命令未找到”，则说明安装成功。

## （2）安装 CLI

在 Windows 或 macOS 终端输入以下命令，一次性安装好目前所需的所有 CLI：

```bash
npm i -g @anthropic-ai/claude-code@latest
npm i -g @openai/codex@latest
npm i -g @google/gemini-cli@latest
```

## （3）测试安装成功

::: important 这一步很重要
请务必运行命令进行测试。因为运行命令后，你的用户目录下才会生成各 CLI 的配置目录，方便后续操作！
:::

### Claude Code

在终端输入以下命令，若出现交互界面或选项，则 Claude Code 安装成功：

```bash
claude
```

### Codex

在终端输入以下命令，若出现交互界面或选项，则 Codex 安装成功：

```bash
codex
```

### Gemini

在终端输入以下命令，若出现交互界面或选项，则 Gemini 安装成功：

```bash
gemini
```

环境就绪后，继续阅读 [Claude Code 配置](./2-claude.md)。
