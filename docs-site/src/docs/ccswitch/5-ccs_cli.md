---
title: CC-Switch-CLI 使用
icon: terminal
order: 5
author: FreeCode Team
date: 2026-01-01
category:
  - CC-Switch 使用
---

# CC-Switch CLI 使用

::: tip 提示
CC-Switch CLI 适合服务器、SSH、macOS 终端和自动化场景使用。如果你更习惯图形界面，可以继续使用前面的 CC-Switch 教程。
:::

## CC-Switch CLI 是什么

CC-Switch CLI 是 CC-Switch 的命令行版本，是 Claude Code、Codex、Gemini、OpenCode 与 OpenClaw 的命令行管理工具，支持 MCP、Skills、提示词、本地代理和环境检查等功能。它包含两部分：

- **完整 CLI 命令**：可完成 Provider 列表查看、切换、环境检查、MCP 同步、Skills 管理、提示词管理、本地代理等操作。
- **完整 TUI 界面**：运行 `cc-switch` 后进入终端图形界面，可像桌面版一样新增 Provider、选择模板、填写 API Key、保存并切换配置。

如果你只是第一次配置 FreeCode，推荐先用 TUI。配置完成后，日常切换、检查和排错可以直接用 CLI 命令完成。

## 安装 CC-Switch CLI

macOS 和 Linux 推荐使用一键安装脚本：

```bash
curl -fsSL https://github.com/SaladDay/cc-switch-cli/releases/latest/download/install.sh | bash
```

默认会安装到 `~/.local/bin`。如果终端提示找不到 `cc-switch`，请确认 `~/.local/bin` 已加入 PATH。

### 手动安装

macOS：

```bash
curl -LO https://github.com/saladday/cc-switch-cli/releases/latest/download/cc-switch-cli-darwin-universal.tar.gz
tar -xzf cc-switch-cli-darwin-universal.tar.gz
chmod +x cc-switch
sudo mv cc-switch /usr/local/bin/

# 如遇 "无法验证开发者" 提示
xattr -cr /usr/local/bin/cc-switch
```

Linux x64：

```bash
curl -LO https://github.com/saladday/cc-switch-cli/releases/latest/download/cc-switch-cli-linux-x64-musl.tar.gz
tar -xzf cc-switch-cli-linux-x64-musl.tar.gz
chmod +x cc-switch
sudo mv cc-switch /usr/local/bin/
```

Windows：前往 GitHub Releases 下载 `cc-switch-cli-windows-x64.zip`，解压后将 `cc-switch.exe` 放到 PATH 目录中，或直接在当前目录运行 `.\cc-switch.exe`。

## 两种使用方式

进入 TUI 界面：

```bash
cc-switch
```

如果要直接配置某个应用，可以加 `--app`：

```bash
cc-switch --app claude
cc-switch --app codex
cc-switch --app gemini
```

使用 CLI 命令：

```bash
cc-switch provider list
cc-switch provider current
cc-switch provider switch <id>
cc-switch env tools
cc-switch env check
```

`claude` 是默认应用。管理其他应用时使用 `--app`：

```bash
cc-switch --app codex provider list
cc-switch --app gemini provider current
```

## 配置前准备

请先确认目标 CLI 已经安装：

```bash
cc-switch env tools
```

建议先运行一次目标 CLI 或帮助命令，让它创建自己的配置目录：

```bash
claude --help
codex --help
gemini --help
```

然后在 FreeCode 创建对应分组的令牌：

- Claude Code：创建 CC 分组令牌
- Codex：创建 Codex 分组令牌
- Gemini：创建 Gemini 分组令牌

## 配置 FreeCode

第一次配置推荐使用 TUI。运行 `cc-switch` 进入交互界面：

1. 在左侧选择 **Providers**，进入供应商管理页面，然后新增供应商。
2. 新增一个自定义供应商，端点填写 `https://freecode.codes`。
3. 在 **API Key** 中填入你从 FreeCode 复制的令牌，然后保存。
4. 回到供应商列表，确认当前选中的是刚刚添加的 FreeCode Provider。
5. 如果你配置的是 Claude Code，进入 **设置**，找到 **跳过 Claude Code 初次安装确认** 并开启（会向 `~/.claude.json` 写入 `hasCompletedOnboarding=true`）。
6. 打开对应 CLI 测试：`claude` / `codex` / `gemini`。

## 常用命令

```bash
cc-switch                         # 进入交互界面
cc-switch env tools               # 检查本地 CLI 是否安装
cc-switch env check               # 检查环境变量冲突

cc-switch provider list           # 查看 Claude 供应商
cc-switch provider current        # 查看当前 Claude 供应商
cc-switch provider switch <id>    # 切换 Claude 供应商

cc-switch --app codex provider list
cc-switch --app gemini provider list

cc-switch provider stream-check <id> # 检查供应商流式响应
cc-switch provider fetch-models <id> # 拉取远端模型列表
cc-switch update                     # 更新 CC-Switch CLI
```

## 常见问题

### 切换 Provider 后没有生效

请先确认目标 CLI 已经初始化配置目录。可以运行一次 `claude --help` / `codex --help` / `gemini --help`，然后重新切换一次 Provider。

### 环境变量覆盖了配置

如果系统里设置了 `ANTHROPIC_API_KEY`、`OPENAI_API_KEY`、`GEMINI_API_KEY` 等环境变量，目标 CLI 可能会优先读取环境变量，导致 CC-Switch CLI 写入的配置没有生效。可以运行：

```bash
cc-switch env check --app claude
cc-switch env check --app codex
cc-switch env check --app gemini
```
