import { z } from "zod";

export const formSchema = z.object({
	token: z.string().min(1, "Required"),
});

export type FormSchema = typeof formSchema;
