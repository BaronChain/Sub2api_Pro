package service

import (
	"context"
	"errors"
	"testing"
)

func TestStreamDisconnectBillable(t *testing.T) {
	errDisc := errors.New("stream usage incomplete: context canceled")

	tests := []struct {
		name   string
		result *streamingResult
		err    error
		want   bool
	}{
		{
			name:   "disconnect_with_usage_billable",
			result: &streamingResult{usage: &ClaudeUsage{InputTokens: 100, OutputTokens: 50}, clientDisconnect: true},
			err:    errDisc,
			want:   true,
		},
		{
			name:   "disconnect_only_cache_tokens_billable",
			result: &streamingResult{usage: &ClaudeUsage{CacheReadInputTokens: 200}, clientDisconnect: true},
			err:    errDisc,
			want:   true,
		},
		{
			name:   "disconnect_no_usage_not_billable",
			result: &streamingResult{usage: &ClaudeUsage{}, clientDisconnect: true},
			err:    errDisc,
			want:   false,
		},
		{
			name:   "not_disconnect_not_billable",
			result: &streamingResult{usage: &ClaudeUsage{OutputTokens: 10}, clientDisconnect: false},
			err:    errDisc,
			want:   false,
		},
		{
			name:   "no_error_not_billable",
			result: &streamingResult{usage: &ClaudeUsage{OutputTokens: 10}, clientDisconnect: true},
			err:    nil,
			want:   false,
		},
		{
			name:   "nil_result_not_billable",
			result: nil,
			err:    errDisc,
			want:   false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := streamDisconnectBillable(tc.result, tc.err); got != tc.want {
				t.Fatalf("streamDisconnectBillable = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestOpenAIStreamDisconnectBillable(t *testing.T) {
	errDisc := errors.New("stream usage incomplete after disconnect")

	if !openaiStreamDisconnectBillable(&openaiStreamingResultPassthrough{
		usage: &OpenAIUsage{OutputTokens: 30}, clientDisconnect: true,
	}, errDisc) {
		t.Error("disconnect with usage should be billable")
	}
	if openaiStreamDisconnectBillable(&openaiStreamingResultPassthrough{
		usage: &OpenAIUsage{}, clientDisconnect: true,
	}, errDisc) {
		t.Error("disconnect without usage should NOT be billable")
	}
	if openaiStreamDisconnectBillable(&openaiStreamingResultPassthrough{
		usage: &OpenAIUsage{OutputTokens: 30}, clientDisconnect: false,
	}, errDisc) {
		t.Error("non-disconnect error should NOT be billable")
	}
	if openaiStreamDisconnectBillable(&openaiStreamingResultPassthrough{
		usage: &OpenAIUsage{OutputTokens: 30}, clientDisconnect: true,
	}, nil) {
		t.Error("nil error should NOT be billable")
	}
}

// 编译期保证：辅助函数签名稳定，避免被重构破坏。
var _ = func() bool { return streamDisconnectBillable(nil, context.Canceled) }
