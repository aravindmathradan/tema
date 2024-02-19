import { env } from "$env/dynamic/private";
import { error } from "@sveltejs/kit";
import type { LayoutServerLoad } from "./$types";
import { createProjectFormSchema, projectListResponseSchema } from "$lib/types/projects";
import { superValidate } from "sveltekit-superforms/server";

export const load: LayoutServerLoad = async ({ locals, cookies, fetch, depends }) => {
	const createProjectForm = await superValidate(createProjectFormSchema);

	const token = cookies.get("auth-token");

	const res = await fetch(`${env.BASE_API_URL}/projects?sort=-updated_at`, {
		method: "GET",
		headers: {
			Authorization: `Bearer ${token}`,
		},
	});

	if (!res.ok) {
		error(500, {
			message: "Something went wrong. Please try again later",
		});
	}

	const { projects } = projectListResponseSchema.parse(await res.json());

	depends("app:projectsFetch");

	return {
		user: locals.user,
		projects,
		createProjectForm,
	};
};
