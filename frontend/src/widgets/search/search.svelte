<script lang="ts">
  let { value = $bindable("") } = $props<{ value?: string; "bind:value"?: string }>();
</script>

<div class=" h-8 p-1 bg-[rgba(255,255,255,0.1)]">
  <input placeholder="Search" class="w-full text-white" bind:value />
</div>
