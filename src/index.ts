#!/usr/bin/env node

import { LabradocClient } from "./client.js";
import { runServer } from "./server.js";

const apiKey = process.env.LABRADOC_API_KEY;
const apiUrl = process.env.LABRADOC_API_URL || "https://labradoc.eu";
const logLevel = process.env.LABRADOC_LOG_LEVEL || "info";

if (!apiKey) {
  console.error("LABRADOC_API_KEY is required");
  process.exit(1);
}

const log: typeof console.error = (level: string, msg: string, meta?: Record<string, unknown>) => {
  const minLevel = logLevel === "debug" ? 0 : logLevel === "info" ? 1 : logLevel === "warn" ? 2 : 3;
  const msgLevel = level === "debug" ? 0 : level === "info" ? 1 : level === "warn" ? 2 : 3;
  if (msgLevel >= minLevel) {
    const prefix = `[${level.toUpperCase()}]`;
    const metaStr = meta ? ` ${JSON.stringify(meta)}` : "";
    if (level === "error") console.error(`${prefix} ${msg}${metaStr}`);
    else console.log(`${prefix} ${msg}${metaStr}`);
  }
};

log("info", "Starting Labradoc MCP server", { api_url: apiUrl, log_level: logLevel });

const client = new LabradocClient(apiKey, apiUrl, log);

runServer(client).catch((err) => {
  log("error", "Server exited with error", { error: err.message });
  process.exit(1);
});
