---
title: Codex 配置
icon: robot
order: 3
author: FreeCode Team
date: 2026-01-01
category:
  - CLI 配置教程
---

# Codex 配置

## Windows

1. 键盘按下 `Win + R`，输入以下内容后回车，打开 Codex 配置目录：

   ```text
   %userprofile%\.codex
   ```

   目录中可能存在多个文件，我们用到的只有三个，需要配置的只有两个：

   - `config.toml`：Codex 的核心配置文件，中转服务与 MCP 等都在此文件配置。
   - `auth.json`：用来配置你在中转站获取的 API Key 秘钥。
   - `AGENTS.md`：用来设置 Codex 全局工作的提示词。

   ::: important 文件不存在请手动创建
   很多人刚安装可能没有这三个文件，你需要手动创建后再写入内容。
   :::

2. 配置 `config.toml`，将以下内容保存：

   ```toml
   disable_response_storage = true
   model = "gpt-5.2"
   model_provider = "freecode"
   model_reasoning_effort = "xhigh"
   model_verbosity = "high"

   [features]
   web_search_request = true

   [model_providers.freecode]
   base_url = "https://freecode.codes/v1"
   name = "freecode"
   requires_openai_auth = true
   wire_api = "responses"
   ```

3. 配置 API Key，将以下内容写入 `auth.json`：

   ```json
   {
     "OPENAI_API_KEY": "xxx"
   }
   ```

4. 回顾 [创建 API 令牌](../register/4-token.md)，在 FreeCode 中创建 Codex 分组的令牌，复制后将 key 填入 `xxx` 部分并保存。

5. 测试对话，在终端输入以下命令，出现交互界面后进行对话测试，有回复则配置成功：

   ```bash
   codex
   ```

## macOS

1. 在访达界面按下 `Command + Shift + G`，输入以下路径并回车，打开 Codex 配置目录：

   ```text
   ~/.codex
   ```

   同样需要配置 `config.toml` 与 `auth.json` 两个文件，初次安装若未自动生成，需要手动创建。

2. 将以下内容保存到 `config.toml`：

   ```toml
   model_provider = "freecode"
   model = "gpt-5.1-codex"
   model_reasoning_effort = "high"
   network_access = "enabled"
   disable_response_storage = true
   windows_wsl_setup_acknowledged = true
   model_verbosity = "high"

   [model_providers.freecode]
   name = "freecode"
   base_url = "https://freecode.codes/v1"
   wire_api = "responses"
   requires_openai_auth = true
   ```

3. 将以下内容写入 `auth.json`：

   ```json
   {
     "OPENAI_API_KEY": "xxx"
   }
   ```

4. 回顾 [创建 API 令牌](../register/4-token.md)，在 FreeCode 中创建 Codex 分组的令牌，复制后将 key 填入 `xxx` 并保存。

5. 在终端执行 `codex`，出现对话界面并能收到回复即表示配置成功。
