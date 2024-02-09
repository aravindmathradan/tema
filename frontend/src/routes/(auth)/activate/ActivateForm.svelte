<script lang="ts">
	import * as Form from "$lib/components/ui/form";
	import type { FormOptions } from "formsnap";
	import { formSchema, type FormSchema } from "./schema";
	import type { SuperValidated } from "sveltekit-superforms";
	import { toast } from "svelte-sonner";
	import { goto } from "$app/navigation";

	export let form: SuperValidated<FormSchema>;
	const options: FormOptions<FormSchema> = {
		applyAction: false,
		onResult({ result }) {
			if (result.status >= 200 && result.status < 400) {
				toast.success("Great! Your account has been activated");
			}
			if (result.type === "redirect") {
				goto(result.location);
			}
		},
		onError({ result }) {
			if (result.status && result.status >= 400) toast.error("Could not activate your account!");
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
