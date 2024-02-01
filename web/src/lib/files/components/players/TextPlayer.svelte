<script>
	export let file;

	let data = null;

	let getData = async (file) => {
		let response = await fetch(file.dataURL);
		data = await response.text();
	};

	$: {
		getData(file);
	}
</script>

<div class="h-full w-full overflow-auto p-4">
	<pre class="font-mono">{data}</pre>
</div>
