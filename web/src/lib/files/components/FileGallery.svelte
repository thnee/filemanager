<script>
	import { onMount, onDestroy } from "svelte";

	import { contextMenu } from "$lib/ContextMenu/store";

	import { fileBrowser, files } from "../stores/fileBrowser";

	import FileIcon from "./FileIcon.svelte";

	onMount(() => {
		window.addEventListener("click", onWindowClick);
	});

	onDestroy(() => {
		window.removeEventListener("click", onWindowClick);
	});

	function onWindowClick(event) {
		let closestFileEl = event.target.closest(".file");
		if (closestFileEl === null) {
			fileBrowser.unselectAllFiles();
		}
	}

	function onFileClick(event, file) {
		event.preventDefault();
		fileBrowser.selectOneFile(file);
	}

	function onFileDblClick(event, file) {
		event.preventDefault();
		file.goto();
	}

	function onFileContextmenu(event, eventFile) {
		event.preventDefault();
		fileBrowser.selectOneFile(eventFile);
		contextMenu.open(event.clientX, event.clientY);
	}
</script>

<div class="flex flex-wrap justify-left gap-1 p-2">
{#each $files as file, index}
	<div
		class="flex flex-col file-item"
		data-index="{index}"
	>
		<button
			class="
				file
				cursor-default
				select-none
				p-1
				{file.selected? "bg-neutral-500/50" : "hover:bg-neutral-500/20"}
			"
			style="
				width:  calc(var(--browser-zoom) * 2.5px);
			"
			on:click={(event) => onFileClick(event, file)}
			on:dblclick={(event) => onFileDblClick(event, file)}
			on:contextmenu={(event) => onFileContextmenu(event, file)}
		>
			<div
				class="flex justify-center"
				style="height: calc(var(--browser-zoom) * 2px);"
			>
				{#if file.category == "image"}
					<div
						style="background-image: url({file.dataURL});"
						class="w-full h-full bg-no-repeat bg-contain bg-center"
					/>
				{:else}
					<FileIcon file={file} />
				{/if}
			</div>
			<div class="text-center break-all line-clamp-5 mt-1">
				{file.name}
			</div>
		</button>
	</div>
{/each}
</div>
