import { call } from "./base";

async function user() {
	let response = await call({
		path: "/auth/user",
	});
	let data = await response.json();
	return data;
}

async function login({username, password}) {
	return await call({
		path: "/auth/login",
		method: "POST",
		body: {
			"username": username,
			"password": password,
		},
	});
}

async function logout() {
	await call({
		path: "/auth/logout",
		method: "POST",
	});
}

export default {
	user,
	login,
	logout,
};
