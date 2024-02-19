import { z } from "zod";

export const metadataSchema = z
	.object({
		currentPage: z.number().optional(),
		pageSize: z.number().optional(),
		firstPage: z.number().optional(),
		lastPage: z.number().optional(),
		totalRecords: z.number().optional(),
	})
	.default({});

export const metadataResponseSchema = z.object({
	metadata: z.optional(metadataSchema),
});
