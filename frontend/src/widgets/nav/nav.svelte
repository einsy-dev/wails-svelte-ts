<script lang="ts">
  import { goto } from "$app/navigation";
  import { Heart, KeyRound } from "@lucide/svelte";
</script>

<nav class="h-8 p-1 bg-[rgba(255,255,255,0.1)]">
  <ul class="flex gap-2">
    <li>
      <button class="hover:cursor-pointer" onclick={() => goto(`/group`)}><Heart /></button>
    </li>
    <li><button class="hover:cursor-pointer" onclick={() => goto("/")}><KeyRound /></button></li>
  </ul>
</nav>
