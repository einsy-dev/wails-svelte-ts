<script lang="ts">
  import "../app.css";
  import { Search } from "../widgets";
  import CopyItem from "../shared/ui/copyItem/copyItem.svelte";
  import { Nav } from "../widgets";
  import { Hide } from "$lib/wailsjs/go/window";
  let { children } = $props();

  let data = Array(10).fill({
    type: "string",
    value: "h-screen text-white flex bg-amber-700"
  });
  window.onblur = () => {
    Hide();
  };
</script>

<div class="max-h-screen overflow-hidden text-white flex flex-row *:w-1/2">
  <div class="flex flex-col">
    <Search />
    <div class="p-1 overflow-x-hidden overflow-y-scroll flex flex-col gap-1">
      {#each data as el}
        <CopyItem
          text={el.value}
          onclick={() => {
            console.log("click!");
          }}
        />
      {/each}
    </div>
  </div>
  <div class="flex flex-col">
    <Nav />
    {@render children?.()}
  </div>
</div>
