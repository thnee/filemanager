<script>
	import FileTableTable from "./FileTableTable.svelte";
	import FileTableHead from "./FileTableHead.svelte";
	import FileTableBody from "./FileTableBody.svelte";

	import { columns } from "../stores/fileTable";
</script>

<FileTableTable bind:columns={$columns}>
	<FileTableHead bind:columns={$columns} />
	<FileTableBody bind:columns={$columns} />
</FileTableTable>
