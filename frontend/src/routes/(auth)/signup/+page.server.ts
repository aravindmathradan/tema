import type { PageServerLoad, Actions } from "./$types";
import { fail, redirect, type NumericRange } from "@sveltejs/kit";
import { message, setError, superValidate } from "sveltekit-superforms/server";
import { formSchema } from "./schema";
import { BASE_API_URL } from "$env/static/private";

export const load: PageServerLoad = async () => {
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

		const res = await fetch(`${BASE_API_URL}/users`, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			credentials: "include",
			body: JSON.stringify({
				name: form.data.name,
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

		event.locals.signedupUser = response.user;

		throw redirect(303, "/activate");
	},
};
