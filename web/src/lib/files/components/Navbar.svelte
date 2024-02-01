<script>
	import { Icon } from "@steeze-ui/svelte-icon";
	import { ArrowLeft, ArrowUp, ArrowRight, LayoutGrid, LayoutList } from "@steeze-ui/lucide-icons";

	import { upPath, zoom } from "../stores/fileBrowser";

	export let views, view;
</script>

<div class="bg-base-200 p-1 flex items-center">
	<button
		class="btn btn-square"
		on:click={() => {window.history.back();}}
	>
		<Icon src={ArrowLeft} class="w-6 h-6" />
	</button>
	<button
		class="btn btn-square"
		on:click={() => {window.history.forward();}}
	>
		<Icon src={ArrowRight} class="w-6 h-6" />
	</button>
	<a
		class="btn btn-square"
		class:btn-disabled={$upPath === null}
		href={$upPath}
	>
		<Icon src={ArrowUp} class="w-6 h-6" />
	</a>

	<div class="ml-auto" />

	<input
		type="range"
		min="20"
		max="100"
		class="range range-xs w-48 mr-8"
		bind:value={$zoom}
	/>

	<button
		class="btn btn-square"
		on:click={() => {view=views.gallery;}}
	>
		<Icon src={LayoutGrid} stroke-width="2" class="w-6 h-6" />
	</button>

	<button
		class="btn btn-square"
		on:click={() => {view=views.table;}}
	>
		<Icon src={LayoutList} class="w-6 h-6" />
	</button>
</div>
