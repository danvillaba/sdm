<script lang="ts">
	import { page } from '$app/stores';
	import PageHeader from 'components/PageHeader.svelte';
	import type { PageData } from './$types';

	export let data: PageData;

	function tabUrl(params: any, url: string) {
		const schemaid = params['schemaid'];
		const tableid = params['tableid'];
		return `/schemas/${schemaid}/tables/${tableid}${url}`;
	}
</script>

<div class="flex flex-col h-full">
	<PageHeader breadcrumbs={data.breadcrumbs}>
		<svelte:fragment slot="actions">
			<button class="btn soft primary md">Alter</button>
		</svelte:fragment>
	</PageHeader>
	<div class="tabs list sm">
		<a
			href={tabUrl($page.params, '')}
			class="item"
			class:active={tabUrl($page.params, '') == $page.url.pathname}
			>Structure</a
		>
		<a
			href={tabUrl($page.params, '/select')}
			class="item"
			class:active={tabUrl($page.params, '/select') ==
				$page.url.pathname}>Preview</a
		>
	</div>
	<div class="overflow-auto p-4">
		<slot />
	</div>
</div>
