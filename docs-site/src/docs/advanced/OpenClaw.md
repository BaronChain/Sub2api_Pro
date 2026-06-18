---
title: OpenClaw
icon: cat
order: 4
author: FreeCode Team
date: 2026-01-01
category:
  - 第三方接入
---

# OpenClaw

::: tip 提示
此教程适合 Linux 云服务器、macOS 系统用户。
:::

## 安装与初始化

登录服务器 SSH，或在 macOS 打开终端，输入以下命令开始安装，耐心等待安装流程结束：

```bash
curl -fsSL https://openclaw.ai/install.sh | bash
```

出现界面提示后，依次进行：

1. 选择 `yes` 回车确认。
2. 选择 `QuickStart` 回车确认。
3. 在选择供应商部分先选择 `Skip for now`，跳过设置。
4. 在适配器选择部分，选择 `anthropic`。
5. 在模型选择部分，选择 `opus-4.5`。
6. 社交软件适配器按个人需要选择（如 Telegram），输入 Bot Token 回车。
7. 安装 Skill 部分先跳过（后续可通过网页安装）。
8. Hook 部分使用空格键全选后回车，等待 Gateway 安装完成。
9. 安装 Shell 补全脚本选择 `yes` 回车，完成安装。

## 渠道与模型配置

在 SSH 控制台或 macOS 终端中新增供应商：

1. 选择 **添加供应商**，新增自定义供应商，端点填写 `https://freecode.codes`。
2. 在获取的模型中选择你要使用的模型（以 `Claude Opus 4.5` 为例）。
3. 输入 [创建 API 令牌](../register/4-token.md) 中创建的相关分组令牌，复制后填入。
4. 选择 **选择模型**，选中刚才配置的模型回车确认，然后退出回到控制台。

重启 Gateway：

```bash
openclaw gateway restart
```

重启成功后，进入 TUI 界面测试模型是否能正常输出，正常则输入 `/quit` 退出：

```bash
openclaw tui
```

## 浏览器访问 Dashboard

在控制台输入命令获取 Dashboard URL，在浏览器访问。

::: tip 服务器部署提示
如果你在服务器运行，请使用 Nginx 或其他反向代理工具反代服务并设置 SSL 证书。同时修改 `~/.openclaw` 下的 `openclaw.json` 文件，在 `gateway` 字段下添加：

```json
"controlUi": {
  "allowInsecureAuth": true
}
```

修改后重启网关：`openclaw gateway restart`。
:::

## 配置 Telegram Bot 访问权限

回到创建机器人的 `@BotFather` 对话，点击机器人链接进行对话，首次对话拿到所需的 **Pairing code**，然后在控制台输入：

```bash
openclaw pairing approve telegram 你的Pairing_code
```
