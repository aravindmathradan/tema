<script lang="ts">
	import { toast } from "svelte-sonner";
	import * as Form from "$lib/components/ui/form";
	import type { FormOptions } from "formsnap";
	import { formSchema, type FormSchema } from "./schema";
	import type { SuperValidated } from "sveltekit-superforms";
	import { goto } from "$app/navigation";

	export let form: SuperValidated<FormSchema>;

	let formError: string;
	const options: FormOptions<FormSchema> = {
		applyAction: false,
		onSubmit() {
			// toast.info("Submitting...");
		},
		onResult({ result }) {
			if (result.status >= 200 && result.status < 400) {
				toast.info("Thank you for joining!");
			}
			if (result.type === "redirect") {
				goto(result.location);
			}
		},
		onError({ result }) {
			toast.error(result.error.message);
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
	<Form.Field {config} name="name">
		<Form.Item>
			<Form.Label class="text-base">Name</Form.Label>
			<Form.Input class="text-base" />
			<Form.Validation />
		</Form.Item>
	</Form.Field>
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
	<Form.Field {config} name="confirmPassword">
		<Form.Item>
			<Form.Label class="text-base">Confirm password</Form.Label>
			<Form.Input class="text-base" type="password" />
			<Form.Validation />
		</Form.Item>
	</Form.Field>
	<Form.Button class="mt-3 bg-primary p-5">
		<p class="text-lg text-primary-foreground">Signup</p>
	</Form.Button>
</Form.Root>
