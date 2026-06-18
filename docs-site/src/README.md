---
home: true
icon: house
title: 主页
heroImage: /logo.png
heroText: FreeCode 使用文档
tagline: 从注册、API 令牌、CLI 接入到部署运维，系统介绍 FreeCode 中转服务的使用与管理方法
actions:
  - text: 快速开始
    icon: rocket
    link: /docs/register/1-register.html
    type: primary

  - text: API 接入
    icon: code
    link: /docs/api/

  - text: 文档目录
    link: /docs/

highlights:
  - header: 新用户从这里开始
    description: 按注册、登录、购买额度、创建令牌、配置 CLI 的顺序操作，可以最快跑通 Claude Code、Codex、Gemini 等常见工具。
    features:
      - title: 快速开始
        icon: rocket
        details: 注册、登录、充值、创建令牌、配置 CLI，按顺序完成基础接入
        link: /docs/register/1-register.html

      - title: API 令牌设置
        icon: key
        details: 了解分组、额度、IP 白名单、模型限制、周期限额和安全注意事项
        link: /docs/user/api-key.html

      - title: CLI 配置教程
        icon: terminal
        details: Claude Code / Codex / Gemini CLI 手动配置详解
        link: /docs/cli/2-claude.html

      - title: CC-Switch 一键切换
        icon: shuffle
        details: 用图形化或 CLI 工具快速切换不同中转配置
        link: /docs/ccswitch/1-common.html

  - header: API、排障与管理
    description: 除了基础教程，文档站也整理了 API 接入、用量计费、管理员配置、部署运维和合规说明。
    features:
      - title: API 接入
        icon: code
        details: OpenAI、Anthropic、Gemini 与图片接口的 Base URL、认证方式和常见错误
        link: /docs/api/

      - title: 用量与计费排查
        icon: chart-line
        details: 解释余额、令牌额度、分组倍率、周期限制和异常消耗排查
        link: /docs/user/usage-billing.html

      - title: 管理后台
        icon: user-shield
        details: 管理用户、账号池、分组、支付、风控、Ops 和渠道监控
        link: /docs/admin/

      - title: 部署运维
        icon: server
        details: 部署架构、反向代理、文档站发布、本地开发和运维检查清单
        link: /docs/deploy/

copyright: false
footer: FreeCode | Copyright © 2026 FreeCode Team
---

## FreeCode 文档站

本 文档站面向三类读者：

- **普通用户**：注册账号、购买额度、创建 API 令牌、配置 CLI 或第三方客户端。
- **开发者**：通过 OpenAI / Anthropic / Gemini 兼容接口接入 FreeCode。
- **管理员与部署者**：维护账号池、分组、支付、风控、监控和文档站部署。

如果你是第一次使用，建议从 [快速开始](/docs/register/1-register.html) 按顺序阅读。
