<script lang="ts">
  import { Paste } from "$lib/wailsjs/go/clipboard";
  import { Hide, PrintLn } from "$lib/wailsjs/go/window";
  import { clickOutside } from "$shared/utils/clickOutside";
  import { contextMenu } from "$store";

  const { text = "" } = $props<{ text?: string }>();

  function callback(e: MouseEvent) {
    Hide();
    Paste(text);
  }
  function context() {
    contextMenu.update((state) => {
      state.active = true;
      return state;
    });
  }
  let button: HTMLButtonElement;
</script>

<button
  bind:this={button}
  class="bg-[rgba(255,255,255,0.1)] rounded px-1 h-[30px]"
  onclick={callback}
  oncontextmenu={context}
>
  <div class="text-nowrap overflow-hidden w-full">
    {JSON.stringify(text)}
  </div>
</button>
