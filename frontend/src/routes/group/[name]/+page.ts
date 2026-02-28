import { GetByGroup } from "$lib/wailsjs/go/clipboard";

export const prerender = false;

export async function load({ params }) {
  const data = await GetByGroup(params.name);
  return { data };
}
