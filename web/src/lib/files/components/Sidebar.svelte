<script>
	import { Icon } from "@steeze-ui/svelte-icon";
	import { Folder } from "@steeze-ui/lucide-icons";

	import { fileBrowser, areas } from "../stores/fileBrowser";

	import Sidebar from "$lib/Sidebar.svelte";
</script>

 <Sidebar>
	<ul class="menu">
		{#each $areas as area}
		<li class="mb-1">
			<a
				href={fileBrowser.getBrowserPath(area.name)}
				class="flex items-center"
				class:active={area.selected}
			>
				<Icon src={Folder} fill="oklch(var(--bc))" stroke="oklch(var(--bc))" class="w-6 h-6" />
				<div class="ml-2">
					{area.name}
				</div>
			</a>
		</li>
		{/each}
	</ul>
</Sidebar>
