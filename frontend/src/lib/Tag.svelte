<script lang="ts">
    export let data;
    export let color;
</script>
  

<span style="background-color: {color}">{data}</span>


<style>
	span {
		display: inline-block;
		padding: .2em 1em;
		margin: 0 0.5em 0 0;
		text-align: center;
		border-radius: 1em;
		color: rgb(145, 145, 145);
	}
</style>