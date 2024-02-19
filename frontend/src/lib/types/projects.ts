import { z } from "zod";
import { timestampsSchema } from "./timestamps";
import { metadataResponseSchema } from "./metadata";

export const projectSchema = z
	.object({
		id: z.number().optional(),
		name: z.string().optional(),
		description: z.string().optional(),
		ownerId: z.number().optional(),
		archived: z.boolean().optional(),
	})
	.merge(timestampsSchema);

export const projectListSchema = z.array(projectSchema.default({}));

export const projectListResponseSchema = z
	.object({
		projects: projectListSchema,
	})
	.merge(metadataResponseSchema);

export const createProjectFormSchema = z.object({
	projectName: z.string().min(1, "Required").max(40, "Name must contain at most 100 characters"),
	projectDescription: z.string().max(500, "Description must contain at most 500 characters"),
});

export type ProjectSchema = z.infer<typeof projectSchema>;
export type ProjectListSchema = z.infer<typeof projectListSchema>;
export type ProjectListResponseSchema = z.infer<typeof projectListResponseSchema>;
export type CreateProjectFormSchema = typeof createProjectFormSchema;
