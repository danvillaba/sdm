<script lang="ts">
	import { page } from '$app/stores';
	import Sidebar from './layout/sidebar/sidebar.svelte';
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
	.root main {
		--apply: flex-1 overflow-y-auto;
	}
</style>
