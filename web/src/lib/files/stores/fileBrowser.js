import { writable, derived, get } from "svelte/store";
import { goto } from "$app/navigation";

import File from "../models/File";
import Area from "../models/Area";
import Path from "../models/Path";

class FileBrowser {
	constructor() {
		this.path = null;

		this.areas = writable([]);

		this.files = writable([]);
		this.file = writable(null);

		this.zoom = writable(42);

		this.state = writable({event: null});

		this.upPath = this.createUpPath();

		this.prevFile = this.createPrevFile();
		this.nextFile = this.createNextFile();

		this.baseURL = "http://127.0.0.1:4000/api";
	}

	async init(path) {
		this.path = new Path(path);

		await this.loadAreas();
		this.setSelectedArea();

		// If no area is open, goto first area
		if (this.path.depth == 0) {
			this.gotoPath(get(this.areas)[0].name);
			return;
		}

		let file = await this.openFile();

		if (file) {
			if (file.isDir) {
				this.loadFiles(file.path.rendered);
			} else {
				this.loadFiles(file.path.parentRendered);
			}
		}
	}

	setSelectedArea() {
		let areas = get(this.areas);
		areas.forEach((area) => {
			if (area.name == this.path.area) {
				area.selected = true;
			} else {
				area.selected = false;
			}
		});
		this.areas.set(areas);
	}

	async loadAreas() {
		let response = await fetch(`${this.baseURL}/files/areas`);
		let data = await response.json();
		let areas = data.areas.map((areaData) => {
			return new Area(areaData);
		});
		this.areas.set(areas);
	}

	async openFile() {
		if (this.path.depth > 0) {
			let response = await fetch(
				`${this.baseURL}/files/file?` +
					new URLSearchParams({
						path: this.path.rendered,
					}),
			);
			let data = await response.json();

			let file = new File(data);
			this.file.set(file);
			this.state.set({event: "file-updated"});
			return file;
		}
	}

	async loadFiles(path) {
		let response = await fetch(`${this.baseURL}/files/list?` + new URLSearchParams({
			path: path,
		}));
		let data = await response.json();

		let files = [];
		files = data.Files.map((fileData) => {
			return new File(fileData);
		});

		this.files.set(files);
		this.state.set({event: "files-updated"});
	}

	selectFile(file) {
		file.selected = true;
		this.files.set(this.files);
	}

	selectOneFile(file) {
		this.files.update((files) => {
			files.map((_file) => {
				if (file === _file) {
					_file.selected = true;
				} else {
					_file.selected = false;
				}
				return _file;
			});
			return files;
		});
	}

	selectAllFiles() {
		this.files.update((files) => {
			files.map((_file) => {
				_file.selected = true;
			});
			return files;
		});
	}

	unselectAllFiles() {
		this.files.update((files) => {
			files.map((_file) => {
				_file.selected = false;
			});
			return files;
		});
	}

	getBrowserPath(path) {
		return `/files/${path}`;
	}

	createUpPath() {
		return derived(this.file, ($file) => {
			if ($file.path.depth > 1) {
				return this.getBrowserPath($file.path.parentRendered);
			} else {
				return null;
			}
		});
	}

	createPrevFile() {
		return derived([this.files, this.file], ([$files, $file]) => {
			let result = null;
			$files.forEach((_file, index) => {
				if (_file.path.equals($file.path)) {
					if (index > 0 && !$files[index - 1].isDir) {
						result = this.getBrowserPath($files[index - 1].path.rendered);
					}
				}
			});
			return result;
		});
	}

	createNextFile() {
		return derived([this.files, this.file], ([$files, $file]) => {
			let result = null;
			$files.forEach((_file, index) => {
				if (_file.path.equals($file.path)) {
					if (index < $files.length - 1 && !$files[index + 1].isDir) {
						result = this.getBrowserPath($files[index + 1].path.rendered);
					}
				}
			});
			return result;
		});
	}

	gotoPath(path) {
		goto(this.getBrowserPath(path));
	}

	gotoFile(file) {
		goto(this.getBrowserPath(file.path.rendered));
	}
}

export const fileBrowser = new FileBrowser();
export const path = fileBrowser.path;
export const areas = fileBrowser.areas;
export const files = fileBrowser.files;
export const file = fileBrowser.file;
export const zoom = fileBrowser.zoom;
export const state = fileBrowser.state;
export const upPath = fileBrowser.upPath;
export const prevFile = fileBrowser.prevFile;
export const nextFile = fileBrowser.nextFile;
