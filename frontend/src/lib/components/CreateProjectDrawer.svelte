<script lang="ts">
	import * as Drawer from "$lib/components/ui/drawer";
	import { Button } from "$lib/components/ui/button";

	import * as Form from "$lib/components/ui/form";
	import type { SuperValidated } from "sveltekit-superforms";
	import type { FormOptions } from "formsnap";
	import { toast } from "svelte-sonner";
	import { createProjectFormSchema, type CreateProjectFormSchema } from "$lib/types/projects";

	export let form: SuperValidated<CreateProjectFormSchema>;

	let open: boolean;
	let createButtonLoading = false;

	const options: FormOptions<CreateProjectFormSchema> = {
		applyAction: false,
		onSubmit() {
			createButtonLoading = true;
		},
		onResult({ result }) {
			createButtonLoading = false;
			if (result.status >= 200 && result.status < 400) open = false;
		},
		onError({ result }) {
			createButtonLoading = false;
			toast.error(result.error.message);
		},
	};
</script>

<Drawer.Root bind:open>
	<Drawer.Trigger class="mt-8">
		<Button
			class="w-full bg-primary p-5 hover:bg-primary/90 shadow-sm hover:shadow-none shadow-primary/50 active:translate-y-[1px] active:translate-x-[1px] transition-all duration-300 ease-out"
		>
			<p class="text-lg font-medium text-primary-foreground">Create new project</p>
		</Button>
	</Drawer.Trigger>
	<Drawer.Content>
		<div class="mx-auto w-full max-w-sm">
			<Drawer.Header>
				<Drawer.Title>Create a new project</Drawer.Title>
				<Drawer.Description
					>You can organize your transcripts and codes in a project</Drawer.Description
				>
			</Drawer.Header>
			<Drawer.Footer>
				<Form.Root
					method="POST"
					{form}
					schema={createProjectFormSchema}
					{options}
					let:config
					class="flex flex-col gap-3"
					action="?/createProject"
				>
					<Form.Field {config} name="projectName">
						<Form.Item>
							<Form.Label class="text-base">Name</Form.Label>
							<Form.Input class="text-base" autocomplete="off" />
							<Form.Validation />
						</Form.Item>
					</Form.Field>
					<Form.Field {config} name="projectDescription">
						<Form.Item>
							<Form.Label class="text-base">Description</Form.Label>
							<Form.Textarea
								class="text-base resize-none"
								rows={5}
								placeholder="Describe your project (optional)"
							/>
							<Form.Validation />
						</Form.Item>
					</Form.Field>
					<Form.Button class="mt-3 bg-primary p-5S" disabled={createButtonLoading}>
						<p class="text-lg text-primary-foreground">Create</p>
						{#if createButtonLoading}
							<span class="icon-[mingcute--loading-line] ml-3 animate-spin"></span>
						{/if}
					</Form.Button>
				</Form.Root>
				<Drawer.Close class="hover:font-medium">Cancel</Drawer.Close>
			</Drawer.Footer>
		</div>
	</Drawer.Content>
</Drawer.Root>
