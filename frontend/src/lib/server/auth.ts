import { BASE_API_URL } from "$env/static/private";
import { type RequestEvent } from "@sveltejs/kit";

async function fetchCurrentUser(authToken: string) {
	const res = await fetch(`${BASE_API_URL}/users/me`, {
		method: "GET",
		headers: {
			Authorization: `Bearer ${authToken}`,
		},
	});

	const response: any = await res.json();
	if (res.ok) {
		return {
			user: response?.user,
			status: res.status,
			error: null,
		};
	}

	return {
		user: null,
		status: res.status,
		error: response.error,
	};
}

async function refreshAuthentication(refreshToken: string) {
	const res = await fetch(`${BASE_API_URL}/tokens/authentication`, {
		method: "POST",
		body: JSON.stringify({
			scope: "refresh",
			refresh_token: refreshToken,
		}),
	});

	const response = await res.json();
	if (res.ok) {
		return {
			authToken: response.authentication_token.token,
			status: res.status,
			error: null,
		};
	}

	return {
		authToken: null,
		status: res.status,
		error: response.error,
	};
}

export const getAuthenticatedUser = async (event: RequestEvent) => {
	const authToken = event.cookies.get("auth-token");

	if (!authToken) return null;

	let userResponse = await fetchCurrentUser(authToken);
	if (userResponse.user) {
		return userResponse.user;
	}

	if (userResponse.status == 401) {
		const refreshToken = event.cookies.get("refresh-token");
		if (!refreshToken) return null;

		let refreshResponse = await refreshAuthentication(refreshToken);
		if (refreshResponse.authToken) {
			event.cookies.set("auth-token", refreshResponse.authToken, {
				path: "/",
				httpOnly: true,
				sameSite: "strict",
				secure: process.env.NODE_ENV === "production",
				maxAge: 60 * 60 * 24 * 7, // 1 week
			});
			userResponse = await fetchCurrentUser(refreshResponse.authToken);
			if (userResponse.user) {
				return userResponse.user;
			}
		}
	}

	event.cookies.delete("refresh-token", { path: "/" });
	event.cookies.delete("auth-token", { path: "/" });
	return null;
};
