<script>
	import FileBrowser from "$lib/files/components/FileBrowser.svelte";
	import { fileBrowser } from "$lib/files/stores/fileBrowser";

	export let data;

	$: fileBrowser.init(data.path);
</script>

<FileBrowser />
