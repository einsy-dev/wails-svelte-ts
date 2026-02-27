<script lang="ts">
  import "../app.css";
  import { Search } from "../widgets";
  import { Nav } from "../widgets";
  import { Hide } from "$lib/wailsjs/go/window";
  import { refreshAll } from "$app/navigation";
  import CopyList from "$widgets/copyList/copyList.svelte";

  let { children, data } = $props();
  let cMenu: any;
</script>

<svelte:window onfocus={() => refreshAll()} onblur={Hide} on:contextmenu|preventDefault={() => {}} />

<div class="max-h-screen text-white flex flex-col">
  <div class="flex *:w-1/2">
    <Search />
    <Nav />
  </div>

  <div class="flex *:w-1/2 overflow-hidden">
    <div class="">
      <CopyList {data} />
    </div>
    <div class="">
      {@render children?.()}
    </div>
  </div>
</div>
