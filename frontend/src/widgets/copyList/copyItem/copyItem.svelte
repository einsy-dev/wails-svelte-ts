<script lang="ts">
  import { Paste } from "$lib/wailsjs/go/clipboard";
  import { Hide } from "$lib/wailsjs/go/window";
  import { getContext } from "svelte";

  const { text = "" } = $props<{ text?: string }>();
  function callback(e: MouseEvent) {
    Hide();
    Paste(text);
  }
  const menu: { active: boolean; x: number; y: number } = getContext("menu");
  function context() {
    menu.active = true;
  }
</script>

<button class="bg-[rgba(255,255,255,0.1)] rounded px-1 h-7.5" onclick={callback} oncontextmenu={context}>
  <div class="text-nowrap overflow-hidden w-full">
    {JSON.stringify(text)}
  </div>
</button>
