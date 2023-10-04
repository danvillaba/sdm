<script lang="ts">
	import type { PageData } from './$types';

	export let data: PageData;
	let rows = data.rows || [];
	let columns = data.columns;
</script>

<table class="table md border-1">
	<thead>
		{#each columns as column}
			<th>{column.name}</th>
		{/each}
	</thead>
	<tbody>
		{#each rows as row}
			<tr>
				{#each row as value}
					<td title={value}>{value}</td>
				{/each}
			</tr>
		{/each}
	</tbody>
</table>
