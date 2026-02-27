<script lang="ts">
  import CopyItem from "./copyItem/copyItem.svelte";
  import { setContext } from "svelte";
  import ContextMenu from "./contextMenu/contextMenu.svelte";
  import { PrintLn } from "$lib/wailsjs/go/window";

  let { data } = $props();
  let menu = $state({
    active: false,
    x: 0,
    y: 0
  });
  $effect(() => {
    PrintLn(JSON.stringify(menu));
  });
  setContext("menu", menu);
</script>

<div class="relative p-1 overflow-x-hidden overflow-y-scroll flex flex-col gap-1 h-full">
  {#each data.data as el}
    <CopyItem text={el.Value} />
  {/each}
  {#if menu.active}
    <ContextMenu />
  {/if}
</div>
