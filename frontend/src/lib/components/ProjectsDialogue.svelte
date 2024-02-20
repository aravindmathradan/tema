<script lang="ts">
	import * as Dialog from "$lib/components/ui/dialog";
	import { Tokens } from "radix-icons-svelte";
	import CreateProjectDrawer from "./CreateProjectDrawer.svelte";
	import type { SuperValidated } from "sveltekit-superforms";
	import type {
		CreateProjectFormSchema,
		ProjectListSchema,
		ProjectSchema,
	} from "$lib/types/projects";
	import { currentProject } from "$lib/stores/project";
	import { toast } from "svelte-sonner";
	import { onDestroy } from "svelte";

	export let projects: ProjectListSchema;
	export let transcriptDropdownOpen: any;
	export let createProjectForm: SuperValidated<CreateProjectFormSchema>;

	let dialogOpen = false;

	if ($currentProject.id == 0) {
		if (projects && projects.length > 0) {
			$currentProject = projects[0];
		} else {
			toast.info("Create a project to get started");
			dialogOpen = true;
		}
	}

	const switchProject = (project: ProjectSchema) => {
		$currentProject = project;
		dialogOpen = false;
	};

	onDestroy(() => {
		$currentProject = {
			id: 0,
			name: "Project",
		};
	});
</script>

<Dialog.Root bind:open={dialogOpen}>
	<Dialog.Trigger
		class="flex gap-[22px] mb-3 p-3 w-full items-center bg-primary hover:bg-primary/90 rounded-md shadow-sm hover:shadow-none shadow-primary/50 active:translate-y-[1px] active:translate-x-[1px] transition-all duration-300 ease-out"
	>
		{#if transcriptDropdownOpen}
			<p class="md:inline-block text-lg font-medium text-primary-foreground">
				{$currentProject.name}
			</p>
		{:else}
			<p class="hidden md:inline-block text-lg font-medium text-primary-foreground">
				{$currentProject.name}
			</p>
		{/if}
		<Tokens class="ml-auto w-7 h-7 md:w-6 md:h-6 text-primary-foreground" />
	</Dialog.Trigger>
	<Dialog.Content class="gap-0">
		<div class="text-xl font-medium mb-6">Select the project</div>
		<div class="overflow-y-auto max-h-96">
			{#if projects && projects.length > 0}
				{#each projects as project}
					<button
						on:click={() => switchProject(project)}
						class="flex gap-[22px] p-3 w-full items-center hover:bg-secondary/55 border-b transition-colors duration-300 ease-out"
					>
						<p class="text-lg capitalize ml-2">{project.name}</p>
					</button>
				{/each}
			{:else}
				<div class="text-center mt-5">You don't have any projects. Create one to get started!</div>
			{/if}
		</div>
		<CreateProjectDrawer form={createProjectForm} />
	</Dialog.Content>
</Dialog.Root>
