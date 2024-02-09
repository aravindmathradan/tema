import { ServerErrorCodes, ServerErrorSubCodes } from "$lib/constants/error-codes";
import { z } from "zod";

const fieldErrorsSchema = z.object({
	subCode: z.nativeEnum(ServerErrorSubCodes),
	message: z.string(),
});

export const errorSchema = z.object({
	code: z.nativeEnum(ServerErrorCodes),
	message: z.string(),
	fields: z.record(fieldErrorsSchema),
});

export const errorResponseSchema = z.object({
	error: errorSchema,
});
