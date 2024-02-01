<script>
	import { Icon } from "@steeze-ui/svelte-icon";
	import { X, Download, Menu, ArrowLeft, ArrowRight } from "@steeze-ui/lucide-icons";

	import { upPath, prevFile, nextFile, file } from "../stores/fileBrowser";

	import FileIcon from "./FileIcon.svelte";

	import ImagePlayer from "./players/ImagePlayer.svelte";
	import TextPlayer from "./players/TextPlayer.svelte";
	import VideoPlayer from "./players/VideoPlayer.svelte";

	let player = null;
	let detailsVisible = true;

	$: {
		if (["text", "code", "shell"].includes($file.category)) {
			player = TextPlayer;
		} else if ($file.category === "image") {
			player = ImagePlayer;
		} else if ($file.category === "video") {
			player = VideoPlayer;
		} else {
			player = "unknown";
		}
	}
</script>

<div class="h-screen flex flex-col">
	<div class="p-4 bg-base-200">
		<div class="flex">
			<div class="flex-none self-center mr-2">
				<FileIcon file={$file} class="h-10 w-10" />
			</div>
			<div class="shrink min-w-0 self-center">
				<h2 class="text-md break-all min-w-64">
					{$file.name}
				</h2>
				<small class="text-xs">
					{$file.path.rendered}
				</small>
			</div>
			<div class="flex-none ml-auto">
				<a
					href={$upPath}
					class="btn btn-square"
				>
					<Icon
						src={X}
						class="w-6 h-6"
					/>
				</a>
			</div>
		</div>

		<div class="flex pt-4 justify-between">
			<a
				class="btn btn-square mr-2"
				href={$file.downloadURL}
			>
				<Icon src={Download} class="w-6 h-6" />
			</a>

			<div>
				<a
					class="btn btn-square"
					class:btn-disabled={$prevFile === null}
					href={$prevFile}
				>
					<Icon src={ArrowLeft} class="w-6 h-6" />
				</a>
				<a
					class="btn btn-square"
					class:btn-disabled={$nextFile === null}
					href={$nextFile}
				>
					<Icon src={ArrowRight} class="w-6 h-6" />
				</a>
			</div>

			<button
				class="btn btn-square"
				on:click={() => {detailsVisible = !detailsVisible;}}
			>
				<Icon src={Menu} class="w-6 h-6" />
			</button>
		</div>
	</div>

	<div class="grow min-h-0 h-full">
		<div class="flex h-full">
			<div class="flex-1 flex">
				{#if player !== null && player !== "unknown"}
					<svelte:component this={player} file={$file} />
				{:else}
					<div class="self-center mx-auto">No preview available for this file type</div>
				{/if}
			</div>

			<div
				class="flex-none h-full bg-base-200 w-64 p-4"
				class:hidden={!detailsVisible}
			>
				<div class="flex flex-col gap-2 break-all">
					<div>
						<div class="font-bold">Name</div>
						<div class="">
							{$file.name}
						</div>
					</div>

					<div>
						<div class="font-bold">Size</div>
						<div>
							{$file.sizeHumanReadable}
						</div>
					</div>

					<div>
						<div class="font-bold">Accessed at</div>
						<div>
							{$file.atime.format()}
						</div>
					</div>

					<div>
						<div class="font-bold">Modified at</div>
						<div>
							{$file.mtime.format()}
						</div>
					</div>

					<div>
						<div class="font-bold">Created at</div>
						<div>
							{$file.ctime.format()}
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>
