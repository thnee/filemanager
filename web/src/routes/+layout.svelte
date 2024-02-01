<script>
	import "../app.css";

	import { onMount } from "svelte";
	import { themeChange } from "theme-change";

	import ContextMenu from "$lib/ContextMenu/ContextMenu.svelte";

	onMount(() => {
		themeChange(false);
	});
</script>

<slot />

<ContextMenu />
