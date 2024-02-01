import { z } from "zod";

export const formSchema = z
	.object({
		name: z.string().min(1, "Required").max(100, "Name must contain at most 100 characters"),
		email: z.string().email(),
		password: z
			.string()
			.min(8, "password must contain at least 8 characters")
			.max(50, "password must contain at most 50 characters"),
		confirmPassword: z.string().min(1, "Required"),
	})
	.refine((data) => data.password === data.confirmPassword, {
		message: "Passwords don't match",
		path: ["confirmPassword"],
	});

export type FormSchema = typeof formSchema;
