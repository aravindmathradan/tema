import { superValidate } from "sveltekit-superforms/server";
import type { Actions } from "./$types";
import { createProjectFormSchema } from "$lib/types/projects";
import { error, fail } from "@sveltejs/kit";
import { env } from "$env/dynamic/private";

export const actions: Actions = {
	createProject: async (event) => {
		const form = await superValidate(event, createProjectFormSchema);
		if (!form.valid) {
			return fail(400, {
				form,
			});
		}

		const token = event.cookies.get("auth-token");

		const res = await event.fetch(`${env.BASE_API_URL}/projects`, {
			method: "POST",
			headers: {
				Authorization: `Bearer ${token}`,
			},
			body: JSON.stringify({
				name: form.data.projectName,
				description: form.data.projectDescription,
			}),
		});

		if (!res.ok) {
			error(500, {
				message: "Something went wrong. Please try again later",
			});
		}

		return {
			form,
		};
	},
};
