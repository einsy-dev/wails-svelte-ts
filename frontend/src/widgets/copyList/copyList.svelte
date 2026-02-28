<script lang="ts">
  import CopyItem from "./copyItem/copyItem.svelte";
  import { setContext } from "svelte";
  import ContextMenu from "./contextMenu/contextMenu.svelte";
  import { PrintLn } from "$lib/wailsjs/go/window";
  import type { models } from "$lib/wailsjs/go/models";

  let { data }: { data: models.Text[] } = $props<{ data: models.Text[] }>();
  let menu = $state({
    active: false,
    x: 0,
    y: 0,
    id: null
  });

  setContext("menu", menu);
</script>

<div
  class="relative p-1 pb-2 overflow-scroll flex flex-col gap-1 h-full"
  onscroll={() => {
    menu.active = false;
  }}
>
  {#each data as el}
    <CopyItem data={el} />
  {/each}
</div>

<div class="absolute" style:left="{menu.x}px" style:top="{menu.y}px">
  {#if menu.active}
    <ContextMenu />
  {/if}
</div>
