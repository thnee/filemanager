<script>
	export let file;
</script>

<img
	src={file.dataURL}
	alt={file.name}
	class="max-h-full max-w-full object-contain self-center mx-auto p-4"
/>
