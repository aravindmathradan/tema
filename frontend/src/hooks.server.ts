import { getAuthenticatedUser } from "$lib/server/auth";
import { redirect, type Handle } from "@sveltejs/kit";

export const handle: Handle = async ({ event, resolve }) => {
	event.locals.user = await getAuthenticatedUser(event);

	if (!event.locals.user) {
		if (event.url.pathname.startsWith("/app")) {
			throw redirect(302, "/login");
		}
	}

	const response = await resolve(event);

	return response;
};
