import { GetGroup } from "$lib/wailsjs/go/clipboard";

export async function load({ fetch }) {
  const res = await GetGroup("favorite");
  return { data: res };
}
