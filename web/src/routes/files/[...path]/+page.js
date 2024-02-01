import { apiAuth } from "$lib/apiAuth";

export const load = async ({ params }) => {
	await apiAuth.loginRequired();

	return {
		path: params.path,
	};
};
