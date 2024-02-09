import { z } from "zod";

export const formSchema = z.object({
	email: z.string().email(),
	password: z.string().min(1, "Required"),
});

export type FormSchema = typeof formSchema;

export const responseSchema = z.object({
	authenticationToken: z.object({
		token: z.string(),
		expiry: z.coerce.date(),
	}),
	refreshToken: z.object({
		token: z.string(),
		expiry: z.coerce.date(),
	}),
});

export type ResponseSchema = z.infer<typeof responseSchema>;
