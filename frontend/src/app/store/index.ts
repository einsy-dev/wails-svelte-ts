import { writable } from "svelte/store";

export const contextMenu = writable({
  active: false
});
