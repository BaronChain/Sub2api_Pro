---
title: Claude Code
icon: robot
order: 1
author: FreeCode Team
date: 2026-01-01
category:
  - 常见问题
---

# Claude Code 相关问题

## Claude Code 无法连接到 Anthropic 服务

使用 npm 安装完 `claude` 后，在命令行输入 `claude` 可能报错：

```text
Unable to connect to Anthropic services
Failed to connect to api.anthropic.com: ERR BAD REQUEST
```

或在初次配置时停在安装确认流程。解决方法：

### Windows

按 `Win + R`，输入 `cmd` 回车，在命令行中运行：

```bat
powershell -Command "$f='%USERPROFILE%\.claude.json';$j=Get-Content $f|ConvertFrom-Json;$j|Add-Member -NotePropertyName 'hasCompletedOnboarding' -NotePropertyValue $true -Force;$j|ConvertTo-Json|Set-Content $f"
```

然后重启 Claude CLI。

### macOS

打开终端运行：

```bash
jq '. + {"hasCompletedOnboarding": true}' ~/.claude.json > /tmp/tmp.json && mv /tmp/tmp.json ~/.claude.json
```

::: tip 提示
如果提示未找到 `jq`，可以输入 `brew install jq` 进行安装。
:::

然后重启 Claude CLI。

## 如何在 VSCode CC 插件中使用 FreeCode

确保你已完成 [环境检查](../cli/1-env.md) 并保证 claude code cli 正常。

### Windows

按 `Win + R`，输入 `%userprofile%\.claude` 回车打开配置目录。如果目录中没有 `config.json`，手动创建后写入以下内容保存：

```json
{
  "primaryApiKey": "FreeCode"
}
```

重启 VSCode 即可使用。

### macOS

在访达按 `Command + Shift + G`，输入 `~/.claude` 回车打开配置目录。如果没有 `config.json`，手动创建后写入相同内容并重启 VSCode。

## 切换回 200K 上下文并禁用非必要流量

如果你希望将 Claude Code 从 1M 上下文切换回 200K，并关闭非必要请求与终端标题变更，可以在 `settings.json` 的 `env` 中加入：

```json
{
  "env": {
    "CLAUDE_CODE_DISABLE_1M_CONTEXT": "1",
    "CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC": "1",
    "CLAUDE_CODE_DISABLE_TERMINAL_TITLE": "1"
  }
}
```

- Windows：`Win + R` → `%userprofile%\.claude` → 打开/创建 `settings.json`。
- macOS：`Command + Shift + G` → `~/.claude` → 打开/创建 `settings.json`。

## Claude Code 常用命令

| 命令 | 功能说明 |
| --- | --- |
| `claude` | 在当前目录启动交互式 REPL |
| `claude "解释这个项目"` | 启动 REPL 并带上初始问题 |
| `claude -p "解释这个函数"` | print 模式一次性问答，便于脚本调用 |
| `claude -c` | 继续当前目录最近的一次会话 |
| `claude -r "abc123" "把这个 PR 完成"` | 通过会话 ID 恢复指定会话 |
| `claude update` | 更新 Claude Code CLI 到最新版本 |
| `claude mcp` | 管理和配置 MCP 服务器 |
| `claude --model sonnet` | 指定会话使用的模型 |
| `claude --resume abc123` | 通过会话 ID 恢复会话 |
| `claude --dangerously-skip-permissions` | 跳过权限确认（高风险，仅在信任环境使用） |
