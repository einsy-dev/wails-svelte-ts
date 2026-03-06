<script lang="ts">
  import "../app.css";
  let { children } = $props();
</script>

<div class="max-h-screen text-white">
  {@render children?.()}
</div>
