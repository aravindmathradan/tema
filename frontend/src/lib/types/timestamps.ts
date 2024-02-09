import { z } from "zod";

export const timestampsSchema = z.object({
	createdAt: z.coerce.date(),
	updatedAt: z.coerce.date(),
});

export type TimestampsSchema = z.infer<typeof timestampsSchema>;
