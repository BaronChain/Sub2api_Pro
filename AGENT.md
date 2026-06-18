# AGENT · 最高优先级索引目录

> 本文件是 **sub2api_Pro** 仓库的最高级索引与导航入口，用于本地**修复、检索、查阅**。
> 任何 AI / 开发者进入本仓库，**先读本文件**，再按索引跳转到目标位置。
> 维护原则：保持简短、只放索引与极高优先级约束；细节沉淀到对应专门文档，不在此处展开。

---

## 0. 仓库一句话定位

多上游 LLM 订阅转 API 的聚合网关：将 OpenAI / Claude / Gemini / Antigravity / Bedrock / Vertex 等上游账号池，统一暴露为兼容 OpenAI / Anthropic / Gemini 的 API，并提供计费、配额、支付、运维监控、用户与管理后台。

技术栈：**Go 1.26 (Ent ORM + Gin + Wire)** 后端 + **Vue3 + TS + Vite + Pinia (pnpm)** 前端；数据库 **PostgreSQL 16 + Redis**。

---

## 1. 极高优先级约束（违反会直接导致 CI 失败或线上事故）

> 这些是「红线」，优先级高于一般开发习惯。改动前后必须遵守。

1. **前端只用 `pnpm`，不用 `npm`**。改了 `package.json` 必须 `pnpm install` 并提交 `pnpm-lock.yaml`（CI 用 `--frozen-lockfile`）。
2. **Go 版本锁定 1.26.x**（CI 强校验），本地不要用其它大版本生成代码或 lint。
3. **改了 `backend/ent/schema/*.go` 必须 `go generate ./ent` 并提交生成代码**，否则改动不生效。
4. **改了 Go interface，必须补全所有 test stub / mock 的新方法**，否则编译失败（搜索 `type.*Stub.*struct` / `type.*Mock.*struct`）。
5. **依赖注入用 Wire**：改了 provider/构造函数后，需重新生成 `wire_gen.go`，不要手改生成文件。
6. **不提交密钥 / 凭据**：`.env`、`config.yaml`、bcrypt hash、OAuth secret 等不得入库。涉及账号凭据的代码改动要走脱敏路径（见 `account_credentials_redact.go`）。
7. **PR 合并前自检（缺一不可）**：
   - `cd backend && go test -tags=unit ./...`
   - `cd backend && go test -tags=integration ./...`
   - `cd backend && golangci-lint run ./...`（golangci-lint **v2.7**）
   - 前端：`pnpm install` + 必要时 `pnpm test` / `pnpm build`
8. **Windows 本地环境注意**（详见 DEV_GUIDE）：用 `127.0.0.1` 而非 `localhost`；PowerShell 中 bcrypt 的 `$` 需转义或走 `psql -f`；无 `make` 时直接用 Makefile 原始命令；psql 不支持中文路径。
9. **网关热路径（hotpath）代码改动需谨慎**：`gateway_*` / `openai_ws_*` / `scheduler_*` 涉及账号调度、计费、并发，改动必须跑对应单测与 benchmark，避免引入选不到账号 / 计费错误 / 串号。

---

## 2. 文档索引（先查文档，再读代码）

| 主题 | 文件 | 用途 |
|------|------|------|
| **本索引** | `AGENT.md` | 总导航 + 红线约束（你正在读） |
| 开发环境 / 坑点 | `DEV_GUIDE.md` | 本地环境、CI、11 个高频坑、常用命令速查 |
| 项目说明 | `README.md` / `README_JA.md` | 功能、特性、部署总览 |
| 部署指南 | `docs/部署指南.md`、`deploy/README.md`、`deploy/DOCKER.md` | 部署流程与 Docker |
| 部署记录 | `docs/部署记录.md` | 历史部署记录 |
| 本地开发与增量部署 | `docs/本地开发与增量部署指南.md` | 本地实时预览 + 一键推送服务器 |
| 支付集成 | `docs/PAYMENT.md` / `docs/PAYMENT_CN.md` | 支付总体设计 |
| 管理端支付 API | `docs/ADMIN_PAYMENT_INTEGRATION_API.md` | 后台支付接口 |
| 数据管理 | `deploy/DATAMANAGEMENTD_CN.md` | data management 守护进程 |
| 法务 | `docs/legal/`、`CLA.md`、`LICENSE` | 合规与许可 |
| 贡献规范 | `DEV_GUIDE.md` 第四、十一节 | PR 流程与检查清单 |

---

## 3. 后端代码索引 `backend/`

入口与装配：
- 主程序入口：`backend/cmd/server/main.go`
- 依赖注入：`backend/cmd/server/wire.go` + `wire_gen.go`（生成，勿手改）
- 版本号：`backend/cmd/server/VERSION`
- 工具命令：`backend/cmd/jwtgen/`（JWT 生成）

