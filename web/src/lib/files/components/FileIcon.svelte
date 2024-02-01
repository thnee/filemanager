<script>
	import { Icon } from "@steeze-ui/svelte-icon";
	import { Folder, File, FileText, FileCode, FileTerminal, Image, Video } from "@steeze-ui/lucide-icons";

	export let file;

	let icons = {
		directory: Folder,
		image: Image,
		video: Video,
		code: FileCode,
		shell: FileTerminal,
		text: FileText,
		unknown: File,
	};

	let strokeWidth, fill, stroke, icon;

	$: {
		file;
		strokeWidth = "1";
		fill = "transparent";
		stroke = "oklch(var(--bc))";
		if (file.category === "directory") {
			strokeWidth = 0;
			fill = "oklch(var(--bc))";
		}
		icon = icons[file.category];
	}
</script>

<Icon
	src={icon}
	stroke-width={strokeWidth}
	fill={fill}
	stroke={stroke}
	class={"file-icon " + $$restProps.class || ""}
/>
