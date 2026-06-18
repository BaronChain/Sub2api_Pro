---
title: AionUi
icon: window-maximize
order: 2
author: FreeCode Team
date: 2026-01-01
category:
  - 第三方接入
---

# AionUi

## 介绍

AionUi 是一个用户友好的图形化界面工具，可与 Gemini CLI、Claude Code、Codex、Qwen Code 等 AI Agent 协同工作。主要特性：

- ✅ **统一图形界面** —— 自动识别本地 CLI 工具，告别命令行
- ✅ **多会话并行** —— 同时开启多个对话，互不干扰
- ✅ **本地数据安全** —— 对话和文件保存在本地 SQLite 数据库
- ✅ **多种格式预览** —— PDF、Word、Excel、PPT、代码、Markdown、图片等
- ✅ **多模型切换** —— Gemini、Claude、OpenAI、Qwen、Ollama 等
- ✅ **完全免费开源** —— Apache-2.0 许可证

## 软件下载

macOS：

```bash
brew install aionui
```

或访问 GitHub Releases 页面下载适合你系统的安装包（`.dmg` / `.zip` / `.AppImage` / `.deb`）。

Linux（Debian/Ubuntu）：

```bash
# 请将 x.x.x 替换为实际版本号
wget https://github.com/iOfficeAI/AionUi/releases/latest/download/AionUi-x.x.x-linux-amd64.deb
sudo dpkg -i AionUi-x.x.x-linux-amd64.deb
```

## 配置

### 获取 API

回顾 [创建 API 令牌](../register/4-token.md)，在 FreeCode 中创建对应分组的令牌并复制 API Key：

- Gemini → 创建 Gemini 分组令牌
- Claude → 创建 CC 分组令牌
- Codex → 创建 Codex 分组令牌

### 配置 LLM 模型

打开 AionUi，点击 **设置 → LLM 配置 → 添加模型**，选择平台 **自定义**，根据下方配置填入对应信息，保存后返回主界面选择该模型即可使用。

| 模型类型 | API 请求地址 | 说明 |
| --- | --- | --- |
| Gemini | `https://freecode.codes` | 使用 Gemini 分组 API Key，选择 Gemini 模型 |
| Claude | `https://freecode.codes` | 使用 CC 分组 API Key，选择 Claude 模型 |
| Codex | `https://freecode.codes/v1` | 使用 Codex 分组 API Key，选择 Codex 模型 |
