import type { PageServerLoad, Actions } from "./$types";
import { fail, redirect, error } from "@sveltejs/kit";
import { setError, superValidate } from "sveltekit-superforms/server";
import { formSchema, responseSchema } from "./schema";
import { env } from "$env/dynamic/private";
import { ServerErrorCodes, ServerErrorSubCodes } from "$lib/constants/error-codes";
import { errorResponseSchema } from "$lib/types/errors";
import { toast } from "svelte-sonner";

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

		const res = await event.fetch(`${env.BASE_API_URL}/users`, {
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

		if (!res.ok) {
			const response = errorResponseSchema.parse(await res.json());
			if (res.status === 422 && response.error.code === ServerErrorCodes.EFAILEDVALIDATION) {
				let { fields } = response.error;
				for (let key in fields) {
					if (key == "email" && fields[key].subCode === ServerErrorSubCodes.EEMAILALREADYEXISTS) {
						return setError(form, "email", "An account is already created with this email address");
					}
				}
				error(422, {
					message: "Something went wrong. Please check your inputs and try again",
				});
			}
			error(500, {
				message: "Something went wrong. Please try again later",
			});
		}

		redirect(303, "/activate");
	},
};
