<script>
	import { Icon } from "@steeze-ui/svelte-icon";
	import { onMount, onDestroy } from "svelte";

	import { contextMenu, contextMenuItems, contextMenuVisible, contextMenuLeft, contextMenuTop } from "./store";

	onMount(() => {
		window.addEventListener("mousedown", onWindowMousedown);
		window.addEventListener("keydown", onWindowKeydown);
		window.addEventListener("blur", onWindowBlur);
	});

	onDestroy(() => {
		window.removeEventListener("mousedown", onWindowMousedown);
		window.removeEventListener("keydown", onWindowKeydown);
		window.removeEventListener("blur", onWindowBlur);
	});

	function onWindowMousedown() {
		contextMenu.close();
	}

	function onWindowBlur() {
		contextMenu.close();
	}

	function onWindowKeydown(event) {
		if (event.key === "Escape") {
			contextMenu.close();
		}
	}
</script>

<div
	class:hidden={!$contextMenuVisible}
	class="absolute bg-base-200 border-solid border py-3 w-48 rounded-lg border-neutral-500 shadow"
	style={`left: ${$contextMenuLeft}px; top: ${$contextMenuTop}px;`}
>
	{#each $contextMenuItems as item}
		<div class="flex items-center hover:bg-base-300 p-2 cursor-default">
			<Icon src={item.icon} class="w-5 h-5 ml-2 mr-3" />
			{item.label}
		</div>
	{/each}
</div>
