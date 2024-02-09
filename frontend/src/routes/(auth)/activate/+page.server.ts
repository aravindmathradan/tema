import type { PageServerLoad, Actions } from "./$types";
import { error, fail, redirect } from "@sveltejs/kit";
import { setError, superValidate } from "sveltekit-superforms/server";
import { formSchema } from "./schema";
import { env } from "$env/dynamic/private";
import { errorResponseSchema } from "$lib/types/errors";
import { ServerErrorCodes } from "$lib/constants/error-codes";

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

		const res = await event.fetch(`${env.BASE_API_URL}/users/activate`, {
			method: "PUT",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify({
				token: form.data.token,
			}),
		});

		if (!res.ok) {
			const response = errorResponseSchema.parse(await res.json());
			if (res.status === 422 && response.error.code === ServerErrorCodes.EFAILEDVALIDATION) {
				return setError(form, "token", "The token is either invalid or expired");
			}
			error(500, {
				message: "Something went wrong. Please try again later by re-generating token",
			});
		}

		redirect(303, "/login");
	},
};
