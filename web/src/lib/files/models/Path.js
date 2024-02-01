export default class Path {
	constructor(path) {
		this.path = path.split("/");
		// Remove empty string elements after split
		this.path = this.path.filter((item) => {return item;});
	}

	get depth() {
		return this.path.length;
	}

	get area() {
		return this.path[0];
	}

	get parentRendered() {
		return this.path.slice(0, -1).join("/");
	}

	get rendered() {
		return this.path.join("/");
	}

	equals(path) {
		return JSON.stringify(this.path) == JSON.stringify(path.path);
	}
}
