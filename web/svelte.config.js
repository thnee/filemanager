import path from "path";
import { fileURLToPath } from "url";

import adapter from "@sveltejs/adapter-static";
import preprocess from "svelte-preprocess";

const __dirname = path.dirname(fileURLToPath(import.meta.url));

export default {
	preprocess: [
		preprocess({
			postcss: {
				configFilePath: path.resolve(__dirname, "postcss.config.js"),
			}
		}),
	],

	kit: {
		adapter: adapter({
			pages: "build",
			assets: "build",
			fallback: "index.html",
			precompress: true,
			strict: true,
		}),
	},

	vitePlugin: {
		inspector: true,
	},
};
