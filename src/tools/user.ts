import type { Client } from "../client.js";

export type UserStatsResult = Client extends { userStats(): Promise<infer R> } ? R : never;
export type BillingCheckoutResult = Client extends { billingCheckout(): Promise<infer R> } ? R : never;

export function newErrorResult(msg: string) {
  return {
    content: [{ type: "text" as const, text: `Error: ${msg}` }],
    isError: true,
  };
}

export async function userStats(client: Client) {
  try {
    const result = await client.userStats();
    return {
      content: [{ type: "json" as const, text: JSON.stringify(result) }],
    };
  } catch (e) {
    return newErrorResult(e instanceof Error ? e.message : String(e));
  }
}

export async function billingCheckout(client: Client) {
  try {
    const result = await client.billingCheckout();
    return {
      content: [{ type: "json" as const, text: JSON.stringify(result) }],
    };
  } catch (e) {
    return newErrorResult(e instanceof Error ? e.message : String(e));
  }
}
