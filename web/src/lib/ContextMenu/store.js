import { writable } from "svelte/store";

class ContextMenu {
	constructor() {
		this.items = writable([]);
		this.visible = writable(false);
		this.left = writable(0);
		this.top = writable(0);
	}

	setItems(items) {
		this.items.set(items);
	}

	open(left, top) {
		this.left.set(left);
		this.top.set(top);
		this.visible.set(true);
	}

	close() {
		this.visible.set(false);
	}
}

export const contextMenu = new ContextMenu();
export const contextMenuItems = contextMenu.items;
export const contextMenuVisible = contextMenu.visible;
export const contextMenuLeft = contextMenu.left;
export const contextMenuTop = contextMenu.top;
