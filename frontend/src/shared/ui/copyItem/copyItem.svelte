<script lang="ts">
  import { Hide } from "$lib/wailsjs/go/window";

  const { text = "", onclick = () => {} } = $props<{ text?: string; onclick?: () => void }>();

  function callback(e: MouseEvent) {
    onclick(e);
    Hide();
  }
</script>

<button class="bg-[rgba(255,255,255,0.1)] rounded px-1 min-h-fit" onclick={callback}>
  <div class="text-nowrap overflow-hidden w-full">
    {text}
  </div>
</button>
