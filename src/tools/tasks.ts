import type { Client } from "../client.js";

export interface TasksCloseArgs {
  ids: string[];
}

export interface TasksCloseResult {
  message: string;
}

export type TasksListResult = Client extends { tasksList(): Promise<infer R> } ? R : never;

export function newErrorResult(msg: string) {
  return {
    content: [{ type: "text" as const, text: `Error: ${msg}` }],
    isError: true,
  };
}

export async function tasksList(client: Client) {
  try {
    const result = await client.tasksList();
    return {
      content: [{ type: "json" as const, text: JSON.stringify(result) }],
    };
  } catch (e) {
    return newErrorResult(e instanceof Error ? e.message : String(e));
  }
}

export async function tasksClose(client: Client, args: TasksCloseArgs) {
  if (!args.ids?.length) return newErrorResult("missing required parameter: ids");
  try {
    await client.tasksClose(args.ids);
    return {
      content: [{ type: "text" as const, text: JSON.stringify({ message: "Tasks closed successfully" }) }],
    };
  } catch (e) {
    return newErrorResult(e instanceof Error ? e.message : String(e));
  }
}
