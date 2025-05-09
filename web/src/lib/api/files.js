import { call } from "./base";

async function getFileAreas() {
	return await call({
		path: "/files/areas",
		method: "GET",
	});
}

async function getFilesList(path) {
	return await call({
		path: "/files/list?" + new URLSearchParams({
			path: path,
		}),
	});
}

async function getFile(path) {
	return await call({
		path: `/files/file?` + new URLSearchParams({
			path: path,
		}),
	});
}

export default {
	getFileAreas,
	getFilesList,
	getFile,
};
