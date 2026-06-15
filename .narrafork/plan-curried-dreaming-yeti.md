# Sub2API 伪装性审计报告

> 目标：从"上游官方 API（Anthropic/OpenAI/Google）能否识别此请求来自中转反代"的角度审查代码

---

## 风险分级总览

| 等级 | 数量 | 含义 |
|------|------|------|
| 🔴 高 | 2 | 上游可直接通过此类特征判定请求来自中转 |
| 🟡 中 | 6 | 结合其他信号可推断，但单独不构成铁证（含原 H2 IP 传递，经核实降级） |
| 🟢 低 | 4 | 理论可检测，实际几乎不会被用于判定 |
| ✅ 已做好 | 6 | 项目已实施的伪装措施（正面评价） |

---

## 一、🔴 高风险（应优先修改）

### H1. HTTP 400 错误体原样透传

**文件**: `backend/internal/service/gateway_service.go:7417-7419`

```go
case 400:
    c.Data(http.StatusBadRequest, "application/json", body)  // ← 上游原始错误体直接返回
```

**风险**: 上游 400 响应体包含平台特有格式和措辞（如 Anthropic 的 `{"type":"error","error":{"type":"invalid_request_error",...}}`），直接透传等于把上游 API 的"指纹"原封不动告诉客户端。虽然这不直接暴露给上游，但如果官方做 entrapment（给特定账号投递带水印的 400 错误），可以追溯中转链路。

**建议**: 对 400 错误体做结构化提取和重新封装，使用 Anthropic 标准格式重新构建错误响应，而非原样透传。

---

### H2. 客户端 IP 通过 X-Forwarded-For / X-Real-IP 传递给上游

**文件**: `backend/internal/pkg/ip/ip.go:14-29`

```go
// 2. Nginx X-Real-IP
if ip := c.GetHeader("X-Real-IP"); ip != "" { ... }
// 3. X-Forwarded-For
if xff := c.GetHeader("X-Forwarded-For"); xff != "" { ... }
```

**Anthropic 路径**: `gateway_service.go:6375-6384` — 请求头白名单透传中 **不含** X-Forwarded-For / X-Real-IP（✅ 安全）

**OpenAI 路径**: `openai_gateway_service.go:110-111` — `codexCLIOnlyDebugHeaderWhitelist` 包含 X-Forwarded-For 和 X-Real-IP，但这**仅用于诊断日志**，不参与上游透传（✅ 安全）

**实际风险点**: 项目本身**没有**主动将客户端 IP 设置到转发给上游的请求中。但如果部署在 Nginx/Caddy 反代后面，且 Nginx 配置了 `proxy_set_header X-Forwarded-For`，Go 的 `net/http` 会将这些头从入站请求复制到出站请求吗？**不会**，Go HTTP 客户端不会自动转发入站头。项目代码中 `buildUpstreamRequest` 是构建全新 `http.NewRequest`，仅复制白名单头。

**结论**: 此项实际**不构成直接风险**（项目白名单已正确排除 IP 头），但需要确保部署时上游反代不会添加此类头。降级为 🟡 中。

---

### H3. OpenAI SSE 错误事件使用非原生类型 `"upstream_error"`

**文件**: `backend/internal/service/openai_gateway_service.go:4636`

