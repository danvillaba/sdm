<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<script lang="ts">
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<div class="overflow-auto flex-1">
	<table class="table md border-1">
		<thead>
			<th>Column</th>
			<th>Type</th>
			<th>Ordinal</th>
			<th>Nullable</th>
			<th>Comment</th>
		</thead>
		<tbody>
			{#each data.columns as column}
				<tr>
					<td title={column.name}>
						<i class="i-mdi-table-column" />
						{column.name}
					</td>
					<td>{column.data_type}</td>
					<td>{column.ordinal_position}</td>
					<td>{column.is_nullable ? 'Yes' : 'No'}</td>
					<td>{column.comment || ''}</td>
				</tr>
			{/each}
		</tbody>
	</table>
</div>
