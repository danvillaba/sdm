<script lang="ts">
	import { page } from '$app/stores';
	import type { Axios } from 'axios';
	import type { ListTablesResponse } from 'pb/sdm/inspector/v1/inspector_service';

	export let schemaid: string;

	let api = $page.data['api'] as Axios;
	const tables = api
		.get<ListTablesResponse>(`/inspector/schemas/${schemaid}/tables`)
		.then((res) => res.data.tables || []);
</script>

{#await tables}
	loading {schemaid} tables
{:then tables}
	<ul class="list pl-4 border-l">
		{#each tables as item}
			<a
				href={`/schemas/${item.schema}/tables/${item.name}`}
				class="item item-btn soft truncate"
				class:primary={$page.params['tableid'] == item.name}
			>
				<div class="decorator">
					{#if item.type == 'view'}
						<i class="i-mdi-eye" />
					{:else if item.type == 'table'}
						<i class="i-mdi-table" />
					{:else}
						<i class="i-mdi-cube" />
					{/if}
				</div>
				{item.name}
			</a>
		{/each}
	</ul>
{/await}
