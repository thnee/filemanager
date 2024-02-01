import { apiAuth } from "$lib/apiAuth";

export const load = async () => {
	await apiAuth.loginRequired();
};
