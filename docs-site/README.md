# FreeCode 文档站

基于 VuePress 2 + Theme Hope 的 FreeCode 使用文档站，内容覆盖普通用户教程、API 接入、CLI 配置、第三方客户端、部署运维、管理后台、支付和合规说明。

## 环境要求

- Node.js 18+（已在 Node 24 验证）
- pnpm（可用 `corepack enable` 启用）

## 常用命令

```bash
pnpm install        # 安装依赖
pnpm docs:dev       # 本地开发预览（默认 http://127.0.0.1:8080）
pnpm docs:build     # 构建静态产物到 src/.vuepress/dist
pnpm docs:clean-dev # 清缓存后启动开发
```

## 目录结构

```text
src/
├── README.md                  # 站点首页
├── .vuepress/
│   ├── config.ts              # 主配置
│   ├── theme.ts               # Theme Hope 主题配置
│   ├── navbar.ts              # 顶部导航
│   ├── sidebar.ts             # 侧边栏
│   └── public/                # 静态资源
docs/
├── register/   # 快速开始
├── user/       # 用户中心、API 令牌、用量、订单、渠道
├── token/      # 模型分组介绍
├── api/        # API 接入与错误排查
├── ccswitch/   # CC-Switch 使用
├── cli/        # CLI 配置教程
├── paint/      # 绘图模型教程
├── advanced/   # 第三方接入
├── deploy/     # 部署运维
├── admin/      # 管理后台
├── faq/        # 常见问题
└── tos/        # 条款与政策
```

## 内容维护原则

- 面向公开用户的内容不要包含服务器私钥、真实 token、OAuth secret、支付密钥或内部账号凭据。
- API Key、令牌、密钥在用户文档中统一解释为同一类访问凭据，标题尽量使用“API 令牌”。
- 部署和管理文档以摘要和操作入口为主，敏感部署细节保留在仓库私有文档或服务器私有配置中。
- 修改文档后运行 `pnpm docs:build`，确认 VuePress 能成功生成静态产物。

## 部署

`pnpm docs:build` 后，将 `src/.vuepress/dist` 中的静态文件部署到文档站 Web 根目录。
