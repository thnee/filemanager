<script>
	import { onMount, onDestroy } from "svelte";

	import { contextMenu } from "$lib/ContextMenu/store";

	import { fileBrowser, files } from "../stores/fileBrowser";

	import FileIcon from "./FileIcon.svelte";

	export let columns;

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

<tbody>
	{#each $files as file}
		<tr
			class="
				file
				{file.selected? "bg-neutral-500/50" : "hover:bg-neutral-500/20"}
			"
			on:click={(event) => onFileClick(event, file)}
			on:dblclick={(event) => onFileDblClick(event, file)}
			on:contextmenu={(event) => onFileContextmenu(event, file)}
		>
			{#each columns as column}
				<td
					class="
						float-left
						relative
						flex items-center
						px-2
					"
					style="
						width: calc(var(--browser-zoom) * {column.width} * 0.025px);
						height: calc(var(--browser-zoom) * 1px);
						font-size: calc(var(--browser-zoom) * 0.4px)
					"
				>
					<div class="overflow-x-hidden whitespace-nowrap text-ellipsis">
						{#if column.name == "icon"}
							<FileIcon file={file} />
						{/if}
						{#if column.name == "name"}
							{file.name}
						{/if}
						{#if column.name == "size"}
							{file.sizeHumanReadable}
						{/if}
						{#if column.name == "atime"}
							{file.atime.format("YYYY-MM-DD HH:mm:ss")}
						{/if}
						{#if column.name == "mtime"}
							{file.mtime.format("YYYY-MM-DD HH:mm:ss")}
						{/if}
						{#if column.name == "ctime"}
							{file.ctime.format("YYYY-MM-DD HH:mm:ss")}
						{/if}
					</div>
				</td>
			{/each}
		</tr>
	{/each}
</tbody>
