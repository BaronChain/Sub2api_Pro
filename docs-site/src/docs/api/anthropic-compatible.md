---
title: Anthropic / Claude 兼容接口
icon: robot
order: 2
author: FreeCode Team
---

# Anthropic / Claude 兼容接口

Claude Code、Claude Desktop 以及部分支持 Anthropic API 的客户端，可以通过 FreeCode 的 Claude 兼容入口接入。

## Base URL

```text
https://freecode.codes
```

Claude Code 配置时通常设置环境变量或 `settings.json`：

```json
{
  "env": {
    "ANTHROPIC_BASE_URL": "https://freecode.codes",
    "ANTHROPIC_AUTH_TOKEN": "你的CC分组API令牌"
  }
}
```

完整步骤请看 [Claude Code 配置](../cli/2-claude.md)。

## 认证方式

不同客户端字段名称可能不同：

| 客户端 | 常见字段 |
|--------|----------|
| Claude Code | `ANTHROPIC_AUTH_TOKEN` |
| Claude Desktop Gateway | `Gateway API key`，认证方式选择 `x-api-key` |
| 通用 Anthropic SDK | `x-api-key` 或客户端封装的 apiKey 参数 |

请使用 CC / Claude 相关分组令牌。

## 常见问题

### 连接到官方 Anthropic 而不是 FreeCode

说明环境变量或配置文件没有生效。请检查：

- 是否存在全局 `ANTHROPIC_BASE_URL` 覆盖；
- Claude Code 是否读取了正确目录下的 `settings.json`；
- 修改配置后是否重启终端或客户端。

### 第三方客户端测试失败，但 Claude Code 可用

某些 CC 分组只面向 Claude Code 场景，不一定支持第三方客户端完整测试。请以站点分组说明和实际客户端支持为准。

### 401 / 403

- 401 通常是 Key 错误、未启用或认证头错误；
- 403 通常是分组、模型、策略或账户权限不匹配。
