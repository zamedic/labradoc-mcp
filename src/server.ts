import { Server } from "@modelcontextprotocol/sdk/server/index.js";
import {
  StdioServerTransport,
} from "@modelcontextprotocol/sdk/server/stdio.js";
import {
  CallToolRequestSchema,
  ListToolsRequestSchema,
  type CallToolRequest,
} from "@modelcontextprotocol/sdk/types.js";
import type { Client } from "./client.js";
import * as files from "./tools/files.js";
import * as email from "./tools/email.js";
import * as tasks from "./tools/tasks.js";
import * as user from "./tools/user.js";
import * as integrations from "./tools/integrations.js";

const TOOLS = [
  {
    name: "files_list",
    description: "List user's files with optional status filter, pagination, and search",
    inputSchema: {
      type: "object",
      properties: {
        status: { type: "string", description: "Filter by file status (active or archived)" },
        page_size: { type: "number", description: "Number of items per page (default: 50)" },
        page_number: { type: "number", description: "Page number to retrieve (default: 1)" },
        query: { type: "string", description: "Search query to filter files" },
      },
    },
  },
  {
    name: "files_search",
    description: "Search files by query string",
    inputSchema: {
      type: "object",
      properties: {
        query: { type: "string", description: "Search query string" },
      },
      required: ["query"],
    },
  },
  {
    name: "files_get",
    description: "Get a single file by ID",
    inputSchema: {
      type: "object",
      properties: {
        file_id: { type: "string", description: "The ID of the file to retrieve" },
      },
      required: ["file_id"],
    },
  },
  {
    name: "files_delete",
    description: "Archive files by their IDs",
    inputSchema: {
      type: "object",
      properties: {
        ids: { type: "array", items: { type: "string" }, description: "Array of file IDs to archive" },
      },
      required: ["ids"],
    },
  },
  {
    name: "email_addresses_list",
    description: "List user's inbound email addresses",
    inputSchema: { type: "object", properties: {} },
  },
  {
    name: "email_addresses_create",
    description: "Request a new inbound email address",
    inputSchema: {
      type: "object",
      properties: {
        description: { type: "string", description: "Optional description for the email address" },
      },
    },
  },
  {
    name: "emails_list",
    description: "List ingested emails",
    inputSchema: { type: "object", properties: {} },
  },
  {
    name: "tasks_list",
    description: "List tasks extracted from documents",
    inputSchema: { type: "object", properties: {} },
  },
  {
    name: "tasks_close",
    description: "Close/complete tasks by their IDs",
    inputSchema: {
      type: "object",
      properties: {
        ids: { type: "array", items: { type: "string" }, description: "Array of task IDs to close" },
      },
      required: ["ids"],
    },
  },
  {
    name: "user_stats",
    description: "Get user statistics including completed pages and unlimited pages status",
    inputSchema: { type: "object", properties: {} },
  },
  {
    name: "billing_checkout",
    description: "Create a Stripe checkout session for AI credits",
    inputSchema: { type: "object", properties: {} },
  },
  {
    name: "google_drive_status",
    description: "Check if Google Drive is connected",
    inputSchema: { type: "object", properties: {} },
  },
  {
    name: "google_drive_connect",
    description: "Start Google Drive OAuth flow to connect your account",
    inputSchema: { type: "object", properties: {} },
  },
  {
    name: "google_gmail_status",
    description: "Check if Gmail is connected",
    inputSchema: { type: "object", properties: {} },
  },
  {
    name: "google_gmail_connect",
    description: "Start Gmail OAuth flow to connect your account",
    inputSchema: { type: "object", properties: {} },
  },
  {
    name: "microsoft_outlook_status",
    description: "Check if Microsoft Outlook is connected",
    inputSchema: { type: "object", properties: {} },
  },
  {
    name: "microsoft_outlook_connect",
    description: "Start Microsoft Outlook OAuth flow to connect your account",
    inputSchema: { type: "object", properties: {} },
  },
] as const;

export function createServer(client: Client) {
  const server = new Server(
    { name: "labradoc-mcp", version: "1.0.0" },
    { capabilities: { tools: {} } },
  );

  server.setRequestHandler(ListToolsRequestSchema, async () => ({
    tools: TOOLS.map((t) => ({ name: t.name, description: t.description, inputSchema: t.inputSchema })),
  }));

  server.setRequestHandler(CallToolRequestSchema, async (request: CallToolRequest) => {
    const { name, arguments: args = {} } = request.params;

    switch (name) {
      case "files_list":
        return files.filesList(client, args as files.FilesListArgs);
      case "files_search":
        return files.filesSearch(client, args as files.FilesSearchArgs);
      case "files_get":
        return files.fileGet(client, args as files.FileGetArgs);
      case "files_delete":
        return files.filesDelete(client, args as files.FilesDeleteArgs);
      case "email_addresses_list":
        return email.emailAddressesList(client);
      case "email_addresses_create":
        return email.emailAddressCreate(client, args as email.EmailAddressCreateArgs);
      case "emails_list":
        return email.emailsList(client);
      case "tasks_list":
        return tasks.tasksList(client);
      case "tasks_close":
        return tasks.tasksClose(client, args as tasks.TasksCloseArgs);
      case "user_stats":
        return user.userStats(client);
      case "billing_checkout":
        return user.billingCheckout(client);
      case "google_drive_status":
        return integrations.googleDriveStatus(client);
      case "google_drive_connect":
        return integrations.googleDriveConnect(client);
      case "google_gmail_status":
        return integrations.googleGmailStatus(client);
      case "google_gmail_connect":
        return integrations.googleGmailConnect(client);
      case "microsoft_outlook_status":
        return integrations.microsoftOutlookStatus(client);
      case "microsoft_outlook_connect":
        return integrations.microsoftOutlookConnect(client);
      default:
        return {
          content: [{ type: "text" as const, text: `Unknown tool: ${name}` }],
          isError: true,
        };
    }
  });

  return server;
}

export async function runServer(client: Client): Promise<void> {
  const server = createServer(client);
  const transport = new StdioServerTransport();
  await server.connect(transport);
}
