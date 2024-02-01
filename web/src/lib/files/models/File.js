import dayjs from "dayjs";

import { fileBrowser } from "../stores/fileBrowser";

import Path from "./Path";

export default class File {
	constructor(data) {
		this.selected = false;
		this.imageData = null;

		this.name = data.Name;
		this.nameStem = data.NameStem;
		this.nameExt = data.NameExt;
		this.mimeType = data.MimeType;
		this.path = new Path(data.Path);
		this.isDir = data.IsDir;
		this.dataURL = data.DataURL;
		this.downloadURL = data.DownloadURL;
		this.size = data.Size;
		this.mode = data.Mode;
		this.atime = dayjs(data.Atime);
		this.mtime = dayjs(data.Mtime);
		this.ctime = dayjs(data.Ctime);
	}

	get nameSlashed() {
		let value = this.Name;
		if (this.isDir) {
			value += "/";
		}
		return value;
	}

	get category() {
		if (this.isDir) {
			return "directory";
		}
		var categories = [
			{category: "image", extensions: ["jpg", "jpeg", "png", "gif"]},
			{category: "video", extensions: ["mkv", "mp4", "mpg", "mpeg", "avi"]},
			{category: "text",  extensions: ["txt"]},
			{category: "code",  extensions: ["py", "rb", "js", "css", "html", "go", "md", "yml", "yaml", "json", "toml", "editorconfig", "gitignore", "gitignore_global", "gitmodules", "gitconfig", "yarnrc", "npmrc", "iyarc"]},
			{category: "shell", extensions: ["sh", "bash", "zsh", "fish"]},
		];
		let category = "unknown";
		categories.forEach((item) => {
			if (item.extensions.includes(this.nameExt.toLowerCase())) {
				category = item.category;
			}
		});
		return category;
	}

	get sizeHumanReadable() {
		let bytes = this.size;
		let si = false;
		let dp = 1;
		const thresh = si ? 1000 : 1024;

		if (Math.abs(bytes) < thresh) {
			return bytes + " B";
		}

		const units = si
			? ["kB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"]
			: ["KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB", "YiB"];
		let u = -1;
		const r = 10 ** dp;

		do {
			bytes /= thresh;
			++u;
		} while (Math.round(Math.abs(bytes) * r) / r >= thresh && u < units.length - 1);

		return bytes.toFixed(dp) + " " + units[u];
	}

	goto() {
		fileBrowser.gotoFile(this);
	}
}
