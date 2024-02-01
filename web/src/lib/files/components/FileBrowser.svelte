<script>
	import { Pencil, Trash2, Globe } from "@steeze-ui/lucide-icons";

	import { contextMenu } from "$lib/ContextMenu/store";

	import { state, file, zoom } from "../stores/fileBrowser";

	import FileTable from "./FileTable.svelte";
	import FileGallery from "./FileGallery.svelte";
	import FileDetail from "./FileDetail.svelte";
	import Navbar from "./Navbar.svelte";
	import Sidebar from "./Sidebar.svelte";

	let views = {
		gallery: {
			comp: FileGallery,
		},
		table: {
			comp: FileTable,
		},
	};
	let view = views.gallery;
	let container;

	$: {
		if ($state.event == "files-updated" && container) {
			container.scrollTop = 0;
		}
	}

	contextMenu.setItems([
		{
			label: "Open in new tab",
			icon: Globe,
		},
		{
			label: "Rename",
			icon: Pencil,
		},
		{
			label: "Delete",
			icon: Trash2,
		},
	]);
</script>

{#if $file !== null}
	{#if $file.isDir}
		<div class="h-screen w-screen overflow-hidden">
			<div
				class="flex"
				style="--browser-zoom: {$zoom}"
			>
				<div class="flex-none">
					<Sidebar />
				</div>

				<div class="
					flex-1 min-w-0
					h-screen
					flex flex-col
				">
					<Navbar views={views} bind:view={view} />
					<div
						class="overflow-auto h-full"
						bind:this={container}
					>
						<svelte:component this={view.comp} />
					</div>
				</div>
			</div>
		</div>
	{:else}
		<FileDetail />
	{/if}
{/if}
