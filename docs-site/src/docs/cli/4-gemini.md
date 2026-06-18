---
title: Gemini 配置
icon: robot
order: 4
author: FreeCode Team
date: 2026-01-01
category:
  - CLI 配置教程
---

# Gemini 配置

## Windows

1. 键盘按下 `Win + R`，输入以下内容后回车，打开 Gemini CLI 配置目录：

   ```text
   %userprofile%\.gemini
   ```

2. 如果目录下没有 `.env` 文件，请新建一个 `.env`。`.env` 是 Gemini CLI 的配置文件，主要设置自定义端点、API Key 与所用模型。写入以下内容：

   ```ini
   GOOGLE_GEMINI_BASE_URL=https://freecode.codes
   GEMINI_API_KEY=xxx
   GEMINI_MODEL=gemini-2.5-pro
   ```

3. 回顾 [创建 API 令牌](../register/4-token.md)，在 FreeCode 中创建 Gemini 分组令牌，复制 API Key 后填入 `xxx`。

4. 打开终端执行 `gemini`，看到交互界面并能正常回复即表示配置成功。

## macOS

1. 在访达界面按下 `Command + Shift + G`，输入以下路径后回车，打开配置目录：

   ```text
   ~/.gemini
   ```

2. 若目录中没有 `.env` 文件，创建并写入以下内容：

   ```ini
   GOOGLE_GEMINI_BASE_URL=https://freecode.codes
   GEMINI_API_KEY=xxx
   GEMINI_MODEL=gemini-2.5-pro
   ```

3. 回顾 [创建 API 令牌](../register/4-token.md)，在 FreeCode 中创建 Gemini 分组令牌，填入 `xxx`。

4. 在终端运行 `gemini`，可正常进入对话并收到回复即表示配置完成。
