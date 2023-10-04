<script lang="ts">
	import PageHeader from 'components/PageHeader.svelte';
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<div class="flex flex-col h-full">
	<PageHeader breadcrumbs={data.breadcrumbs}>
		<svelte:fragment slot="actions" />
	</PageHeader>

	<div class="overflow-auto flex-1 p-4">
		<table class="table md border-1">
			<thead>
				<th>Name</th>
				<th>Type</th>
			</thead>
			<tbody>
				{#each data.tables as table}
					<tr>
						<td>{table.name}</td>
						<td>{table.type}</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</div>
</div>
