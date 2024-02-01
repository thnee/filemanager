import { writable } from "svelte/store";

export let columns = writable([
	{
		name: "icon",
		label: "",
		width: 42,
	},
	{
		name: "name",
		label: "Name",
		width: 500,
	},
	{
		name: "size",
		label: "Size",
		width: 100,
	},
	{
		name: "mtime",
		label: "Modified",
		width: 165,
	},
	{
		name: "ctime",
		label: "Created",
		width: 165,
	},
]);
