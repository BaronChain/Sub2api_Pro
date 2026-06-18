---
title: Claude Desktop
icon: desktop
order: 1
author: FreeCode Team
date: 2026-01-01
category:
  - 第三方接入
---

# Claude Desktop

## 软件下载

进入 Claude Desktop 下载页面，根据自己的系统下载对应的安装包。

## 软件安装

### Windows

Windows 系统下安装需要请求 Anthropic 官方，需要用代理挂 **全局服务（TUN 模式）**，或用命令行运行安装程序使其强制走代理，否则会安装失败。

确认你当前使用的代理端口号（例如 Clash Verge 默认 `7897`），在安装程序所在目录运行 cmd，分别输入以下命令运行安装程序：

```bat
set HTTP_PROXY=http://127.0.0.1:7897
set HTTPS_PROXY=http://127.0.0.1:7897
"Claude Setup.exe"
```

### macOS

macOS 系统下直接正常安装即可。

## 绕过登录并配置第三方接口

### 开启开发者模式

- **Windows**：鼠标点击邮件输入框获取焦点，键盘 `Tab` 跳到左上角菜单，按下回车，依次进入 `help → troubleshooting → enable developer mode`。
- **macOS**：直接在左上角菜单依次进入 `help → troubleshooting → enable developer mode`。

开启后等待软件重启。

### 配置第三方 API

同样的方法打开菜单，依次进入 `Developer → Configure third-party inference`：

1. 在 **Gateway base URL** 填入 `https://freecode.codes`。
2. 将 **Gateway auth scheme** 更改为 `x-api-key`。
3. **Gateway API key** 填入生成的 **CC 分组** 的 API Key。
4. 打开最下方 **Skip login-mode chooser** 选项。
5. 点击右下角 **Apply locally** 按钮使配置生效。

进行愉快的对话吧~
