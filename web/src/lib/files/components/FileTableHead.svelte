<script>
	import { onMount, onDestroy } from "svelte";

	import { Icon } from "@steeze-ui/svelte-icon";
	import { Sparkle } from "@steeze-ui/lucide-icons";

	import { debounce } from "$lib/utils";

	export let columns;

	let startPos = 0;
	let startWidth = 0;
	let isResizing = false;
	let isMoving = false;
	let isHoldingColumn = false;
	let curIndex = "";
	let tr = null;

	onMount(() => {
		window.addEventListener("mousemove", onWindowMousemove);
		window.addEventListener("mouseup", onWindowMouseup);
	});

	onDestroy(() => {
		window.removeEventListener("mousemove", onWindowMousemove);
		window.removeEventListener("mouseup", onWindowMouseup);
	});

	function onResizerMousedown(event, index) {
		event.preventDefault();
		event.stopPropagation();
		isResizing = true;
		curIndex = index;
		startPos = event.clientX;
		startWidth = columns[index].width;
	}

	function onColumnMousedown(event, index) {
		event.preventDefault();
		event.stopPropagation();

		isHoldingColumn = true;
		curIndex = index;
	}

	function onWindowMousemove(event) {
		if (isResizing) {
			event.preventDefault();
			setResizingPosition(event);
		}

		if (isHoldingColumn) {
			event.preventDefault();
			isMoving = true;

			if (tr.querySelector("td.moving") == null) {
				let td = tr.querySelector(`td[data-index="${curIndex}"]`);
				let clone = td.cloneNode(true);
				clone.classList.add("moving");
				let rect = td.getBoundingClientRect();
				clone.style.height = `${rect.height}px`;
				clone.style.width = `${rect.width}px`;
				tr.appendChild(clone);
			}

			setMovingPosition(event);
		}
	}

	function setResizingPosition(event) {
		if (event.clientX === 0) {
			return false;
		}
		debounce(() => {
			let width = startWidth - (startPos - event.clientX);
			columns[curIndex].width = Math.max(width, 30);
		}, 10)();
	}

	function setMovingPosition(event) {
		let trRect = tr.getBoundingClientRect();
		let td = tr.querySelector("td.moving");
		let left = event.clientX - trRect.left - 50;
		td.style.left = `${left}px`;
	}

	function onWindowMouseup(event) {
		if (isResizing) {
			event.preventDefault();
			isResizing = false;
			let width = startWidth - (startPos - event.clientX);
			columns[curIndex].width = Math.max(width, 30);
		}

		if (isHoldingColumn && !isMoving) {
			event.preventDefault();
			isHoldingColumn = false;
		}

		if (isHoldingColumn && isMoving) {
			event.preventDefault();
			isHoldingColumn = false;
			isMoving = false;

			let td = tr.querySelector("td.moving");
			td.remove();

			let targetIndex = null;
			let tds = tr.querySelectorAll("td:not(.moving)");
			tds.forEach((td) => {
				let rect = td.getBoundingClientRect();
				if (event.clientX > rect.left) {
					targetIndex = td.dataset.index;
				}
			});
			let sourceCols = columns.splice(curIndex, 1);
			columns = columns;
			columns.splice(targetIndex, 0, sourceCols[0]);
			columns = columns;
		}
	}
</script>

<thead class="sticky top-0 z-10 bg-base-200">
	<tr bind:this={tr}>
		{#each columns as column, index}
			<td
				class="
					float-left
					relative
					flex items-center
					border-r-2 border-neutral-500
					px-2
				"
				class:opacity-30={isMoving}
				style="
					width: calc(var(--browser-zoom) * {column.width} * 0.025px);
					height: calc(var(--browser-zoom) * 1px);
					font-size: calc(var(--browser-zoom) * 0.4px)
				"
				data-name="{column.name}"
				data-index="{index}"
				on:mousedown={() => {onColumnMousedown(event, index);}}
			>
				<div class="overflow-x-hidden whitespace-nowrap text-ellipsis">
					{#if column.name == "icon"}
						<div class="h-3/4 mx-auto ">
							<Icon src={Sparkle} />
						</div>
					{:else}
						{column.label}
						<div
							class="
								absolute
								top-0 right-0
								w-2 h-full
								cursor-col-resize
							"
							on:mousedown={() => {onResizerMousedown(event, index);}}
							aria-hidden
						></div>
					{/if}
				</div>
			</td>
		{/each}
	</tr>
</thead>

<style lang="postcss">
	td {
		&:global(.moving) {
			position: absolute;
			top: 0;
			cursor: move;
			@apply bg-base-300;
			width: 100px !important;
			border: 0;
		}
	}
</style>
