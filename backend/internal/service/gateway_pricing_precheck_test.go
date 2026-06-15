package service

import (
	"context"
	"testing"

	"github.com/Wei-Shaw/sub2api/internal/config"
)

// TestHasPricingForModel 覆盖转发前计费预检：
//   - 已知模型（命中 fallback）→ 有定价
//   - 未知模型 → 无定价（应被拦截）
//   - 空模型 → 放行（交由上游正常校验）
func TestHasPricingForModel(t *testing.T) {
	billing := NewBillingService(&config.Config{}, nil)
	svc := &GatewayService{billingService: billing}
	apiKey := &APIKey{} // Group 为 nil，跳过渠道定价

	tests := []struct {
		name  string
		model string
		want  bool
	}{
		{name: "known_claude_sonnet_fallback", model: "claude-sonnet-4-5-20250929", want: true},
		{name: "known_claude_opus_fallback", model: "claude-opus-4-5", want: true},
		{name: "unknown_model_rejected", model: "totally-unknown-model-xyz", want: false},
		{name: "unknown_openai_rejected", model: "gpt-9-imaginary", want: false},
		{name: "empty_model_passthrough", model: "", want: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := svc.HasPricingForModel(context.Background(), tc.model, apiKey)
			if got != tc.want {
				t.Fatalf("HasPricingForModel(%q) = %v, want %v", tc.model, got, tc.want)
			}
		})
	}
}

// TestCalculateTokenCostPricingUnavailable 验证未知模型计费时标记
// PricingUnavailable 并按 0 计费（不丢请求）。
func TestCalculateTokenCostPricingUnavailable(t *testing.T) {
	billing := NewBillingService(&config.Config{}, nil)
	svc := &GatewayService{billingService: billing}

	result := &ForwardResult{
		Usage: ClaudeUsage{InputTokens: 100, OutputTokens: 50},
	}
	apiKey := &APIKey{}
	opts := &recordUsageOpts{}

	cost := svc.calculateTokenCost(context.Background(), result, apiKey, "totally-unknown-model-xyz", 1.0, opts)
	if cost == nil {
		t.Fatal("expected non-nil cost breakdown")
	}
	if !cost.PricingUnavailable {
		t.Errorf("expected PricingUnavailable=true for unknown model")
	}
	if cost.ActualCost != 0 {
		t.Errorf("expected ActualCost=0 for unknown model, got %v", cost.ActualCost)
	}
}
