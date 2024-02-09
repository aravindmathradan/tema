import { z } from "zod";
import { timestampsSchema } from "./timestamps";

export const userSchema = z
	.object({
		id: z.number(),
		name: z.string(),
		email: z.string().email(),
		activated: z.boolean(),
	})
	.merge(timestampsSchema);

export type UserSchema = z.infer<typeof userSchema>;
