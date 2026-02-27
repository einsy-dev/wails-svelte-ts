<script lang="ts">
  import { PrintLn } from "$lib/wailsjs/go/window";
  import { clickOutside } from "$shared/utils/clickOutside";
  import { getContext } from "svelte";
  const menu: { active: boolean; x: number; y: number } = getContext("menu");
</script>

<div class="absolute" use:clickOutside onclick_outside={() => (menu.active = false)}>
  <div>Context</div>
</div>
