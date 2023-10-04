<script lang="ts">
	import type { Object } from 'pb/sdm/inspector/v1/object';
	import SidebarItem from './SidebarItem.svelte';

	let schemaList: Object[];
	let schemaId: string;
</script>

<aside class="root">
	{#if !schemaList}
		<h1>loading...</h1>
	{:else}
		<input
			type="text"
			class="input w-full search"
			placeholder="search"
		/>
		<ul class="list">
			{#each schemaList as item}
				<SidebarItem
					name={item.name || ''}
					active={item.name == schemaId}
				/>
			{/each}
		</ul>
	{/if}
</aside>

<style>
	:global(.item) {
		--ListItem-radius: theme('borderRadius.lg');
	}
	aside {
		--apply: h-full w-xs overflow-y-auto p-4 space-y-4;
		--apply: border-r-1 border-r-dashed bg-neutral-50;
	}
</style>
