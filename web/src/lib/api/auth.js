import { call } from "./base";

async function user() {
	return await call({
		path: "/auth/user",
	});
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
