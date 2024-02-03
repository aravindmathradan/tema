import type { LayoutServerLoad } from "./$types";

export const load: LayoutServerLoad = async (event) => {
	let user = event.locals.user;

	return {
		user,
	};
};
