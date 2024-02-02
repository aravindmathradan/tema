import type { RequestEvent } from "@sveltejs/kit";

export const getAuthenticatedUser = (event: RequestEvent) => {
	const { cookies } = event;

	const token = cookies.get("auth-token");

	if (!token) {
		event.locals.user = null;
		return null;
	}

	// check the token against db and return user if exists

	const user = {
		id: 1,
		email: "aravindmathradan@gmail.com",
		name: "aravind",
	};
	return user;
};
