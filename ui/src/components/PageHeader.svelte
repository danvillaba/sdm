<script lang="ts">
	export let breadcrumbs: string[] = [];
</script>

<header>
	<div class="breadcrumb">
		<div class="decorator">
			<i class="i-mdi-table" />
		</div>
		<p>{breadcrumbs.join(' / ')}</p>
	</div>
	<div class="actions">
		<slot name="actions" />
	</div>
</header>

<style>
	header {
		--apply: p4 flex items-center justify-between;
	}

	header .breadcrumb {
		--apply: flex items-center text-xl gap-4;
	}
	header .breadcrumb .decorator {
		--apply: w-10 h-10 flex items-center justify-center;
		--apply: p-2 rounded-lg bg-neutral-100;
	}

	header .actions {
		--apply: flex items-center gap-2;
	}
</style>
