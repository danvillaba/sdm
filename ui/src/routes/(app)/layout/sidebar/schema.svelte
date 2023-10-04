<script lang="ts">
	import { page } from '$app/stores';
	import Tables from './tables.svelte';

	export let name: string;

	let open = false;

	function handlePress() {
		open = !open;
	}
</script>

<button
	class="item item-btn soft"
	class:primary={$page.params['schemaid'] == name}
	on:click={handlePress}
>
	<div class="decorator">
		{#if open}
			<i class="i-mdi-chevron-down" />
		{:else}
			<i class="i-mdi-chevron-right" />
		{/if}
		<i class="i-mdi-file-tree" />
	</div>
	{name}
</button>
{#if open}
	<Tables schemaid={name} />
{/if}
