import { z } from "zod";

export const formSchema = z.object({
	token: z.string().length(26, "Token must be 26 characters long"),
});

export type FormSchema = typeof formSchema;
