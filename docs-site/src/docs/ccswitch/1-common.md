---
title: 通用步骤
icon: download
order: 1
author: FreeCode Team
date: 2026-01-01
category:
  - CC-Switch 使用
---

# 通用步骤

## CC-Switch 介绍

CC-Switch 是一个 Claude Code / Codex / Gemini CLI 全方位辅助工具（第三方开源项目），从供应商切换器演进为 AI CLI 一体化管理平台，统一管理 Claude Code、Codex 与 Gemini CLI 的供应商配置、MCP 服务器、Skills 扩展和系统提示词。

使用 CC-Switch，你可以：

- ✅ **一键切换 API 配置** —— 在多个 API 提供商之间快速切换
- ✅ **可视化配置管理** —— 通过图形界面轻松管理所有配置
- ✅ **MCP 服务器管理** —— 管理 Model Context Protocol 服务器
- ✅ **系统托盘快捷操作** —— 通过托盘菜单快速切换

::: tip 关于 FreeCode 配置
在 CC-Switch 中接入 FreeCode 时，新增一个自定义供应商，将端点填写为 `https://freecode.codes`，并填入对应分组的 API Key 即可。
:::

## 软件下载

### Windows

进入 CC-Switch 的 GitHub Release 页面，下载适合自己版本的安装包。Windows 系统推荐下载 `.msi` 后缀的安装包进行安装。安装后运行 CC-Switch 主程序。

### macOS

推荐使用 Homebrew，开启终端后分别运行以下命令：

```bash
# 添加 tap 源
brew tap farion1231/ccswitch

# 安装 CC-Switch
brew install --cask cc-switch
```

安装完成后，在“启动台”或“应用程序”文件夹中找到 CC-Switch 并启动。

### Linux

Debian / Ubuntu 系统：

```bash
# 下载 .deb 包
wget https://github.com/farion1231/cc-switch/releases/latest/download/cc-switch_x.x.x_amd64.deb

# 安装
sudo dpkg -i cc-switch_x.x.x_amd64.deb
```

## 环境检查

::: warning 注意
请你最好进行环境检查步骤！如果你有经验，能确认 Node.js 环境以及 cc、codex、gemini 的 CLI 安装没问题、配置目录也都存在，可以忽略这一步，直接进入后续的 CC-Switch 配置。
:::

如何进行环境检查，请参考 [环境检查（通用步骤）](../cli/1-env.md)。
