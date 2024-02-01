<script>
	import { Icon } from "@steeze-ui/svelte-icon";
	import { Box, Settings, HelpCircle, LogOut } from "@steeze-ui/lucide-icons";
	import { apiAuth } from "$lib/apiAuth";
</script>

<div class="w-64 flex flex-col sticky top-0 h-screen bg-base-200">
	<div class="pt-4 pl-5 pb-2 pr-4 select-none">
		<a href="/" class="flex items-center">
			<Icon src={Box} class="w-8 h-8" />
			<span class="text-lg ml-3">
				My Box
			</span>
		</a>
	</div>

  <slot />

	<div class="mt-auto flex gap-2 pl-5 mb-2">
		<a href="/settings" class="btn btn-ghost btn-square">
			<Icon src={Settings} class="w-6" />
		</a>

    <a href="/settings" class="btn btn-ghost btn-square">
			<Icon src={HelpCircle} class="w-6" />
		</a>

		<button
			class="btn btn-ghost btn-square"
			on:click={apiAuth.logout}
		>
			<Icon src={LogOut} class="w-6" />
		</button>
	</div>
</div>
