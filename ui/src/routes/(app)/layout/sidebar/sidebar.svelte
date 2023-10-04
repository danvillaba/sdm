<script lang="ts">
	import SchemaList from './schemaList.svelte';
	import Selector from './selector.svelte';
</script>

<aside>
	<header>
		<Selector />
		<div class="tabs list md">
			<div class="item item-btn active">Schemas</div>
			<div class="item item-btn">Worksheets</div>
		</div>
	</header>
	<div class="content">
		<SchemaList />
	</div>
</aside>

<style>
	:global(.item) {
		--ListItem-radius: theme('borderRadius.lg');
	}
	header {
		--apply: space-y-2;
	}
	aside {
		--apply: h-full w-xs py-4 space-y-4;
		--apply: border-r-1 border-r-dashed bg-neutral-50;
		--apply: overflow-hiden flex flex-col;
		/* --apply */
	}
	aside header {
		--apply: px-4;
	}
	aside .content {
		--apply: overflow-y-scroll flex-1 px-4;
	}
</style>
