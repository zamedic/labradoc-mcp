# Labradoc MCP Server

A Model Context Protocol (MCP) server that exposes Labradoc API capabilities as MCP tools. This enables AI assistants like Claude Desktop, OpenAI Agents, and other MCP-compatible clients to interact with Labradoc's document management, email ingestion, task extraction, and integration features.

## Features

- **Files**: List, search, get, and archive documents
- **Email**: Manage inbound email addresses and view ingested emails
- **Tasks**: List and close tasks extracted from documents
- **User**: View usage statistics and purchase AI credits
- **Integrations**: Connect Google Drive, Gmail, and Microsoft Outlook

## Prerequisites

- Go 1.21 or later
- A Labradoc API key

## Getting an API Key

1. Log in to your Labradoc account at [https://labradoc.eu](https://labradoc.eu)
2. Navigate to Settings → API
3. Generate a new API key
4. Copy the key and keep it secure

## Configuration

The server uses environment variables for configuration:

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `LABRADOC_API_KEY` | Yes | - | Your Labradoc API key |
| `LABRADOC_API_URL` | No | `https://labradoc.eu` | Labradoc API base URL |
| `LABRADOC_LOG_LEVEL` | No | `info` | Log level: debug, info, warn, error |
| `LABRADOC_LOG_FORMAT` | No | `text` | Log format: text, json |

You can also use a `.env` file (copy from `.env.example`):

```bash
cp .env.example .env
# Edit .env and add your API key
```

## Building

```bash
# Clone the repository
git clone https://github.com/zamedic/labradoc-mcp.git
cd labradoc-mcp

# Install dependencies
go mod tidy

# Build
go build -o labradoc-mcp ./cmd/server

# Run
LABRADOC_API_KEY=your-key ./labradoc-mcp
```

## Available MCP Tools

### Files

| Tool | Description |
|------|-------------|
| `files_list` | List user's files with optional status filter, pagination, and search |
| `files_search` | Search files by query string |
| `files_get` | Get a single file by ID |
| `files_delete` | Archive files by their IDs |

### Email

| Tool | Description |
|------|-------------|
| `email_addresses_list` | List user's inbound email addresses |
| `email_addresses_create` | Request a new inbound email address |
| `emails_list` | List ingested emails |

### Tasks

| Tool | Description |
|------|-------------|
| `tasks_list` | List tasks extracted from documents |
| `tasks_close` | Close/complete tasks by their IDs |

### User & Billing

| Tool | Description |
|------|-------------|
| `user_stats` | Get user statistics (completed pages, unlimited pages status) |
| `billing_checkout` | Create a Stripe checkout session for AI credits |

### Integrations

| Tool | Description |
|------|-------------|
| `google_drive_status` | Check if Google Drive is connected |
| `google_drive_connect` | Start Google Drive OAuth flow |
| `google_gmail_status` | Check if Gmail is connected |
| `google_gmail_connect` | Start Gmail OAuth flow |
| `microsoft_outlook_status` | Check if Microsoft Outlook is connected |
| `microsoft_outlook_connect` | Start Microsoft Outlook OAuth flow |

## Usage

### Claude Desktop (macOS/Windows)

1. Install Claude Desktop from [claude.ai/desktop](https://claude.ai/desktop)
2. Edit your Claude Desktop configuration file:
   - **macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
   - **Windows**: `%APPDATA%\Claude\claude_desktop_config.json`

3. Add the MCP server configuration:

```json
{
  "mcpServers": {
    "labradoc": {
      "command": "path/to/labradoc-mcp",
      "env": {
        "LABRADOC_API_KEY": "your-api-key"
      }
    }
  }
}
```

4. Restart Claude Desktop

### Claude CLI

```bash
claude mcp add labradoc -- path/to/labradoc-mcp
```

Set the API key:
```bash
export LABRADOC_API_KEY=your-api-key
```

### OpenAI Agents SDK

```python
from agents import Agent, AgentRegistry

# Register the Labradoc MCP server
AgentRegistry.register(
    name="labradoc",
    command="path/to/labradoc-mcp",
    env={"LABRADOC_API_KEY": "your-api-key"}
)

# Use in an agent
agent = Agent(
    name="document-assistant",
    mcp_servers=["labradoc"]
)
```

### Direct stdio Usage

You can also run the server directly and communicate via stdio:

```bash
LABRADOC_API_KEY=your-key ./labradoc-mcp
```

The server reads JSON-RPC requests from stdin and writes responses to stdout.

## Architecture

```
├── cmd/server/main.go       # Entry point, config, and startup
├── internal/
│   ├── server/mcp.go         # MCP server implementation
│   ├── labradoc/client.go    # Labradoc API client
│   └── tools/                # MCP tool implementations
│       ├── files.go
│       ├── email.go
│       ├── tasks.go
│       ├── user.go
│       └── integrations.go
├── .env.example
└── README.md
```

## Error Handling

All tools return structured error messages when failures occur. The `IsError` field in the response indicates whether the result is an error.

## Logging

The server uses structured logging via `log/slog`. Configure log level and format via environment variables:

```bash
LABRADOC_LOG_LEVEL=debug LABRADOC_LOG_FORMAT=json ./labradoc-mcp
```

## License

MIT
