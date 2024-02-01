<script>
	export let columns;

	$: tableWidth = Object.values(columns).reduce((a, b) => {
		let sum = 0;
		if (a) {
			sum += a.width;
		}
		if (b) {
			sum += b.width;
		}
		return {width: sum};
	})["width"];
</script>

<table
	class="table-fixed select-none"
	style="width: round(up, calc(var(--browser-zoom) * {tableWidth} * 0.025px), 1px)"
>
	<slot />
</table>
