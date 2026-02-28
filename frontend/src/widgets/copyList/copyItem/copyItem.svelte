<script lang="ts">
  import { Paste } from "$lib/wailsjs/go/clipboard";
  import type { models } from "$lib/wailsjs/go/models";
  import { Hide, PrintLn } from "$lib/wailsjs/go/window";
  import { getContext } from "svelte";

  const { data }: { data: models.Text } = $props<{ data: models.Text }>();
  function callback(e: MouseEvent) {
    Hide();
    Paste(data.Value);
  }
  const menu: { active: boolean; id: number; x: number; y: number } = getContext("menu");

  function context(e: MouseEvent) {
    menu.x = e.pageX;
    menu.y = e.pageY;
    menu.active = true;
    menu.id = data.ID;
  }
</script>

<button class="bg-[rgba(255,255,255,0.1)] rounded px-1 h-7.5" onclick={callback} oncontextmenu={context}>
  <div class="text-nowrap overflow-hidden w-full">
    {data.Value}
  </div>
</button>