分层目录（`backend/internal/`）：

| 目录 | 职责 | 修复时优先看 |
|------|------|--------------|
| `config/` | 配置加载与校验 | `config.go`、`validate_dingtalk.go` |
| `handler/` | HTTP 处理器（网关 + 业务） | 见下「网关核心」 |
| `handler/admin/`、`handler/dto/`、`handler/quotaview/` | 后台 / DTO / 配额视图 | |
| `service/` | 业务逻辑（仓库最大模块） | 见下「服务核心」 |
| `repository/` | 数据访问层 | |
| `ent/` + `ent/schema/` | ORM 模型与生成代码 | 见下「数据模型」 |
| `migrations/` | 数据库迁移脚本 | |
| `server/routes/` | 路由注册 | `gateway.go`、`admin.go`、`auth.go`、`user.go`、`payment.go` |
| `server/middleware/`、`middleware/` | 中间件 | |
| `payment/provider/` | 支付渠道实现 | `alipay/wxpay/stripe/airwallex/easypay`，工厂 `factory.go` |
| `integration/`、`setup/` | 集成与初始化 | |
| `pkg/` | 上游协议适配与工具 | 见下「上游适配」 |
| `util/`、`web/`、`testutil/` | 工具 / 前端嵌入 / 测试辅助 | |

### 3.1 网关核心（请求转发热路径）`internal/handler/`
- 网关主入口：`gateway_handler.go`、`gateway_helper.go`、`endpoint.go`
- OpenAI 兼容：`openai_gateway_handler.go`、`openai_chat_completions.go`、`openai_embeddings.go`、`openai_images.go`、`openai_stream_validation.go`
- Gemini：`gemini_v1beta_handler.go`、`gemini_cli_session_test.go`
- 故障切换：`failover_loop.go`、`gateway_handler_stream_failover_test.go`
- 幂等 / 限流 / 限并发：`idempotency_helper.go`、`image_concurrency_limiter.go`、`request_body_limit.go`
- 错误处理：`ops_error_logger.go`、`stream_error_event.go`、`concurrency_error_response.go`

### 3.2 服务核心（`internal/service/`，按主题）
> 该目录文件极多，按前缀检索：
- 账号池 / 调度：`account_*.go`、`openai_account_scheduler*.go`、`scheduler_*.go`、`gateway_account_selection_test.go`
- 网关业务：`gateway_*.go`、`openai_gateway_*.go`、`antigravity_gateway_service.go`、`gemini_*compat*.go`
- OpenAI WebSocket：`openai_ws_*.go`（含 v2 `openai_ws_v2/`）
- 计费：`billing_*.go`、`pricing_service.go`、`model_pricing_resolver.go`、`usage_*.go`
- 配额 / 限流：`ratelimit_service*.go`、`quota_fetcher.go`、`*_quota*.go`、`rpm_cache.go`、`concurrency_service.go`
- 鉴权 / OAuth：`auth_*.go`、`oauth_*.go`、`*_oauth_service.go`、`totp_service.go`、`turnstile_service.go`
- Token 刷新与缓存：`token_*.go`、`*_token_provider.go`、`refresh_*.go`
- 支付：`payment_*.go`、`promo_*.go`、`redeem_*.go`、`affiliate_service.go`
- 运维监控（Ops）：`ops_*.go`、`dashboard_*.go`、`channel_monitor_*.go`
- 上游协议：`bedrock_*.go`、`vertex_service_account.go`、`gemini_*.go`、`openai_codex_*.go`、`antigravity_*.go`
- 通知 / 邮件：`email_*.go`、`notification_email_service.go`、`balance_notify_*.go`、`content_moderation*.go`
- 系统：`leader_lock.go`、`system_operation_lock_service.go`、`timing_wheel_service.go`、`backup_service.go`、`update_service.go`、`data_management_*.go`

### 3.3 上游协议适配 `internal/pkg/`
`openai/`、`openai_compat/`、`claude/`、`gemini/`、`geminicli/`、`googleapi/`、`antigravity/`、`apicompat/`、`oauth/`、`websearch/`、`tlsfingerprint/`、`httpclient/`、`proxyurl/`、`proxyutil/`、`usagestats/`、`pagination/`、`logger/`、`response/`、`errors/`、`ip/`、`timezone/`、`sysutil/`、`ctxkey/`

