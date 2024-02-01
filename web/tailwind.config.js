import themes from "./themes.js";

export default {
	content: ["./src/**/*.{html,js,css,scss,svelte}"],
	theme: {
		extend: {},
	},
	plugins: [
		require("daisyui"),
	],
	daisyui: {
		themes: themes,
	},
};
