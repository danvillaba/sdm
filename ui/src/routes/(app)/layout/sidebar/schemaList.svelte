<script lang="ts">
	import { page } from '$app/stores';
	import type { Axios } from 'axios';
	import type { ListSchemasResponse } from 'pb/sdm/inspector/v1/inspector_service';
	import Schema from './schema.svelte';

	let api: Axios = $page.data['api'];

	let schemas = api
		.get<ListSchemasResponse>('/inspector/schemas')
		.then((res) => res.data.schemas || []);
</script>

{#await schemas}
	loading schemas
{:then schemas}
	<ul class="list sm">
		{#each schemas as item}
			<Schema name={item.name || ''} />
		{/each}
	</ul>
{/await}

<style>
	:global(.item) {
		--ListItem-radius: theme('borderRadius.lg');
	}
</style>
