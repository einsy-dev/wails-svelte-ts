<script lang="ts">
  import { Paste } from "$lib/wailsjs/go/clipboard";
  import { Hide } from "$lib/wailsjs/go/window";

  const { text = "" } = $props<{ text?: string }>();

  function callback(e: MouseEvent) {
    Hide();
    Paste(text);
  }
</script>

<button class="bg-[rgba(255,255,255,0.1)] rounded px-1 min-h-fit" onclick={callback}>
  <div class="text-nowrap overflow-hidden w-full">
    {text}
  </div>
</button>
