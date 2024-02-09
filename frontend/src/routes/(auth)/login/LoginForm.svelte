<script lang="ts">
	import * as Form from "$lib/components/ui/form";
	import type { FormOptions } from "formsnap";
	import { formSchema, type FormSchema } from "./schema";
	import type { SuperValidated } from "sveltekit-superforms";
	import { toast } from "svelte-sonner";
	import { goto } from "$app/navigation";

	export let form: SuperValidated<FormSchema>;

	let formError: string;
	const options: FormOptions<FormSchema> = {
		applyAction: false,
		onResult({ result }) {
			if (result.type === "redirect") {
				goto(result.location);
			}
		},
		onError({ result }) {
			if (result.status === 401) formError = result.error.message;
			else toast.error(result.error.message);
		},
	};
</script>

{#if formError}
	<span class="text-destructive">{formError}</span>
{/if}
<Form.Root
	method="POST"
	{form}
	schema={formSchema}
	{options}
	let:config
	class="flex flex-col gap-3"
>
	<Form.Field {config} name="email">
		<Form.Item>
			<Form.Label class="text-base">Email</Form.Label>
			<Form.Input class="text-base" type="email" />
			<Form.Validation />
		</Form.Item>
	</Form.Field>
	<Form.Field {config} name="password">
		<Form.Item>
			<Form.Label class="text-base">Password</Form.Label>
			<Form.Input class="text-base" type="password" />
			<Form.Validation />
		</Form.Item>
	</Form.Field>
	<Form.Button class="mt-3 bg-primary p-5">
		<p class="text-lg text-primary-foreground">Login</p>
	</Form.Button>
</Form.Root>
