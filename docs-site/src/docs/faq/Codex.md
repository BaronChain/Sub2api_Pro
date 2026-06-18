---
title: Codex
icon: robot
order: 2
author: FreeCode Team
date: 2026-01-01
category:
  - 常见问题
---

# Codex 相关问题

## 如何更高效地使用 Codex

很多人使用一段时间后认为模型出现“降智”。其实关键在于如何合理使用：

- **任务划分**：不要提交非常笼统的任务（如“请帮我写一个管理系统后台”），Codex 的特点是严谨有序、指哪打哪，需要你把任务拆分细致。
- **掌控之内**：开始任务前先评估是否拆分得足够细、是否符合“模块化”准则，提交前应能预估改动会修改哪些文件。
- **避免压缩**：多数场景下任务最多用约 60% 上下文就能解决。如果超过且需要压缩，说明拆分工作没做好，需要更精细地拆分。

## Windows 系统下流畅使用 Codex

::: important 此方法同时解决读写文件、乱码、Token 耗费高、项目无记忆等多个痛点
:::

确保 Codex CLI 与 VSCode Codex 插件正常运行后，按 `Win + R` 输入 `%userprofile%\.codex` 打开用户目录，编辑 `config.toml`：

```toml
model_provider = "freecode"
model = "gpt-5.4"
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

打开目录下 `AGENTS.md`（没有则手动创建），写入全局工作指南，例如：

```markdown
# Codex 全局工作指南

## 回答风格:
 - 回答必须使用中文
 - 对总结、Plan、Task 等长内容，优先逻辑整理后使用美观的 Table 格式输出；普通内容正常输出
```

## Codex 常用命令

| 命令 | 说明 |
| --- | --- |
| `/model` | 选择当前使用的模型 |
| `/approvals` | 设置本会话的审批规则 |
| `/review` | 让 Codex 审查当前工作区变更 |
| `/resume` | 从历史会话列表选择并继续 |
| `/new` | 开启新对话 |
| `/init` | 在当前目录生成 AGENTS.md 模板 |
| `/compact` | 总结对话内容以释放上下文 |
| `/diff` | 查看当前 git diff |
| `/status` | 查看会话配置和 token 使用情况 |
| `/mcp` | 列出当前可用的 MCP 工具 |
| `/exit` | 退出 Codex CLI |

## Codex 在 Windows 下乱码问题

按 `Win + R` 输入 `intl.cpl` 回车，点击上侧 **管理** 选项卡，再点击 **更改系统区域设置**，勾选“使用 Unicode UTF-8 提供全球语言支持”，确定后重启电脑再使用 codex 即可避免乱码。

## Codex 开启内置网络搜索

打开 `config.toml`，加入以下内容后运行 Codex：

```toml
[features]
web_search_request = true
```

## Connection failed 报错

报错信息类似 `Connection failed: error sending request for url`。出现这种情况通常是本机网络问题：

1. 检查本机网络是否通畅，能否访问其他页面。
2. 检查电脑是否使用了网络代理（梯子），如果有请关闭。
3. 在终端运行 `codex` 测试，判断是否是 VSCode 插件问题，如是请重启 VSCode。
4. 若仍不行，带上报错截图联系客服或群友。

## 401 报错

报错信息类似 `exceeded retry limit, last status: 401 Unauthorized`。通常是环境变量覆盖了配置。

检查并清除环境变量：

Windows：

```bat
cmd /c "setx OPENAI_API_KEY \"\" & setx OPENAI_BASE_URL \"\""
```

macOS：

```bash
unset OPENAI_API_KEY OPENAI_BASE_URL
```

然后检查 `~/.codex/auth.json` 中的 API Key 与 `~/.codex/config.toml` 中的请求地址是否正确。

## 403 报错

报错信息类似 `403 Forbidden: {"error":{"message":"Usage not included in your plan"...}}`。这是号池中的账号出现问题：

1. 使用 `Ctrl + C` 打断对话（VSCode 中点击停止按钮）。
2. 重新发起对话，观察是否再次出现。
3. 如果重试 3 次以上无效，带上报错截图联系客服或群友。
