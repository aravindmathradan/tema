import { BASE_API_URL } from "$env/static/private";
import type { RequestEvent } from "@sveltejs/kit";

export const getAuthenticatedUser = async (event: RequestEvent) => {
	const { cookies } = event;

	const token = cookies.get("auth-token");

	if (!token) {
		event.locals.user = null;
		return null;
	}

	// check the token against db and return user if exists

	const res = await fetch(`${BASE_API_URL}/users/me`, {
		method: "GET",
		headers: {
			Authorization: `Bearer ${token}`,
		},
	});

	if (res.ok) {
		const response: any = await res.json();
		return response?.user;
	}

	event.locals.user = null;
	cookies.delete("auth-token", { path: "/" });
	return null;
};
