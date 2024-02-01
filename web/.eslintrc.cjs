module.exports = {
	root: true,
	env: {
		"browser": true,
		"es2021": true,
		"node": true,
	},
	extends: [
		"eslint:recommended",
		"plugin:svelte/recommended",
	],
	parserOptions: {
		"sourceType": "module",
		"ecmaVersion": 2020,
		"extraFileExtensions": [".svelte"],
	},
	rules: {
		"indent": ["error", "tab"],
		"linebreak-style": ["error", "unix"],
		"quotes": ["error", "double"],
		"semi": ["error", "always"],
	},
};
