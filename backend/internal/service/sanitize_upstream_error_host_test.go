package service

import (
	"strings"
	"testing"
)

func TestSanitizeUpstreamErrorMessage_HostAndURL(t *testing.T) {
	tests := []struct {
		name       string
		in         string
		wantSubstr string // 期望结果包含
		notSubstr  []string
	}{
		{
			name:       "anthropic_url",
			in:         "Post \"https://api.anthropic.com/v1/messages?beta=true\": dial tcp timeout",
			wantSubstr: "upstream service",
			notSubstr:  []string{"api.anthropic.com", "/v1/messages"},
		},
		{
			name:       "googleapis_url",
			in:         "error from https://generativelanguage.googleapis.com/v1beta/models: 500",
			wantSubstr: "upstream service",
			notSubstr:  []string{"generativelanguage.googleapis.com"},
		},
		{
			name:       "bare_host",
			in:         "connection refused by api.openai.com",
			wantSubstr: "upstream service",
			notSubstr:  []string{"api.openai.com"},
		},
		{
			name:       "vertex_host",
			in:         "cloudcode-pa.googleapis.com returned 403",
			wantSubstr: "upstream service",
			notSubstr:  []string{"cloudcode-pa.googleapis.com"},
		},
		{
			name:       "sensitive_query_param_in_url",
			in:         "GET https://example-upstream.anthropic.com/v1?key=sk-secret-123 failed",
			wantSubstr: "upstream service",
			notSubstr:  []string{"sk-secret-123", "anthropic.com"},
		},
		{
			name:       "no_url_unchanged",
			in:         "internal timeout while reading stream",
			wantSubstr: "internal timeout while reading stream",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := sanitizeUpstreamErrorMessageForClient(tc.in)
			if tc.wantSubstr != "" && !strings.Contains(got, tc.wantSubstr) {
				t.Errorf("result %q should contain %q", got, tc.wantSubstr)
			}
			for _, ns := range tc.notSubstr {
				if strings.Contains(got, ns) {
					t.Errorf("result %q should NOT contain %q", got, ns)
				}
			}
		})
	}
}