```go
payload := `{"type":"error","sequence_number":0,"error":{"type":"upstream_error","message":...`
```

**风险**: 原生 OpenAI Responses API 的错误类型有 `rate_limit_error`、`server_error`、`invalid_request_error` 等，**从未使用** `upstream_error`。如果客户端将此错误类型转发给 OpenAI（如 telemetry），或者官方检查客户端日志中的错误类型模式，`upstream_error` 是中转服务的明确标志。

**建议**: 将 `"upstream_error"` 替换为原生 OpenAI 错误类型（如 `"server_error"` 或 `"rate_limit_error"`），保持与上游一致。

---

## 二、🟡 中风险（建议修改）

### M1. `x-codex-*` 响应头强制放行

**文件**: `backend/internal/service/openai_gateway_service.go:4079-4112`

```go
for _, rawKey := range []string{
    "x-codex-primary-used-percent",
    "x-codex-secondary-used-percent",
    ...
} {
    // 强制透传这些头
}
```

**风险**: 这些头是 OpenAI Codex 专有的速率限制头。虽然白名单过滤了其他所有头，但这里额外放行了 7 个 x-codex-* 头。对于客户端使用来说这是功能必需（显示用量），但**如果官方检测代理返回的响应头集合**，x-codex-* 的大量存在可能被视为间接证据。不过这些头原生 Codex 客户端也会收到，**风险有限**。

---

### M2. 前端 HTML title 直接暴露项目名

**文件**: `frontend/index.html:7`

```html
<title>Sub2API - AI API Gateway</title>
```

**风险**: 虽然这个页面不会发给上游 API，但**你的中转站对客户端可见**。如果用户把中转站 URL 分享到公开平台，搜索引擎或爬虫抓取到 `<title>Sub2API</title>` 就直接暴露了使用的程序。

**建议**: 改为自定义名称。

---

### M3. SSE keepalive ping 行为

**文件**: `backend/internal/service/gateway_service.go:7681-7694`

```go
keepaliveInterval := ...
// 下游 keepalive：防止代理/Cloudflare Tunnel 因连接空闲而断开
```

**风险**: 真实的 Anthropic/OpenAI API **不会**在等待首个 token 时发送 `data: {"type": "ping"}` 类型的 keepalive。如果客户端将这些 ping 的时序/格式记录下来并上报，可能暴露中转行为。但**默认未启用**（需配置 `stream_keepalive_interval`），风险可控。

---

### M4. ChatCompletions/Responses 路径的上游错误消息

**文件**: 
- `gateway_forward_as_chat_completions.go:177`
- `gateway_forward_as_responses.go:175`

```go
writeGatewayCCError(c, mapUpstreamStatusCode(resp.StatusCode), "server_error", upstreamMsg)
```

**风险**: `upstreamMsg` 经过 `sanitizeUpstreamErrorMessage` 处理，但该函数仅做 URL 参数脱敏（替换 `key=xxx` → `key=***`），**不会**清理上游特有措辞。例如 Anthropic 的 `"Your API key is invalid"` 或 OpenAI 的 `"You exceeded your current quota"` 会原样传递。

---

### M5. 项目元数据指纹（go.mod 模块路径、二进制名等）

| 来源 | 暴露内容 |
|------|---------|
| `backend/go.mod` | `module github.com/Wei-Shaw/sub2api` |
| Dockerfile | `LABEL ... "Sub2API - AI API Gateway Platform"` |
| Dockerfile | `LABEL org.opencontainers.image.source="https://github.com/Wei-Shaw/sub2api"` |
| Dockerfile | 二进制路径 `/app/sub2api` |
| 前端 | `siteName = 'Sub2API'`、i18n 多处 |
| 配置 | `LOG_SERVICE_NAME=sub2api`、Redis key prefix `sub2api:` |

**风险**: 这些信息**不直接被上游 API 看到**（上游只看到 HTTP 请求），但对于服务器安全审计、反指纹有影响。如果你不想让任何人知道你用的是 sub2api，这些都需要改。

---

## 三、🟢 低风险（了解即可）

### L1. Mock 拦截响应使用固定消息 ID

**文件**: `backend/internal/handler/gateway_handler.go:1920-1921`

```go
msgID = "msg_mock_warmup"        // 固定值
msgID = "msg_mock_suggestion"    // 固定值
```

**说明**: 这仅用于拦截预热/建议请求（不发给上游），不会被上游看到。但客户端如果分析响应 ID 格式，可以发现 `msg_mock_*` 不是真实的 Anthropic 格式。已有 `generateRealisticMsgID()` 用于 haiku 探测响应，建议统一。

### L2. SSE ping 格式差异

**文件**: `backend/internal/service/gateway_helper.go:106-115`

不同平台使用不同的 ping 格式：
- Anthropic: `data: {"type": "ping"}\n\n`
- OpenAI: SSE 注释 `:\n\n`

**说明**: 这与各平台原生行为一致，**不应修改**。

### L3. x-request-id 透传

**说明**: 透传上游的 x-request-id 给客户端。上游的 request-id 格式是标准的，不构成明显指纹。风险极低。

### L4. X-Accel-Buffering: no 头

**说明**: 所有 SSE 路径设置 `X-Accel-Buffering: no`。这是 Nginx 专有头，暴露后端使用 Nginx，但不暴露"中转"身份。功能必需。

---

## 四、✅ 已做好的伪装措施（正面评价）

| 措施 | 文件/位置 | 说明 |
|------|-----------|------|
| **响应头白名单过滤** | `responseheaders.go` | 默认启用，仅允许 16 个头透传，阻断 Server/Via/CF-*/X-Powered-By 等 |
| **错误码通用化** | `gateway_service.go:7428-7451` | 401/403/429/529/5xx 全部映射为通用消息 |
| **Panic 不泄露堆栈** | `recovery.go` | 只返回 `"internal error"` |
| **Claude Code 指纹注入** | `gateway_service.go:6301-6407` | OAuth 账号自动注入 x-stainless-*、User-Agent、metadata.user_id 等完整 CC 指纹 |
| **TLS 指纹伪装** | `req_client_pool.go` + `tlsfingerprint` 包 | 使用 `ImpersonateChrome()` 伪装 TLS 握手 |
| **模型名替换** | 流式/非流式均支持 | 请求和响应中的模型名自动映射 |

---

## 五、推荐修改优先级

按**投入产出比**排序（最值得先做的在前）：

| 序号 | 修改项 | 风险等级 | 改动范围 | 说明 |
|------|--------|---------|---------|------|
| 1 | OpenAI SSE error 类型伪装 | 🔴 高 | 1 个文件 | `"upstream_error"` → 原生类型，小改动大收益 |
| 2 | 400 错误体结构化重封装 | 🔴 高 | 1 个文件 | 阻止上游水印追溯 |
| 3 | 前端 title/i18n 消品牌化 | 🟡 中 | 少量前端文件 | 防止站名暴露 |
| 4 | 上游错误消息深度脱敏 | 🟡 中 | 1-2 个文件 | 清理上游措辞特征 |
| 5 | keepalive ping 格式优化 | 🟡 中 | 1 个文件 | 让 ping 行为更贴近原生 |
| 6 | x-codex-* 头评估 | 🟡 中 | 1 个文件 | 可选：是否保留功能性透传 |
| 7 | 二进制/模块/容器消品牌化 | 🟡 中 | 多文件 | 全方位去除 sub2api 痕迹 |

**注**: 第 1-2 项是"上游反检测"的核心修改；第 3-7 项更多是"运营消品牌"层面。如果你只关心上游不被检测到，重点做 1-2 即可。如果你想彻底去除 sub2api 痕迹（让用户也不知道用的是 sub2api），则 3-7 也需要做。
