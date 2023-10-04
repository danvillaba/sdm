<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
<script lang="ts">
	import { goto } from '$app/navigation';
	import type { Table } from 'pb/sdm/table';
	import type { GetSchemaTablesResponse } from 'pb/sdm/v1/service';
	import api from 'provider/api';

	export let name: string;
	let open = false;
	let items: Table[];
	export let active: boolean;

	async function schemaClick() {
		if (open) {
			if (!active) return goto(`/schemas/${name}`);
			return (open = !open);
		}

		if (!active) {
			goto(`/schemas/${name}`);
		}

		const res = await api<GetSchemaTablesResponse>(
			`/schemas/${name}/tables`
		);

		items = res.data.tables || [];
		open = true;
	}
</script>

<li
	class="item item-btn"
	class:bg-blue-100={active}
	class:text-blue-800!={active}
	on:click={schemaClick}
>
	<!-- arrow icon -->
	{#if !open}
		<i class="i-mdi-chevron-right" />
	{:else}
		<i class="i-mdi-chevron-down" />
	{/if}
	<!-- database icon -->
	<div class="decorator">
		<i class="i-mdi-database" />
	</div>
	<!-- content -->
	{name}
</li>

{#if open}
	<ul class="list ml-4! border-l-1">
		{#if !items.length}
			<h1>without tables</h1>
		{:else}
			{#each items as item}
				<a
					class="item item-btn pl-6! truncate"
					href={`/schemas/${name}/tables/${item.name}`}
				>
					<div class="decorator">
						<i class="i-mdi-table-large" />
					</div>
					{item.name}
				</a>
			{/each}
		{/if}
	</ul>
{/if}
