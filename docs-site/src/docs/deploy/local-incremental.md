---
title: 本地开发与增量部署
icon: code-branch
order: 2
author: FreeCode Team
---

# 本地开发与增量部署

本页面向维护者，说明从本地修改到服务器更新的基本流程。具体脚本和服务器信息请以仓库内维护文档和私有配置为准，不应写入公开文档站。

## 本地开发建议

- 前端包管理统一使用 `pnpm`；
- 修改前端依赖后同步提交 `pnpm-lock.yaml`；
- 后端修改 Go interface 后补齐测试 stub / mock；
- 修改 Ent schema 后运行代码生成；
- 修改 Wire provider 后重新生成 `wire_gen.go`；
- 不要提交 `.env`、真实 `config.yaml`、OAuth secret、支付密钥或服务器私钥。

## 推荐验证

后端：

```bash
cd backend
go test -tags=unit ./...
go test -tags=integration ./...
golangci-lint run ./...
```

前端或文档站：

```bash
cd frontend
pnpm install
pnpm build

cd ../docs-site
pnpm install
pnpm docs:build
```

## 增量部署思路

一般流程：

1. 本地完成修改并通过必要测试；
2. 构建前端或文档站静态产物；
3. 将产物上传到服务器目标目录；
4. 如后端有变更，重启服务并查看日志；
5. 访问线上页面和关键 API 做冒烟验证。

## 文档站部署

文档站构建产物位于：

```text
docs-site/src/.vuepress/dist
```

部署时只需要把该目录内的静态文件同步到文档站 Web 根目录。不要把 `node_modules`、源码中的私有配置或临时文件上传到公开目录。
