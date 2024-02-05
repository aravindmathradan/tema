import { redirect } from "@sveltejs/kit";
import type { Actions, PageServerLoad } from "./$types";
import { BASE_API_URL } from "$env/static/private";

export const load: PageServerLoad = async ({}) => {
	redirect(302, "/");
};

export const actions: Actions = {
	default: async ({ cookies, locals }) => {
		const token = cookies.get("auth-token");

		await fetch(`${BASE_API_URL}/tokens/authentication`, {
			method: "DELETE",
			headers: {
				Authorization: `Bearer ${token}`,
			},
		});

		cookies.delete("auth-token", { path: "/" });
		cookies.delete("refresh-token", { path: "/" });
		locals.user = null;

		redirect(302, "/");
	},
};
