import type { PageServerLoad, Actions } from "./$types";
import { fail, redirect } from "@sveltejs/kit";
import { setError, superValidate } from "sveltekit-superforms/server";
import { formSchema } from "./schema";
import { BASE_API_URL } from "$env/static/private";

export const load: PageServerLoad = async ({ locals }) => {
	if (locals.user) {
		redirect(302, "/app");
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

		const res = await fetch(`${BASE_API_URL}/users/activate`, {
			method: "PUT",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify({
				token: form.data.token,
			}),
		});

		const response = await res.json();
		if (!res.ok) {
			for (const field in response.error) {
				return setError(form, field, response.error[field]);
			}
		}

		redirect(303, "/login");
	},
};
