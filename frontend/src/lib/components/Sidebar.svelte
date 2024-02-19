<script lang="ts">
	import {
		CardStack,
		CaretDown,
		CaretUp,
		Frame,
		Home,
		MagnifyingGlass,
		Reader,
	} from "radix-icons-svelte";
	import * as Collapsible from "$lib/components/ui/collapsible";
	import ProjectsDialogue from "./ProjectsDialogue.svelte";
	import type { CreateProjectFormSchema, ProjectListSchema } from "$lib/types/projects";
	import type { SuperValidated } from "sveltekit-superforms";

	const sidebarOptions = [
		{
			icon: Home,
			title: "home",
			href: "/",
			contents: [],
		},
		{
			icon: MagnifyingGlass,
			title: "search",
			href: "/",
			contents: [],
		},
		{
			icon: Frame,
			title: "codes",
			href: "/",
			contents: [],
		},
		{
			icon: CardStack,
			title: "transcripts",
			href: "/",
			contents: [
				{
					title: "Transcript 1",
					href: "/",
				},
				{
					title: "Transcript 2",
					href: "/",
				},
				{
					title: "Transcript 3",
					href: "/",
				},
				{
					title: "Transcript 4",
					href: "/",
				},
				{
					title: "Transcript 5",
					href: "/",
				},
				{
					title: "Transcript 6",
					href: "/",
				},
				{
					title: "Transcript 7",
					href: "/",
				},
				{
					title: "Transcript 8",
					href: "/",
				},
				{
					title: "Transcript 9",
					href: "/",
				},
				{
					title: "Transcript 10",
					href: "/",
				},
				{
					title: "Transcript 11",
					href: "/",
				},
				{
					title: "Transcript 12",
					href: "/",
				},
			],
		},
	];

	export let projects: ProjectListSchema;
	export let createProjectForm: SuperValidated<CreateProjectFormSchema>;

	let transcriptDropdownOpen: boolean = false;
	function onTranscriptDropdownOpenChange(state: boolean) {
		transcriptDropdownOpen = state;
	}
	function closeTranscriptDropdown() {
		if (window.innerWidth < 768) transcriptDropdownOpen = false;
	}
</script>

<div
	class="px-3 md:flex-[0.22] md:px-2 border flex justify-between flex-col sticky z-10 top-0 left-0"
>
	<div class="flex flex-col flex-1 mt-3 min-h-0 overflow-y-auto">
		<ProjectsDialogue {projects} {transcriptDropdownOpen} {createProjectForm} />
		{#each sidebarOptions as { icon, title, href, contents }}
			{#if contents.length > 0}
				<Collapsible.Root
					onOpenChange={onTranscriptDropdownOpenChange}
					bind:open={transcriptDropdownOpen}
					class="flex flex-col min-h-0 mb-3"
				>
					<Collapsible.Trigger
						class="flex gap-[22px] p-3 w-full items-center hover:bg-secondary/55 rounded-md transition-colors duration-300 ease-out"
					>
						<svelte:component this={icon} class="w-7 h-7" />
						{#if transcriptDropdownOpen}
							<p class="text-lg capitalize md:inline-block">{title}</p>
							<CaretUp class="hidden md:inline-block w-6 h-6 ml-auto" />
						{:else}
							<p class="text-lg capitalize hidden md:inline-block">{title}</p>
							<CaretDown class="hidden md:inline-block w-6 h-6 ml-auto" />
						{/if}
					</Collapsible.Trigger>
					<Collapsible.Content class="overflow-y-auto bg-secondary/10 border-y rounded-md">
						{#each contents as { title, href }}
							<a
								on:click={closeTranscriptDropdown}
								{href}
								class="flex gap-[22px] p-2 w-full items-center hover:text-accent hover:font-medium hover:border-y transition-colors duration-300 ease-out"
							>
								<Reader class="ml-3 md:ml-8 w-4 h-4" />
								<p class="text-base capitalize">{title}</p>
							</a>
						{/each}
					</Collapsible.Content>
				</Collapsible.Root>
			{:else}
				<a
					{href}
					class="flex gap-[22px] p-3 w-full items-center hover:bg-secondary/55 rounded-md transition-colors duration-300 ease-out"
				>
					<svelte:component this={icon} class="w-7 h-7" />
					{#if transcriptDropdownOpen}
						<p class="text-lg capitalize md:inline-block pr-5">{title}</p>
					{:else}
						<p class="text-lg capitalize hidden md:inline-block pr-5">{title}</p>
					{/if}
				</a>
			{/if}
		{/each}
	</div>
</div>
