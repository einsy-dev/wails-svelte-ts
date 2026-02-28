<script lang="ts">
  import { refreshAll } from "$app/navigation";
  import { SetGroup } from "$lib/wailsjs/go/clipboard";
  import { PrintLn } from "$lib/wailsjs/go/window";
  import { clickOutside } from "$shared/utils/clickOutside";
  import { getContext } from "svelte";
  const menu: { active: boolean; id: number; x: number; y: number } = getContext("menu");
</script>

<div class="bg-black flex flex-col px-2 py-1 rounded" use:clickOutside onclick_outside={() => (menu.active = false)}>
  <button
    class="cursor-pointer"
    onclick={() => {
      SetGroup(menu.id, "favorites");
      menu.active = false;
      refreshAll();
    }}>Pin</button
  >
  <button
    class="cursor-pointer"
    onclick={() => {
      SetGroup(menu.id, "passwords");
      menu.active = false;
      refreshAll();
    }}>Add to ...</button
  >
  <button
    class="cursor-pointer"
    onclick={() => {
      SetGroup(menu.id, "history");
      menu.active = false;
      refreshAll();
    }}>Delete</button
  >
</div>
