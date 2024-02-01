<script>
	import { onMount } from "svelte";

	export let file;

	let canPlay = false;

	onMount(async () => {
		let tempVideo = document.createElement("video");
		let canPlayType = tempVideo.canPlayType(file.mimeType);
		if (canPlayType == "probably") {
			canPlay = true;
		}
	});
</script>

<div class="self-center mx-auto">
	{#if canPlay}
		<video
			class="max-h-full max-w-full object-contain self-center mx-auto"
			controls
			preload="auto"
		>
			<source
				src={file.dataURL}
				type={file.mimeType}
			/>
			<track kind="captions" />
		</video>
	{:else}
		Cannot play this video type
	{/if}
</div>
