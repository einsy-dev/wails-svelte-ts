import { GetHistory } from "$lib/wailsjs/go/clipboard/index.js";
import { PrintLn } from "$lib/wailsjs/go/window";

export const prerender = true;
export const ssr = false;

export async function load() {
  const res = await GetHistory();
  PrintLn(JSON.stringify(res[0]));
  return { data: res };
}
