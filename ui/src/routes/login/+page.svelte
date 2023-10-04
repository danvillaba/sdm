<script>
	import { handleSubmit } from './actions';
</script>

<div class="root">
	<form autocomplete="off" on:submit|preventDefault={handleSubmit}>
		<div>
			<label for="">Engine</label>
			<select class="select neutral soft" name="scheme" required>
				<option value="postgres">Postgres</option>
				<option value="mysql">Mysql</option>
			</select>
		</div>
		<div class="flex gap-2">
			<div class="w-full">
				<label for="">Host</label>
				<input class="input" type="text" name="host" required />
			</div>
			<div>
				<label for="">Port</label>
				<input class="input" type="number" name="port" />
			</div>
		</div>
		<div>
			<label for="">Username</label>
			<input class="input" type="text" name="username" />
		</div>
		<div>
			<label for="">Password</label>
			<input class="input" type="password" name="password" />
		</div>
		<div>
			<label for="">Database</label>
			<input class="input" type="text" name="database" />
		</div>
		<button type="submit" class="btn neutral soft">Connect</button>
	</form>
</div>

<style>
	.root {
		--apply: h-screen w-full flex items-center justify-center;
	}
	.root form {
		--apply: m-auto min-w-md bg-neutral-100 flex flex-col gap-2;
		--apply: p-8 rounded-lg;
	}
	.root input,
	.root select {
		--apply: w-full;
	}
</style>
