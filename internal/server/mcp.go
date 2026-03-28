package server

import (
	"context"
	"log/slog"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/zamedic/labradoc-mcp/internal/labradoc"
	"github.com/zamedic/labradoc-mcp/internal/tools"
)

// MCPServer wraps the MCP server with Labradoc tools.
type MCPServer struct {
	client *labradoc.Client
	logger *slog.Logger
	server *mcp.Server
}

// NewMCPServer creates a new MCP server with Labradoc tools.
func NewMCPServer(client *labradoc.Client, logger *slog.Logger) *MCPServer {
	s := mcp.NewServer(&mcp.Implementation{
		Name:    "labradoc-mcp",
		Version: "1.0.0",
	}, nil)

	srv := &MCPServer{
		client: client,
		logger: logger,
		server: s,
	}

	// Register all tools
	srv.registerTools()

	return srv
}

// Run starts the MCP server using stdio transport.
func (s *MCPServer) Run(ctx context.Context) error {
	s.logger.Info("Starting MCP server with stdio transport")
	return s.server.Run(ctx, &mcp.StdioTransport{})
}

// registerTools registers all MCP tools.
func (s *MCPServer) registerTools() {
	// File tools
	s.registerFileTools()

	// Email tools
	s.registerEmailTools()

	// Task tools
	s.registerTaskTools()

	// User tools
	s.registerUserTools()

	// Integration tools
	s.registerIntegrationTools()
}

// registerFileTools registers file-related MCP tools.
func (s *MCPServer) registerFileTools() {
	// files_list
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "files_list",
		Description: "List user's files with optional filtering and pagination",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args tools.FilesListArgs) (*mcp.CallToolResult, tools.FilesListResult, error) {
		return tools.FilesList(ctx, s.client, args)
	})

	// files_search
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "files_search",
		Description: "Search files by query string",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args tools.FilesSearchArgs) (*mcp.CallToolResult, tools.FilesSearchResult, error) {
		return tools.FilesSearch(ctx, s.client, args)
	})

	// files_get
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "files_get",
		Description: "Get a single file by ID",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args tools.FileGetArgs) (*mcp.CallToolResult, tools.FileGetResult, error) {
		return tools.FileGet(ctx, s.client, args)
	})

	// files_delete
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "files_delete",
		Description: "Archive files by their IDs",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args tools.FilesDeleteArgs) (*mcp.CallToolResult, tools.FilesDeleteResult, error) {
		return tools.FilesDelete(ctx, s.client, args)
	})
}

// registerEmailTools registers email-related MCP tools.
func (s *MCPServer) registerEmailTools() {
	// email_addresses_list
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "email_addresses_list",
		Description: "List user's inbound email addresses",
	}, func(ctx context.Context, req *mcp.CallToolRequest, _ any) (*mcp.CallToolResult, tools.EmailAddressesListResult, error) {
		return tools.EmailAddressesList(ctx, s.client)
	})

	// email_addresses_create
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "email_addresses_create",
		Description: "Request a new inbound email address",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args tools.EmailAddressCreateArgs) (*mcp.CallToolResult, tools.EmailAddressCreateResult, error) {
		return tools.EmailAddressCreate(ctx, s.client, args)
	})

	// emails_list
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "emails_list",
		Description: "List ingested emails",
	}, func(ctx context.Context, req *mcp.CallToolRequest, _ any) (*mcp.CallToolResult, tools.EmailsListResult, error) {
		return tools.EmailsList(ctx, s.client)
	})
}

// registerTaskTools registers task-related MCP tools.
func (s *MCPServer) registerTaskTools() {
	// tasks_list
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "tasks_list",
		Description: "List tasks extracted from documents",
	}, func(ctx context.Context, req *mcp.CallToolRequest, _ any) (*mcp.CallToolResult, tools.TasksListResult, error) {
		return tools.TasksList(ctx, s.client)
	})

	// tasks_close
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "tasks_close",
		Description: "Close/complete tasks by their IDs",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args tools.TasksCloseArgs) (*mcp.CallToolResult, tools.TasksCloseResult, error) {
		return tools.TasksClose(ctx, s.client, args)
	})
}

// registerUserTools registers user-related MCP tools.
func (s *MCPServer) registerUserTools() {
	// user_stats
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "user_stats",
		Description: "Get user statistics including completed pages and unlimited pages status",
	}, func(ctx context.Context, req *mcp.CallToolRequest, _ any) (*mcp.CallToolResult, tools.UserStatsResult, error) {
		return tools.UserStats(ctx, s.client)
	})

	// billing_checkout
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "billing_checkout",
		Description: "Create a Stripe checkout session for AI credits",
	}, func(ctx context.Context, req *mcp.CallToolRequest, _ any) (*mcp.CallToolResult, tools.BillingCheckoutResult, error) {
		return tools.BillingCheckout(ctx, s.client)
	})
}

// registerIntegrationTools registers integration-related MCP tools.
func (s *MCPServer) registerIntegrationTools() {
	// google_drive_status
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "google_drive_status",
		Description: "Check if Google Drive is connected",
	}, func(ctx context.Context, req *mcp.CallToolRequest, _ any) (*mcp.CallToolResult, tools.GoogleDriveStatusResult, error) {
		return tools.GoogleDriveStatus(ctx, s.client)
	})

	// google_drive_connect
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "google_drive_connect",
		Description: "Start Google Drive OAuth flow to connect your account",
	}, func(ctx context.Context, req *mcp.CallToolRequest, _ any) (*mcp.CallToolResult, tools.GoogleDriveConnectResult, error) {
		return tools.GoogleDriveConnect(ctx, s.client)
	})

	// google_gmail_status
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "google_gmail_status",
		Description: "Check if Gmail is connected",
	}, func(ctx context.Context, req *mcp.CallToolRequest, _ any) (*mcp.CallToolResult, tools.GoogleGmailStatusResult, error) {
		return tools.GoogleGmailStatus(ctx, s.client)
	})

	// google_gmail_connect
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "google_gmail_connect",
		Description: "Start Gmail OAuth flow to connect your account",
	}, func(ctx context.Context, req *mcp.CallToolRequest, _ any) (*mcp.CallToolResult, tools.GoogleGmailConnectResult, error) {
		return tools.GoogleGmailConnect(ctx, s.client)
	})

	// microsoft_outlook_status
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "microsoft_outlook_status",
		Description: "Check if Microsoft Outlook is connected",
	}, func(ctx context.Context, req *mcp.CallToolRequest, _ any) (*mcp.CallToolResult, tools.MicrosoftOutlookStatusResult, error) {
		return tools.MicrosoftOutlookStatus(ctx, s.client)
	})

	// microsoft_outlook_connect
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "microsoft_outlook_connect",
		Description: "Start Microsoft Outlook OAuth flow to connect your account",
	}, func(ctx context.Context, req *mcp.CallToolRequest, _ any) (*mcp.CallToolResult, tools.MicrosoftOutlookConnectResult, error) {
		return tools.MicrosoftOutlookConnect(ctx, s.client)
	})
}
