import type { PageServerLoad, Actions } from "./$types";
import { fail, redirect, error } from "@sveltejs/kit";
import { superValidate } from "sveltekit-superforms/server";
import { formSchema, responseSchema } from "./schema";
import { env } from "$env/dynamic/private";
import { ServerErrorCodes } from "$lib/constants/error-codes";
import { DateTime, Interval } from "luxon";
import { errorResponseSchema } from "$lib/types/errors";

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

		const res = await event.fetch(`${env.BASE_API_URL}/tokens/authentication`, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			credentials: "include",
			body: JSON.stringify({
				email: form.data.email,
				password: form.data.password,
				scope: "authentication",
			}),
		});

		if (!res.ok) {
			const response = errorResponseSchema.parse(await res.json());
			if (res.status === 401 && response.error.code === ServerErrorCodes.EINVALIDCREDENTIALS) {
				error(401, {
					message: "Email/Password do not match our records",
				});
			}
			error(500, {
				message: "Something went wrong. Please try again later",
			});
		}

		const response = responseSchema.parse(await res.json());
		let authInterval = Interval.fromDateTimes(
			DateTime.now(),
			DateTime.fromJSDate(response.authenticationToken.expiry),
		);
		let refreshInterval = Interval.fromDateTimes(
			DateTime.now(),
			DateTime.fromJSDate(response.refreshToken.expiry),
		);

		event.cookies.set("auth-token", response.authenticationToken.token, {
			path: "/",
			httpOnly: true,
			sameSite: "strict",
			secure: process.env.NODE_ENV === "production",
			maxAge: authInterval.length("seconds"),
		});

		event.cookies.set("refresh-token", response.refreshToken.token, {
			path: "/",
			httpOnly: true,
			sameSite: "strict",
			secure: process.env.NODE_ENV === "production",
			maxAge: refreshInterval.length("seconds"),
		});

		redirect(303, "/app");
	},
};
