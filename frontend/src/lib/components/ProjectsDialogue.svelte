<script lang="ts">
	import * as Dialog from "$lib/components/ui/dialog";
	import { Tokens } from "radix-icons-svelte";
	import CreateProjectDrawer from "./CreateProjectDrawer.svelte";
	import type { SuperValidated } from "sveltekit-superforms";
	import type { CreateProjectFormSchema } from "$lib/types/projects";

	export let projects: any;
	export let transcriptDropdownOpen: any;
	export let createProjectForm: SuperValidated<CreateProjectFormSchema>;
</script>

<Dialog.Root>
	<Dialog.Trigger
		class="flex gap-[22px] mb-3 p-3 w-full items-center bg-primary hover:bg-primary/90 rounded-md shadow-sm hover:shadow-none shadow-primary/50 active:translate-y-[1px] active:translate-x-[1px] transition-all duration-300 ease-out"
	>
		{#if transcriptDropdownOpen}
			<p class="md:inline-block text-lg font-medium text-primary-foreground">Project 1</p>
		{:else}
			<p class="hidden md:inline-block text-lg font-medium text-primary-foreground">Project 1</p>
		{/if}
		<Tokens class="ml-auto w-7 h-7 md:w-6 md:h-6 text-primary-foreground" />
	</Dialog.Trigger>
	<Dialog.Content class="gap-0">
		<div class="text-xl font-medium mb-6">Select the project</div>
		<div class="overflow-y-auto max-h-96">
			{#if projects && projects.length > 0}
				{#each projects as { name, id }}
					<button
						class="flex gap-[22px] p-3 w-full items-center hover:bg-secondary/55 border-b transition-colors duration-300 ease-out"
					>
						<p class="text-lg capitalize ml-2">{name}</p>
					</button>
				{/each}
			{:else}
				<div class="text-center mt-5">You don't have any projects. Create one to get started!</div>
			{/if}
		</div>
		<CreateProjectDrawer form={createProjectForm} />
	</Dialog.Content>
</Dialog.Root>
