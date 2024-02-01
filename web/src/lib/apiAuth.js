import { writable, get } from "svelte/store";
import { redirect } from "@sveltejs/kit";
import { goto } from "$app/navigation";

import api from "$lib/api";

export class User {
	constructor(data) {
		if (data !== null) {
			this.id = data.ID;
			this.username = data.Username;
			this.email = data.Email;
		} else {
			this.id = null;
		}
	}

	get isLoggedIn() {
		return this.id != null;
	}
}

export const apiUser = writable(new User(null));

export const apiAuth = {
	async updateUser() {
		let data = await api.auth.user();
		apiUser.set(new User(data));
	},

	async login(params) {
		await api.auth.login(params);
		await this.updateUser();
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
