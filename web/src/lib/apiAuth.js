import { writable, get } from "svelte/store";
import { redirect } from "@sveltejs/kit";
import { goto } from "$app/navigation";

import api from "$lib/api";

export class User {
	constructor(data) {
		if (data !== null) {
			this.username = data.Username;
		} else {
			this.username = null;
		}
	}

	get isLoggedIn() {
		return this.username != null;
	}
}

export const apiUser = writable(new User(null));

export const apiAuth = {
	async updateUser() {
		let r = await api.auth.user();
		let data = await r.json();
		apiUser.set(new User(data));
	},

	async login(params) {
		let r = await api.auth.login(params);
		await this.updateUser();
		return r;
	},

	async logout() {
		await api.auth.logout();
		await goto("/login");
		apiUser.set(new User(null));
	},

	async loginRequired() {
		if (!get(apiUser).isLoggedIn) {
			await this.updateUser();
		}

		if (!get(apiUser).isLoggedIn) {
			redirect(302, "/login");
		}
	},
};
