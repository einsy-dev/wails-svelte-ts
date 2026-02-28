<script lang="ts">
  import { goto } from "$app/navigation";
  import { Eraser, Heart, KeyRound, TimerReset } from "@lucide/svelte";
</script>

<nav class="h-8 p-1 bg-[rgba(255,255,255,0.1)]">
  <ul class="flex gap-2">
    <li>
      <button class="hover:cursor-pointer" onclick={() => goto(`/group/favorites`)}><Heart /></button>
    </li>
    <li><button class="hover:cursor-pointer" onclick={() => goto("/group/passwords")}><KeyRound /></button></li>
    <li><button class="hover:cursor-pointer" onclick={() => goto("/totp/main")}><TimerReset /></button></li>
  </ul>
</nav>
