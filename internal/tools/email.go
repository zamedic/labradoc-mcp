package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/zamedic/labradoc-mcp/internal/labradoc"
)

// EmailAddressCreateArgs represents input args for email_addresses_create tool.
type EmailAddressCreateArgs struct {
	Description string `json:"description,omitempty" jsonschema:"Optional description for the email address"`
}

// EmailAddressCreateResult represents output for email_addresses_create tool.
type EmailAddressCreateResult = labradoc.EmailAddress

// EmailAddressesListResult represents output for email_addresses_list tool.
type EmailAddressesListResult = labradoc.EmailAddressesResponse

// EmailsListResult represents output for emails_list tool.
type EmailsListResult = labradoc.EmailsResponse

// EmailAddressesList lists user's inbound email addresses.
func EmailAddressesList(ctx context.Context, client *labradoc.Client) (*mcp.CallToolResult, EmailAddressesListResult, error) {
	result, err := client.EmailAddressesList(ctx)
	if err != nil {
		return errorResult(err), labradoc.EmailAddressesResponse{}, nil
	}

	return nil, *result, nil
}

// EmailAddressCreate requests a new inbound email address.
func EmailAddressCreate(ctx context.Context, client *labradoc.Client, args EmailAddressCreateArgs) (*mcp.CallToolResult, EmailAddressCreateResult, error) {
	params := labradoc.EmailAddressCreateParams{
		Description: args.Description,
	}

	result, err := client.EmailAddressCreate(ctx, params)
	if err != nil {
		return errorResult(err), labradoc.EmailAddress{}, nil
	}

	return nil, *result, nil
}

// EmailsList lists ingested emails.
func EmailsList(ctx context.Context, client *labradoc.Client) (*mcp.CallToolResult, EmailsListResult, error) {
	result, err := client.EmailsList(ctx)
	if err != nil {
		return errorResult(err), labradoc.EmailsResponse{}, nil
	}

	return nil, *result, nil
}
