import { GetHistory } from "$lib/wailsjs/go/clipboard/index.js";

export const prerender = true;
export const ssr = false;

export async function load({ fetch }) {
  const res = await GetHistory();
  return { data: res };
}
