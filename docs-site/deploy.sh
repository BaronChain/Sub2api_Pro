#!/usr/bin/env bash
# =============================================================================
# 文档站一键发布脚本
# 用法: cd docs-site && bash deploy.sh
# 功能: 本地构建 → 上传服务器 → 验证
# =============================================================================
set -euo pipefail

# ── 配置 ──────────────────────────────────────────────
REMOTE_HOST="47.84.98.213"
REMOTE_USER="root"
SSH_KEY="$HOME/.ssh/id_ed25519"
REMOTE_DIR="/var/www/docs.freecode.codes"
SITE_URL="https://docs.freecode.codes"
# ─────────────────────────────────────────────────────

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
cd "$SCRIPT_DIR"

echo "▶ 构建文档站..."
pnpm docs:build 2>&1 | grep -vE "INVALID_ANNOTATION|comment ignored|Help:.*pure" || true

DIST="src/.vuepress/dist"
if [ ! -d "$DIST" ]; then
  echo "❌ 构建产物目录不存在: $DIST"
  exit 1
fi

PAGE_COUNT=$(find "$DIST" -name "*.html" | wc -l)
SIZE=$(du -sh "$DIST" | cut -f1)
echo "✅ 构建完成: ${PAGE_COUNT} 个页面, ${SIZE}"

echo "▶ 上传到服务器 ${REMOTE_HOST}:${REMOTE_DIR} ..."
tar czf - -C "$DIST" . | \
  ssh -i "$SSH_KEY" -o StrictHostKeyChecking=accept-new "${REMOTE_USER}@${REMOTE_HOST}" \
    "rm -rf ${REMOTE_DIR}/* && cd ${REMOTE_DIR} && tar xzf - && chown -R www-data:www-data ${REMOTE_DIR} && chmod -R 755 ${REMOTE_DIR}"
echo "✅ 上传完成"

echo "▶ 验证线上..."
sleep 1
HTTP_CODE=$(curl -s -o /dev/null -w "%{http_code}" --connect-timeout 10 "$SITE_URL/" 2>/dev/null || echo "000")
if [ "$HTTP_CODE" = "200" ]; then
  echo "✅ 发布成功! $SITE_URL (HTTP $HTTP_CODE)"
else
  echo "⚠️  线上验证返回 HTTP $HTTP_CODE（可能是 Cloudflare 缓存延迟，稍后再试）"
fi
