<script lang="ts">
	import * as Form from "$lib/components/ui/form";
	import { formSchema, type FormSchema } from "./schema";
	import type { SuperValidated } from "sveltekit-superforms";

	export let form: SuperValidated<FormSchema>;
</script>

<Form.Root method="POST" {form} schema={formSchema} let:config class="flex flex-col gap-3">
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
