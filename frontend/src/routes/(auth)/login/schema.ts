import { z } from "zod";

export const formSchema = z.object({
	email: z.string().min(1, "Required").email(),
	password: z.string().min(1, "Required"),
});

export type FormSchema = typeof formSchema;
