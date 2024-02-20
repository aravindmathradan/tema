import { type ProjectSchema } from "$lib/types/projects";
import { writable } from "svelte/store";

export const currentProject = writable<ProjectSchema>({
	id: 0,
	name: "Project",
});
