//go:build integration

package repository

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/Wei-Shaw/sub2api/internal/pkg/tlsfingerprint"
)

// TestH2OverUTLS_RealEndpoint 验证 h2-over-utls 端到端真实可用：
// 对一个已知支持 HTTP/2 的公共 HTTPS 端点发起请求，断言协议协商为 HTTP/2.0。
//
// 运行：go test -tags=integration -run TestH2OverUTLS_RealEndpoint ./internal/repository/...
func TestH2OverUTLS_RealEndpoint(t *testing.T) {
	settings := poolSettings{
		maxIdleConns:          10,
		maxIdleConnsPerHost:   10,
		idleConnTimeout:       30 * time.Second,
		responseHeaderTimeout: 15 * time.Second,
	}
	profile := &tlsfingerprint.Profile{Name: "Built-in Default (Node.js 24.x)"}

	rt, err := buildUpstreamTransportWithTLSFingerprint(settings, (*url.URL)(nil), profile)
	if err != nil {
		t.Fatalf("build transport: %v", err)
	}
	client := &http.Client{Transport: rt, Timeout: 20 * time.Second}

	// www.cloudflare.com 支持 HTTP/2 与 Node.js 风格 ALPN。
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.cloudflare.com/", nil)
	resp, err := client.Do(req)
	if err != nil {
		t.Skipf("external service unavailable: %v", err)
	}
	defer resp.Body.Close()
	_, _ = io.Copy(io.Discard, resp.Body)

	if resp.ProtoMajor != 2 {
		t.Errorf("expected HTTP/2 over utls, got %s", resp.Proto)
	} else {
		t.Logf("negotiated protocol: %s", resp.Proto)
	}
}
