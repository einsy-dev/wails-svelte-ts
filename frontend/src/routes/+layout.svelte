<script lang="ts">
  import "../app.css";
  import { Search } from "../widgets";
  import CopyItem from "../shared/ui/copyItem/copyItem.svelte";
  import { Nav } from "../widgets";
  import { Hide, PrintLn } from "$lib/wailsjs/go/window";
  import { refreshAll } from "$app/navigation";
  import ContextMenu from "$widgets/contextMenu/contextMenu.svelte";

  let { children, data } = $props();
  let cMenu: any;
</script>

<svelte:window onfocus={() => refreshAll()} onblur={Hide} on:contextmenu|preventDefault={() => {}} />

<div class="max-h-screen overflow-hidden text-white flex flex-row *:w-1/2">
  <div class="flex flex-col">
    <Search />
    <ContextMenu bind:this={cMenu} />
    <button onclick={() => cMenu.test()}>fefef</button>
    <div class="p-1 overflow-x-hidden overflow-y-scroll flex flex-col gap-1">
      {#each data.data as el}
        <CopyItem text={el.Value} />
      {/each}
    </div>
  </div>
  <div class="flex flex-col">
    <Nav />
    {@render children?.()}
  </div>
</div>
