import { z } from "zod";

export const timestampsSchema = z.object({
	createdAt: z.coerce.date().optional(),
	updatedAt: z.coerce.date().optional(),
});

export type TimestampsSchema = z.infer<typeof timestampsSchema>;
