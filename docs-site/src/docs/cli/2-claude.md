---
title: Claude Code 配置
icon: robot
order: 2
author: FreeCode Team
date: 2026-01-01
category:
  - CLI 配置教程
---

# Claude Code 配置

## Windows

1. 键盘按下 `Win + R`，输入以下内容后回车，打开 Claude Code 配置目录：

   ```text
   %userprofile%\.claude
   ```

2. 如果目录中没有 `settings.json`，你需要手动创建。`settings.json` 是 Claude 的主要配置文件，用来配置中转站地址、API Key，以及一些 hooks、plugins 等。

3. 将以下内容写入 `settings.json`：

   ```json
   {
     "env": {
       "ANTHROPIC_BASE_URL": "https://freecode.codes",
       "ANTHROPIC_AUTH_TOKEN": "xxx",
       "CLAUDE_CODE_ATTRIBUTION_HEADER": "0",
       "CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC": "1",
       "CLAUDE_CODE_DISABLE_TERMINAL_TITLE": "1"
     }
   }
   ```

4. 回顾 [创建 API 令牌](../register/4-token.md)，在 FreeCode 中创建 CC 分组的令牌，替换上方 `xxx` 部分。

5. 在 Windows 终端运行 `claude`，出现对话界面后进行对话测试，能收到回复即表示配置成功。

## macOS

1. 在访达界面按下 `Command + Shift + G`，输入以下路径后回车，打开配置目录：

   ```text
   ~/.claude
   ```

2. 若目录不存在 `settings.json`，需要你手动创建。

3. 将以下内容写入 `settings.json`：

   ```json
   {
     "env": {
       "ANTHROPIC_BASE_URL": "https://freecode.codes",
       "ANTHROPIC_AUTH_TOKEN": "xxx",
       "CLAUDE_CODE_ATTRIBUTION_HEADER": "0",
       "CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC": "1",
       "CLAUDE_CODE_DISABLE_TERMINAL_TITLE": "1"
     }
   }
   ```

4. 回顾 [创建 API 令牌](../register/4-token.md)，在 FreeCode 中创建 CC 分组的令牌，替换上方 `xxx`。

5. 在终端运行 `claude`，看到对话界面并能正常回复即表示配置完成。

::: important 配置后仍报错？
如果配置完仍然提示需要登录或无法连接，请参考 [常见问题 - Claude Code](../faq/CC.md) 解决。
:::
