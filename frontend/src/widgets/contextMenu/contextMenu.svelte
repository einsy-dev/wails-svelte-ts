<script lang="ts">
  import { PrintLn } from "$lib/wailsjs/go/window";
  import { clickOutside } from "$shared/utils/clickOutside";
  import { contextMenu } from "$store";

  let active = $state(false);

  contextMenu.subscribe((state) => {
    active = state.active;
  });
</script>

<div
  class=""
  use:clickOutside
  onclick_outside={() => {
    PrintLn("Out");
  }}
>
  {#if active}
    <div>Context</div>
  {/if}
</div>
