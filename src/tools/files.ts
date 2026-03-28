import type { Client } from "../client.js";

export interface FilesListArgs {
  status?: string;
  page_size?: number;
  page_number?: number;
  query?: string;
}

export interface FilesListResult {
  items: Client extends { filesList(_: any): Promise<infer R> } ? R["items"] : never;
  page_size: number;
  page_number: number;
  total_pages: number;
  total_items: number;
}

export interface FilesSearchArgs {
  query: string;
}

export type FilesSearchResult = FilesListResult;

export interface FileGetArgs {
  file_id: string;
}

export type FileGetResult = Client extends { fileGet(_: any): Promise<infer R> } ? R : never;

export interface FilesDeleteArgs {
  ids: string[];
}

export interface FilesDeleteResult {
  message: string;
}

export function newErrorResult(msg: string) {
  return {
    content: [{ type: "text" as const, text: `Error: ${msg}` }],
    isError: true,
  };
}

export async function filesList(client: Client, args: FilesListArgs) {
  const result = await client.filesList({
    status: args.status,
    page_size: args.page_size,
    page_number: args.page_number,
    query: args.query,
  });
  return {
    content: [{ type: "json" as const, text: JSON.stringify(result) }],
  };
}

export async function filesSearch(client: Client, args: FilesSearchArgs) {
  if (!args.query) return newErrorResult("missing required parameter: query");
  const result = await client.filesSearch(args.query);
  return {
    content: [{ type: "json" as const, text: JSON.stringify(result) }],
  };
}

export async function fileGet(client: Client, args: FileGetArgs) {
  if (!args.file_id) return newErrorResult("missing required parameter: file_id");
  try {
    const result = await client.fileGet(args.file_id);
    return {
      content: [{ type: "json" as const, text: JSON.stringify(result) }],
    };
  } catch (e) {
    return newErrorResult(e instanceof Error ? e.message : String(e));
  }
}

export async function filesDelete(client: Client, args: FilesDeleteArgs) {
  if (!args.ids?.length) return newErrorResult("missing required parameter: ids");
  try {
    await client.filesArchive(args.ids);
    return {
      content: [{ type: "text" as const, text: JSON.stringify({ message: "Files archived successfully" }) }],
    };
  } catch (e) {
    return newErrorResult(e instanceof Error ? e.message : String(e));
  }
}
