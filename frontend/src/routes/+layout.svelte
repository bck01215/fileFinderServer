<script lang="ts">
	import Nav from '../components/nav.svelte';
</script>

<title>File Finder</title>

<Nav />
<slot />

<style lang="scss">
	:global(body) {
		background-color: #ffffff;
		color: #000000;
		margin: 0;
	}
</style>
