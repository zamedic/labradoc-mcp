import type { Client } from "../client.js";

export type GoogleDriveStatusResult = Client extends { googleDriveStatus(): Promise<infer R> } ? R : never;
export type GoogleDriveConnectResult = Client extends { googleDriveConnect(): Promise<infer R> } ? R : never;
export type GoogleGmailStatusResult = Client extends { googleGmailStatus(): Promise<infer R> } ? R : never;
export type GoogleGmailConnectResult = Client extends { googleGmailConnect(): Promise<infer R> } ? R : never;
export type MicrosoftOutlookStatusResult = Client extends { microsoftOutlookStatus(): Promise<infer R> } ? R : never;
export type MicrosoftOutlookConnectResult = Client extends { microsoftOutlookConnect(): Promise<infer R> } ? R : never;

export function newErrorResult(msg: string) {
  return {
    content: [{ type: "text" as const, text: `Error: ${msg}` }],
    isError: true,
  };
}

export async function googleDriveStatus(client: Client) {
  try {
    const result = await client.googleDriveStatus();
    return { content: [{ type: "json" as const, text: JSON.stringify(result) }] };
  } catch (e) {
    return newErrorResult(e instanceof Error ? e.message : String(e));
  }
}

export async function googleDriveConnect(client: Client) {
  try {
    const result = await client.googleDriveConnect();
    return { content: [{ type: "json" as const, text: JSON.stringify(result) }] };
  } catch (e) {
    return newErrorResult(e instanceof Error ? e.message : String(e));
  }
}

export async function googleGmailStatus(client: Client) {
  try {
    const result = await client.googleGmailStatus();
    return { content: [{ type: "json" as const, text: JSON.stringify(result) }] };
  } catch (e) {
    return newErrorResult(e instanceof Error ? e.message : String(e));
  }
}

export async function googleGmailConnect(client: Client) {
  try {
    const result = await client.googleGmailConnect();
    return { content: [{ type: "json" as const, text: JSON.stringify(result) }] };
  } catch (e) {
    return newErrorResult(e instanceof Error ? e.message : String(e));
  }
}

export async function microsoftOutlookStatus(client: Client) {
  try {
    const result = await client.microsoftOutlookStatus();
    return { content: [{ type: "json" as const, text: JSON.stringify(result) }] };
  } catch (e) {
    return newErrorResult(e instanceof Error ? e.message : String(e));
  }
}

export async function microsoftOutlookConnect(client: Client) {
  try {
    const result = await client.microsoftOutlookConnect();
    return { content: [{ type: "json" as const, text: JSON.stringify(result) }] };
  } catch (e) {
    return newErrorResult(e instanceof Error ? e.message : String(e));
  }
}
