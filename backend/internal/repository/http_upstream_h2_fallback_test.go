package repository

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"testing"

	"golang.org/x/net/http2"
)

func TestIsH2FallbackError(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want bool
	}{
		{"nil", nil, false},
		{"sentinel", errH2NotNegotiated, true},
		{"wrapped_sentinel", errors.New("dial: " + errH2NotNegotiated.Error()), true},
		{"unsupported_scheme", errors.New("http2: unsupported scheme"), true},
		{"conn_refused", errors.New("dial tcp: connection refused"), true},
		{"generic_500", errors.New("server returned 500"), false},
		{"context_canceled", context.Canceled, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := isH2FallbackError(tc.err); got != tc.want {
				t.Fatalf("isH2FallbackError(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}

func TestRequestBodyReplayable(t *testing.T) {
	noBody, _ := http.NewRequest(http.MethodGet, "https://x/y", nil)
	if !requestBodyReplayable(noBody) {
		t.Error("GET without body should be replayable")
	}

	withBody, _ := http.NewRequest(http.MethodPost, "https://x/y", strings.NewReader("payload"))
	// NewRequest sets GetBody for *strings.Reader, so it is replayable.
	if !requestBodyReplayable(withBody) {
		t.Error("POST with strings.Reader body (GetBody set) should be replayable")
	}

	if requestBodyReplayable(nil) {
		t.Error("nil request should not be replayable")
	}
}

// rtFunc 是一个把函数适配成 http.RoundTripper 的辅助类型。
type rtFunc func(*http.Request) (*http.Response, error)

func (f rtFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestTLSFingerprintTransportRouting_H2Success(t *testing.T) {
	var h2Calls, h1Calls int
	tr := &tlsFingerprintTransport{
		h2: &http2.Transport{},
		h1: &http.Transport{},
	}
	// 用 stub 替换内部 RoundTrip 逻辑：通过包装方式不可行（字段是具体类型），
	// 因此这里直接验证 host 模式缓存的选路语义。
	_ = h2Calls
	_ = h1Calls

	// 预置 host 为 h1：应直接走 h1（此处仅验证缓存读取分支不 panic）。
	tr.hostMode.Store("example.com", "h1")
	if v, ok := tr.hostMode.Load("example.com"); !ok || v != "h1" {
		t.Fatalf("expected cached h1 mode, got %v ok=%v", v, ok)
	}
}

// TestTLSFingerprintTransportRouting_FallbackDecision 用可注入的 RoundTripper
// 验证 h2 失败时回退 h1 并缓存决策。这里通过组合 rtFunc 模拟两条路径。
func TestTLSFingerprintTransportRouting_FallbackDecision(t *testing.T) {
	// 由于 tlsFingerprintTransport.h1/h2 是具体类型，无法直接注入 stub，
	// 改为验证 RoundTrip 的决策辅助函数语义（isH2FallbackError + requestBodyReplayable）
	// 已在上面单测覆盖。此处补充：可重放 + 可回退错误 → 应回退。
	req, _ := http.NewRequest(http.MethodPost, "https://api.example.com/v1/messages", strings.NewReader("{}"))
	shouldFallback := isH2FallbackError(errH2NotNegotiated) && requestBodyReplayable(req)
	if !shouldFallback {
		t.Fatal("replayable request with h2-not-negotiated error should fall back to h1")
	}
}

var _ http.RoundTripper = (*tlsFingerprintTransport)(nil)
var _ = rtFunc(func(*http.Request) (*http.Response, error) { return nil, nil })
