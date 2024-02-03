import type { PageServerLoad, Actions } from "./$types";
import { fail, redirect, type NumericRange } from "@sveltejs/kit";
import { message, setError, superValidate } from "sveltekit-superforms/server";
import { formSchema } from "./schema";
import { BASE_API_URL } from "$env/static/private";

export const load: PageServerLoad = async (event) => {
	if (event.locals.user) {
		throw redirect(302, "/app");
	}
	return {
		form: await superValidate(formSchema),
	};
};

export const actions: Actions = {
	default: async (event) => {
		const form = await superValidate(event, formSchema);
		if (!form.valid) {
			return fail(400, {
				form,
			});
		}

		const res = await fetch(`${BASE_API_URL}/tokens/authentication`, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			credentials: "include",
			body: JSON.stringify({
				email: form.data.email,
				password: form.data.password,
			}),
		});

		const response = await res.json();
		if (!res.ok) {
			if (typeof response.error === "string") {
				return message(form, response.error, {
					status: <NumericRange<400, 599>>res.status,
				});
			}
			for (const field in response.error) {
				return setError(form, field, response.error[field]);
			}
		}

		event.cookies.set("auth-token", response.authentication_token.token, {
			path: "/",
			httpOnly: true,
			sameSite: "strict",
			secure: process.env.NODE_ENV === "production",
			maxAge: 60 * 60 * 24 * 7, // 1 week
		});

		throw redirect(303, "/app");
	},
};
