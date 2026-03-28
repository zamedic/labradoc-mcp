package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/zamedic/labradoc-mcp/internal/labradoc"
)

// GoogleDriveStatusResult represents output for google_drive_status tool.
type GoogleDriveStatusResult = labradoc.IntegrationStatus

// GoogleDriveConnectResult represents output for google_drive_connect tool.
type GoogleDriveConnectResult = labradoc.BillingCheckoutResponse

// GoogleGmailStatusResult represents output for google_gmail_status tool.
type GoogleGmailStatusResult = labradoc.IntegrationStatus

// GoogleGmailConnectResult represents output for google_gmail_connect tool.
type GoogleGmailConnectResult = labradoc.BillingCheckoutResponse

// MicrosoftOutlookStatusResult represents output for microsoft_outlook_status tool.
type MicrosoftOutlookStatusResult = labradoc.IntegrationStatus

// MicrosoftOutlookConnectResult represents output for microsoft_outlook_connect tool.
type MicrosoftOutlookConnectResult = labradoc.BillingCheckoutResponse

// GoogleDriveStatus checks if Google Drive is connected.
func GoogleDriveStatus(ctx context.Context, client *labradoc.Client) (*mcp.CallToolResult, GoogleDriveStatusResult, error) {
	result, err := client.GoogleDriveStatus(ctx)
	if err != nil {
		return errorResult(err), labradoc.IntegrationStatus{}, nil
	}

	return nil, *result, nil
}

// GoogleDriveConnect starts Google Drive OAuth flow.
func GoogleDriveConnect(ctx context.Context, client *labradoc.Client) (*mcp.CallToolResult, GoogleDriveConnectResult, error) {
	result, err := client.GoogleDriveConnect(ctx)
	if err != nil {
		return errorResult(err), labradoc.BillingCheckoutResponse{}, nil
	}

	return nil, *result, nil
}

// GoogleGmailStatus checks if Gmail is connected.
func GoogleGmailStatus(ctx context.Context, client *labradoc.Client) (*mcp.CallToolResult, GoogleGmailStatusResult, error) {
	result, err := client.GoogleGmailStatus(ctx)
	if err != nil {
		return errorResult(err), labradoc.IntegrationStatus{}, nil
	}

	return nil, *result, nil
}

// GoogleGmailConnect starts Gmail OAuth flow.
func GoogleGmailConnect(ctx context.Context, client *labradoc.Client) (*mcp.CallToolResult, GoogleGmailConnectResult, error) {
	result, err := client.GoogleGmailConnect(ctx)
	if err != nil {
		return errorResult(err), labradoc.BillingCheckoutResponse{}, nil
	}

	return nil, *result, nil
}

// MicrosoftOutlookStatus checks if Microsoft Outlook is connected.
func MicrosoftOutlookStatus(ctx context.Context, client *labradoc.Client) (*mcp.CallToolResult, MicrosoftOutlookStatusResult, error) {
	result, err := client.MicrosoftOutlookStatus(ctx)
	if err != nil {
		return errorResult(err), labradoc.IntegrationStatus{}, nil
	}

	return nil, *result, nil
}

// MicrosoftOutlookConnect starts Microsoft Outlook OAuth flow.
func MicrosoftOutlookConnect(ctx context.Context, client *labradoc.Client) (*mcp.CallToolResult, MicrosoftOutlookConnectResult, error) {
	result, err := client.MicrosoftOutlookConnect(ctx)
	if err != nil {
		return errorResult(err), labradoc.BillingCheckoutResponse{}, nil
	}

	return nil, *result, nil
}
