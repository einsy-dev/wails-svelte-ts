<script lang="ts">
  import { CopyItem } from "$shared/ui";

  let { data } = $props();
</script>

<div class="p-1 overflow-x-hidden overflow-y-scroll flex flex-col gap-1">
  {#each data.data as el}
    <CopyItem text={el} />
  {/each}
</div>
