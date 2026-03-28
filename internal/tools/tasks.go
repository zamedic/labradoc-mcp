package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/zamedic/labradoc-mcp/internal/labradoc"
)

// TasksCloseArgs represents input args for tasks_close tool.
type TasksCloseArgs struct {
	IDs []string `json:"ids" jsonschema:"Array of task IDs to close"`
}

// TasksCloseResult represents output for tasks_close tool.
type TasksCloseResult struct {
	Message string `json:"message"`
}

// TasksListResult represents output for tasks_list tool.
type TasksListResult = labradoc.TasksResponse

// TasksList lists tasks extracted from documents.
func TasksList(ctx context.Context, client *labradoc.Client) (*mcp.CallToolResult, TasksListResult, error) {
	result, err := client.TasksList(ctx)
	if err != nil {
		return errorResult(err), labradoc.TasksResponse{}, nil
	}

	return nil, *result, nil
}

// TasksClose closes/completes tasks by their IDs.
func TasksClose(ctx context.Context, client *labradoc.Client, args TasksCloseArgs) (*mcp.CallToolResult, TasksCloseResult, error) {
	if len(args.IDs) == 0 {
		return errorResult(errMissingRequired("ids")), TasksCloseResult{}, nil
	}

	err := client.TasksClose(ctx, args.IDs)
	if err != nil {
		return errorResult(err), TasksCloseResult{}, nil
	}

	return nil, TasksCloseResult{Message: "Tasks closed successfully"}, nil
}
