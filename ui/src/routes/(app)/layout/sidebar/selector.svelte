<script lang="ts">
	import { page } from '$app/stores';
	import type { Axios } from 'axios';
	import type { ListDatabasesResponse } from 'pb/sdm/inspector/v1/inspector_service';
	import type { InfoResponse } from 'pb/sdm/session/v1/session_service';
	import type { ChangeEventHandler } from 'svelte/elements';

	let session = $page.data['session'] as InfoResponse;
	let api = $page.data['api'] as Axios;

	let options = api.get<Required<ListDatabasesResponse>>(
		'/inspector/databases'
	);

	const handleChange: ChangeEventHandler<HTMLSelectElement> = (e) => {
		console.log(e.currentTarget.value);
	};
</script>

<select
	id="db-select"
	value={session.database}
	on:change={handleChange}
>
	{#await options}
		loading
	{:then res}
		{#each res.data.databases as item}
			<option value={item.name}>{item.name}</option>
		{/each}
	{/await}
</select>

<style>
	select {
		--apply: w-full p-2 rounded-lg bg-neutral-200;
		--apply: outline-none cursor-pointer;
	}
</style>
