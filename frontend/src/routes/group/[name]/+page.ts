import { GetGroup } from "$lib/wailsjs/go/clipboard";

export const prerender = false;

export async function load({ params }) {
  const res = await GetGroup(params.name);
  return { data: res };
}
