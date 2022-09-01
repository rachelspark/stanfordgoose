<script lang="ts">
    import { encodeQueryHash } from "../App.svelte";
  
    export let query: string; // for binding
    export let value: string;
  
</script>
  
  <a
    class="text-rose-600 hover:underline"
    href={encodeQueryHash(value)}
    on:click|preventDefault={() => (query = value)}>{value}</a
  >
  