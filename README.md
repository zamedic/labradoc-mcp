# Labradoc MCP Server

A Model Context Protocol (MCP) server that exposes Labradoc API capabilities as MCP tools. This enables AI assistants like Claude Desktop, OpenAI Agents, and other MCP-compatible clients to interact with Labradoc's document management, email ingestion, task extraction, and integration features.

## Features

- **Files**: List, search, get, and archive documents
- **Email**: Manage inbound email addresses and view ingested emails
- **Tasks**: List and close tasks extracted from documents
- **User**: View usage statistics and purchase AI credits
- **Integrations**: Connect Google Drive, Gmail, and Microsoft Outlook

## Requirements

- Node.js 20+
- A Labradoc API key

## Quick Start

```bash
npx @labradoc/mcp
```

Or install globally:

```bash
npm install -g @labradoc/mcp
LABRADOC_API_KEY=your-key labradoc-mcp
```

## Configuration

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `LABRADOC_API_KEY` | Yes | - | Your Labradoc API key |
| `LABRADOC_API_URL` | No | `https://labradoc.eu` | Labradoc API base URL |
| `LABRADOC_LOG_LEVEL` | No | `info` | Log level: debug, info, warn, error |

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

Edit your Claude Desktop configuration file:
- **macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
- **Windows**: `%APPDATA%\Claude\claude_desktop_config.json`

```json
{
  "mcpServers": {
    "labradoc": {
      "command": "npx",
      "args": ["@labradoc/mcp"]
    }
  }
}
```

Or with a global install:

```json
{
  "mcpServers": {
    "labradoc": {
      "command": "labradoc-mcp"
    }
  }
}
```

Set the API key in your environment or `.env` file:
```
LABRADOC_API_KEY=your-api-key
```

### Claude CLI

```bash
claude mcp add labradoc -- npx @labradoc/mcp
```

### Direct stdio Usage

```bash
LABRADOC_API_KEY=your-key npx @labradoc/mcp
```

## Development

```bash
npm install
npm run build
LABRADOC_API_KEY=your-key node dist/index.js
```

## License

MIT
