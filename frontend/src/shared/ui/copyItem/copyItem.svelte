<script lang="ts">
  const { text = "", onclick = () => {} } = $props<{ text?: string; onclick?: () => void }>();
</script>

<button class="bg-[rgba(255,255,255,0.1)] rounded px-1 min-h-fit" {onclick}>
  <div class="text-nowrap overflow-hidden w-full">
    {text}
  </div>
</button>
