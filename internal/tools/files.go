package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/zamedic/labradoc-mcp/internal/labradoc"
)

// FilesListArgs represents input args for files_list tool.
type FilesListArgs struct {
	Status     string `json:"status,omitempty" jsonschema:"Filter by file status (active or archived)"`
	PageSize   int    `json:"page_size,omitempty" jsonschema:"Number of items per page (default: 50)"`
	PageNumber int    `json:"page_number,omitempty" jsonschema:"Page number to retrieve (default: 1)"`
	Query      string `json:"query,omitempty" jsonschema:"Search query to filter files"`
}

// FilesListResult represents output for files_list tool.
type FilesListResult struct {
	Items      []labradoc.File `json:"items"`
	PageSize   int             `json:"page_size"`
	PageNumber int             `json:"page_number"`
	TotalPages int             `json:"total_pages"`
	TotalItems int             `json:"total_items"`
}

// FilesSearchArgs represents input args for files_search tool.
type FilesSearchArgs struct {
	Query string `json:"query" jsonschema:"Search query string"`
}

// FilesSearchResult represents output for files_search tool.
type FilesSearchResult = FilesListResult

// FileGetArgs represents input args for files_get tool.
type FileGetArgs struct {
	FileID string `json:"file_id" jsonschema:"The ID of the file to retrieve"`
}

// FileGetResult represents output for files_get tool.
type FileGetResult = labradoc.File

// FilesDeleteArgs represents input args for files_delete tool.
type FilesDeleteArgs struct {
	IDs []string `json:"ids" jsonschema:"Array of file IDs to archive"`
}

// FilesDeleteResult represents output for files_delete tool.
type FilesDeleteResult struct {
	Message string `json:"message"`
}

// FilesList lists user's files with optional filtering and pagination.
func FilesList(ctx context.Context, client *labradoc.Client, args FilesListArgs) (*mcp.CallToolResult, FilesListResult, error) {
	params := labradoc.FilesListParams{
		Status:     args.Status,
		PageSize:   args.PageSize,
		PageNumber: args.PageNumber,
		Query:      args.Query,
	}

	if params.PageSize <= 0 {
		params.PageSize = 50
	}
	if params.PageNumber <= 0 {
		params.PageNumber = 1
	}

	result, err := client.FilesList(ctx, params)
	if err != nil {
		return errorResult(err), FilesListResult{}, nil
	}

	return nil, FilesListResult{
		Items:      result.Items,
		PageSize:   result.PageSize,
		PageNumber: result.PageNumber,
		TotalPages: result.TotalPages,
		TotalItems: result.TotalItems,
	}, nil
}

// FilesSearch searches for files by query string.
func FilesSearch(ctx context.Context, client *labradoc.Client, args FilesSearchArgs) (*mcp.CallToolResult, FilesSearchResult, error) {
	if args.Query == "" {
		return errorResult(errMissingRequired("query")), FilesSearchResult{}, nil
	}

	result, err := client.FilesSearch(ctx, args.Query)
	if err != nil {
		return errorResult(err), FilesSearchResult{}, nil
	}

	return nil, FilesSearchResult{
		Items:      result.Items,
		PageSize:   result.PageSize,
		PageNumber: result.PageNumber,
		TotalPages: result.TotalPages,
		TotalItems: result.TotalItems,
	}, nil
}

// FileGet returns a single file by ID.
func FileGet(ctx context.Context, client *labradoc.Client, args FileGetArgs) (*mcp.CallToolResult, FileGetResult, error) {
	if args.FileID == "" {
		return errorResult(errMissingRequired("file_id")), labradoc.File{}, nil
	}

	result, err := client.FileGet(ctx, args.FileID)
	if err != nil {
		return errorResult(err), labradoc.File{}, nil
	}

	return nil, *result, nil
}

// FilesDelete archives files by their IDs.
func FilesDelete(ctx context.Context, client *labradoc.Client, args FilesDeleteArgs) (*mcp.CallToolResult, FilesDeleteResult, error) {
	if len(args.IDs) == 0 {
		return errorResult(errMissingRequired("ids")), FilesDeleteResult{}, nil
	}

	err := client.FilesDelete(ctx, args.IDs)
	if err != nil {
		return errorResult(err), FilesDeleteResult{}, nil
	}

	return nil, FilesDeleteResult{Message: "Files archived successfully"}, nil
}

// ---- shared helpers ----

type missingRequiredError struct {
	param string
}

func (e missingRequiredError) Error() string {
	return "missing required parameter: " + e.param
}

func errMissingRequired(param string) error {
	return missingRequiredError{param: param}
}

func errorResult(err error) *mcp.CallToolResult {
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "Error: " + err.Error()},
		},
		IsError: true,
	}
}
