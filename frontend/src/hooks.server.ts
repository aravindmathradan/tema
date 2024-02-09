import { getAuthenticatedUser } from "$lib/server/auth";
import { redirect, type Handle, type HandleFetch, error } from "@sveltejs/kit";
import camelcaseKeys from "camelcase-keys";

export const handle: Handle = async ({ event, resolve }) => {
	event.locals.user = await getAuthenticatedUser(event);

	if (!event.locals.user) {
		if (event.url.pathname.startsWith("/app")) {
			redirect(302, "/login");
		}
	}

	const response = await resolve(event);

	return response;
};

export const handleFetch: HandleFetch = async ({ request, fetch }) => {
	try {
		const res = await fetch(request);
		const response = await res.json();

		let camelcaseResponse = camelcaseKeys(response, { deep: true });

		return new Response(JSON.stringify(camelcaseResponse), {
			headers: res.headers,
			status: res.status,
			statusText: res.statusText,
		});
	} catch (err: any) {
		if (err.cause?.code === "ECONNREFUSED" || err.cause?.code === "ECONNTIMEOUT") {
			error(500, {
				message: "Could not connect to the server. Please try again later",
			});
		} else {
			error(500, {
				message: "Something went wrong. Please try again later",
			});
		}
	}
};
