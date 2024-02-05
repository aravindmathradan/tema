<script lang="ts">
	import * as Form from "$lib/components/ui/form";
	import type { FormOptions } from "formsnap";
	import { formSchema, type FormSchema } from "./schema";
	import type { SuperValidated } from "sveltekit-superforms";
	import { toast } from "svelte-sonner";

	export let form: SuperValidated<FormSchema>;
	let formError: string;
	const options: FormOptions<FormSchema> = {
		onSubmit() {
			// toast.info("Submitting...");
		},
		onResult({ result }) {
			formError = result.data?.form.message;
			if (result.status === 200) toast.success("Account activated!");
			if (result.status >= 400) toast.error("Could not activate your account!");
		},
	};
</script>

<Form.Root
	method="POST"
	{form}
	schema={formSchema}
	{options}
	let:config
	class="flex flex-col gap-3"
>
	<Form.Field {config} name="token">
		<Form.Item>
			<Form.Label class="text-base">Activation Code</Form.Label>
			<Form.Input class="text-base" />
			<Form.Validation />
		</Form.Item>
	</Form.Field>
	<Form.Button class="mt-3 bg-primary p-5">
		<p class="text-lg text-primary-foreground">Activate</p>
	</Form.Button>
</Form.Root>
