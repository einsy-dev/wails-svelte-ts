<script lang="ts">
  import { CopyItem } from "../shared/ui";

  let data = Array(10).fill({
    type: "string",
    value: "h-screen text-white flex bg-amber-700"
  });
</script>

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
