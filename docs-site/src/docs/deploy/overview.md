---
title: 部署总览
icon: diagram-project
order: 1
author: FreeCode Team
---

# 部署总览

FreeCode / Sub2API 是一个多上游 LLM 订阅转 API 的聚合网关。典型部署包含：

- 后端服务：Go + Gin + Ent；
- 前端管理与用户控制台：Vue3；
- 数据库：PostgreSQL；
- 缓存与限流：Redis；
- 反向代理：Caddy、Nginx 或其他网关；
- 可选：支付回调、邮件通知、监控告警、备份任务。

## 部署前准备

建议准备：

- 一台 Linux 服务器；
- 可访问的域名和 HTTPS 证书；
- PostgreSQL 15/16；
- Redis 7；
- 用于后台管理员的初始账号；
- 明确的上游账号、分组和计费策略。

## 常见部署方式

| 方式 | 适合场景 |
|------|----------|
| 脚本安装 | 单机快速部署，适合首次搭建 |
| Docker / Compose | 需要容器化、易迁移和统一依赖 |
| 手动二进制部署 | 需要自定义 systemd、目录和反代规则 |
| 自行构建镜像 | Fork 后有定制代码，需要自己的发布流程 |

详细命令请以仓库内 `docs/部署指南.md`、`deploy/README.md` 和 `deploy/DOCKER.md` 为准。

## 反向代理注意事项

如果使用 Nginx 且需要兼容 Codex / sticky session 场景，请确认配置中允许带下划线的请求头：

```nginx
underscores_in_headers on;
```

同时建议：

- 提高长连接和上游超时时间，避免图片生成或长上下文请求被提前断开；
- 正确转发 `Authorization`、`x-api-key` 等认证头；
- 保留真实客户端 IP，以便审计、限流和 IP 白名单生效；
- 对管理后台开启额外访问控制或强密码策略。

## 上线后检查

1. 打开首页和控制台；
2. 创建管理员并登录后台；
3. 检查数据库和 Redis 连接；
4. 创建测试分组和测试 API 令牌；
5. 调用 `/v1/models` 或客户端测试；
6. 检查用量记录是否写入；
7. 如果启用支付，完成一笔小额测试订单；
8. 配置备份和日志清理策略。
