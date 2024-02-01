<script>
	import { goto } from "$app/navigation";
	import { apiAuth } from "$lib/apiAuth";

	async function login(event) {
		event.preventDefault();
		await apiAuth.login({
			username: event.target.elements.username.value,
			password: event.target.elements.password.value,
		});
		goto("/files");
	}
</script>

<div class="
	card card-compact
	bg-base-200 shadow-md
	w-96 mx-auto
	p-4 mt-20
	"
>
	<h2 class="text-2xl mb-4">Login</h2>
	<form
		on:submit={login}
	>
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
			class="btn mt-4"
		/>
	</form>
</div>
