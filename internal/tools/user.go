package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/zamedic/labradoc-mcp/internal/labradoc"
)

// UserStatsResult represents output for user_stats tool.
type UserStatsResult = labradoc.UserStats

// BillingCheckoutResult represents output for billing_checkout tool.
type BillingCheckoutResult = labradoc.BillingCheckoutResponse

// UserStats returns user statistics.
func UserStats(ctx context.Context, client *labradoc.Client) (*mcp.CallToolResult, UserStatsResult, error) {
	result, err := client.UserStats(ctx)
	if err != nil {
		return errorResult(err), labradoc.UserStats{}, nil
	}

	return nil, *result, nil
}

// BillingCheckout creates a Stripe checkout session for AI credits.
func BillingCheckout(ctx context.Context, client *labradoc.Client) (*mcp.CallToolResult, BillingCheckoutResult, error) {
	result, err := client.BillingCheckout(ctx)
	if err != nil {
		return errorResult(err), labradoc.BillingCheckoutResponse{}, nil
	}

	return nil, *result, nil
}
