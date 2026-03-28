import type { Client } from "../client.js";

export interface EmailAddressCreateArgs {
  description?: string;
}

export type EmailAddressCreateResult = Client extends { emailAddressCreate(_: any): Promise<infer R> } ? R : never;
export type EmailAddressesListResult = Client extends { emailAddressesList(): Promise<infer R> } ? R : never;
export type EmailsListResult = Client extends { emailsList(): Promise<infer R> } ? R : never;

export function newErrorResult(msg: string) {
  return {
    content: [{ type: "text" as const, text: `Error: ${msg}` }],
    isError: true,
  };
}

export async function emailAddressesList(client: Client) {
  try {
    const result = await client.emailAddressesList();
    return {
      content: [{ type: "json" as const, text: JSON.stringify(result) }],
    };
  } catch (e) {
    return newErrorResult(e instanceof Error ? e.message : String(e));
  }
}

export async function emailAddressCreate(client: Client, args: EmailAddressCreateArgs) {
  try {
    const result = await client.emailAddressCreate({ description: args.description });
    return {
      content: [{ type: "json" as const, text: JSON.stringify(result) }],
    };
  } catch (e) {
    return newErrorResult(e instanceof Error ? e.message : String(e));
  }
}

export async function emailsList(client: Client) {
  try {
    const result = await client.emailsList();
    return {
      content: [{ type: "json" as const, text: JSON.stringify(result) }],
    };
  } catch (e) {
    return newErrorResult(e instanceof Error ? e.message : String(e));
  }
}
