<script>
	import Sidebar from './AppSidebar.svelte';
</script>

<div class="root">
	<Sidebar />
	<main>
		<slot />
	</main>
</div>

<style>
	.root {
		--apply: flex h-screen overflow-hidden;
	}
	main {
		--apply: flex-1 overflow-y-auto;
	}
</style>