### 3.4 数据模型 `backend/ent/schema/`
用户域：`user.go`、`user_subscription.go`、`user_platform_quota.go`、`user_allowed_group.go`、`user_attribute_*.go`、`auth_identity*.go`
账号/渠道：`account.go`、`account_group.go`、`group.go`、`channel_monitor*.go`、`proxy.go`、`tls_fingerprint_profile.go`
计费/支付：`subscription_plan.go`、`payment_order.go`、`payment_provider_instance.go`、`payment_audit_log.go`、`promo_code*.go`、`redeem_code.go`
系统：`setting.go`、`announcement*.go`、`usage_log.go`、`usage_cleanup_task.go`、`idempotency_record.go`、`error_passthrough_rule.go`、`security_secret.go`、`pending_auth_session.go`

---

## 4. 前端代码索引 `frontend/src/`

| 目录 | 职责 |
|------|------|
| `api/` + `api/admin/` | 后端接口调用封装 |
| `views/admin/` | 管理后台页面（账号/渠道/用户/分组/订阅/运维 Ops/支付/兑换/公告/风控/备份） |
| `views/user/` | 用户端页面（仪表盘/密钥/用量/订阅/支付/邀请/渠道状态） |
| `views/auth/` | 登录注册 / 各 OAuth 回调（微信/钉钉/LinuxDo/OIDC/邮箱）|
| `views/public/`、`views/setup/` | 法务文档页 / 安装向导 |
| `components/` | 组件（account/admin/auth/channels/charts/keys/payment/user/common/Guide）|
| `stores/` | Pinia：`auth/app/payment/subscriptions/announcements/onboarding/adminSettings/adminCompliance` |
| `composables/`、`utils/`、`constants/`、`types/` | 组合函数 / 工具 / 常量 / 类型 |
| `i18n/locales/` | 国际化文案 |
| `router/` | 路由 |

入口：`frontend/index.html` → `src/main.ts`（如有）；构建配置 `vite.config.ts`、测试 `vitest.config.ts`。

---

## 5. 部署与运维索引 `deploy/`

- Docker：`Dockerfile`、`docker-compose.yml` / `.dev.yml` / `.local.yml` / `.standalone.yml`、`docker-entrypoint.sh`、`build_image.sh`、`docker-deploy.sh`
- 配置样例：`deploy/config.example.yaml`、`deploy/.env.example`
- 反向代理：`deploy/Caddyfile`
- systemd：`deploy/sub2api.service`、`deploy/sub2api-datamanagementd.service`、`install*.sh`
- 发布：根目录 `.goreleaser.yaml` / `.goreleaser.simple.yaml`、`Dockerfile.goreleaser`
- CI：`.github/workflows/`（`backend-ci.yml` / `security-scan.yml` / `release.yml`）
- 脚本/工具：`scripts/`、`tools/`（如 `tools/check_pnpm_audit_exceptions.py`）

---

## 6. 高频任务 → 快速入口

| 我要做… | 去哪里 |
|---------|--------|
| 配本地环境 / 踩坑排查 | `DEV_GUIDE.md` |
| 修网关转发 / 上游报错 | `internal/handler/gateway_*`、`internal/service/gateway_*` + `pkg/<上游>/` |
| 修账号选不到 / 调度 | `service/account_*`、`service/*scheduler*`、`service/ratelimit_service*` |
| 修计费 / 用量 | `service/billing_*`、`service/usage_*`、`service/pricing_service.go` |
| 加 / 改上游平台 | `pkg/<平台>/` + `service/<平台>_*` + `ent/schema/account.go` |
| 改数据库结构 | `ent/schema/*.go` → `go generate ./ent` → `migrations/` |
| 改支付渠道 | `payment/provider/` + `service/payment_*` + `docs/PAYMENT*.md` |
| 加路由 | `server/routes/*.go` |
| 改后台/用户页面 | `frontend/src/views/{admin,user}/` + 对应 `api/` + `stores/` |
| 部署 / Docker | `deploy/` + `docs/部署指南.md` |
| 本地预览 / 增量推送服务器 | `scripts/deploy-local/` + `docs/本地开发与增量部署指南.md` |
| 文档站（docs.freecode.codes） | `docs-site/`（VuePress 2 + Theme Hope） |

---

## 7. 服务器 SSH 操作（必读）

> **所有涉及 SSH / 服务器操作的任务，必须先读取 `SERVER_INFO.md` 获取连接信息（IP、端口、用户、私钥路径）。**
> 该文件已被 `.gitignore` 排除，不会入库。

- SSH 命令模板：`ssh -i ~/.ssh/id_ed25519 root@47.84.98.213 '<命令>'`
- 文档站静态文件部署在 `/var/www/docs.freecode.codes/`，反代由 Caddy 管理（`/etc/caddy/Caddyfile`）。
- 文档站更新流程：`cd docs-site && pnpm docs:build` → tar 管道上传到 `/var/www/docs.freecode.codes/`。

---

_本索引随结构演进更新；新增顶层模块或红线约束时，请同步修订本文件。_
