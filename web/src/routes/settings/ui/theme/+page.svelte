<script>
	import ThemePreview from "./ThemePreview.svelte";

	const themes = {
		"Dark": [
			"dark",
			"dim",
			"business",
			"dracula",
			"black",
			"forest",
			"halloween",
			"night",
			"sunset",
			"luxury",
			"coffee",
		],
		"Light": [
			"light",
			"cupcake",
			"bumblebee",
			"corporate",
			"cmyk",
			"wireframe",
			"lofi",
			"nord",
			"emerald",
			"garden",
			"winter",
			"pastel",
			"autumn",
			"fantasy",
			"acid",
			"lemonade",
		],
		"Other": [
			"synthwave",
			"cyberpunk",
			"retro",
			"valentine",
			"aqua",
		],
	};
</script>

<h2 class="text-xl mb-4">Select theme</h2>

<ThemePreview themes={themes} />
