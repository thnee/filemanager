<script>
	import { goto } from "$app/navigation";
	import { apiAuth } from "$lib/apiAuth";

	let errors = [];

	async function login(event) {
		event.preventDefault();
		let r = await apiAuth.login({
			username: event.target.elements.username.value,
			password: event.target.elements.password.value,
		});
		if (r.status == 200) {
			console.log("going");
			goto("/files");
		}
		if (r.status == 400) {
			let data = await r.json();
			errors = data.errors;
		}
	}
</script>

<div
	class="
		card card-compact
		bg-base-200 shadow-md
		w-96 mx-auto
		p-4 mt-20
	"
>
	<h2 class="text-2xl mb-4">Login</h2>
	<form on:submit={login}>
		<label
			for="username"
			class="block mt-2 mb-1"
		>
			Username
		</label>
		<input
			id="username"
			type="text"
			class="input input-bordered w-full"
		/>

		<label
			for="password"
			class="block mt-2 mb-1"
		>
			Password
		</label>
		<input
			id="password"
			type="password"
			class="input input-bordered w-full"
		/>

		<input
			type="submit"
			value="Login"
			class="btn btn-neutral rounded mt-4"
		/>
	</form>

	{#if errors}
		<div class="text-red-400 mt-4">
			{#each errors as error}
				{error}
			{/each}
		</div>
	{/if}
</div>
