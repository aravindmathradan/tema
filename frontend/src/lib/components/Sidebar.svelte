<script lang="ts">
	import {
		CaretSort,
		CaretUp,
		Exit,
		Frame,
		Home,
		MagnifyingGlass,
		Person,
		Reader,
		Tokens,
	} from "radix-icons-svelte";
	import * as Avatar from "$lib/components/ui/avatar";
	import * as Collapsible from "$lib/components/ui/collapsible";
	import * as Popover from "$lib/components/ui/popover";
	import * as Dialog from "$lib/components/ui/dialog";
	import Button from "./ui/button/button.svelte";

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
			icon: Reader,
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
					title: "Transcript 1",
					href: "/",
				},
				{
					title: "Transcript 2",
					href: "/",
				},
				{
					title: "Transcript 1",
					href: "/",
				},
			],
		},
	];

	const projects = [
		{
			title: "Project 1",
			href: "/",
		},
		{
			title: "Project 2",
			href: "/",
		},
		{
			title: "Project 3",
			href: "/",
		},
	];

	let transcriptDropdownOpen: boolean = false;
	function onTranscriptDropdownOpenChange(state: boolean) {
		transcriptDropdownOpen = state;
	}
	function closeTranscriptDropdown() {
		if (window.innerWidth < 768) transcriptDropdownOpen = false;
	}
</script>

<div
	class="px-3 md:flex-[0.15] md:px-2 border flex justify-between flex-col sticky z-10 top-0 left-0"
>
	<div class="mt-3">
		<Dialog.Root>
			<Dialog.Trigger
				class="flex gap-[22px] mb-3 p-3 w-full items-center bg-primary rounded-md shadow-sm hover:bg-primary/75"
			>
				{#if transcriptDropdownOpen}
					<p class="md:inline-block text-lg font-medium text-primary-foreground">Project 1</p>
				{:else}
					<p class="hidden md:inline-block text-lg font-medium text-primary-foreground">
						Project 1
					</p>
				{/if}
				<Tokens class="ml-auto w-7 h-7 md:w-6 md:h-6 text-primary-foreground" />
			</Dialog.Trigger>
			<Dialog.Content class="gap-0">
				<div class="text-xl font-medium mb-6">Select the project</div>
				{#each projects as { title, href }, index}
					<a
						{href}
						class="flex gap-[22px] p-3 w-full items-center hover:bg-black/5 dark:hover:bg-white/10 border-b-2"
					>
						<p class="text-lg capitalize ml-2">{title}</p>
					</a>
				{/each}
				<Button class="mt-8 bg-accent p-5 hover:bg-accent/75">
					<p class="text-lg font-medium text-accent-foreground">Create new project</p>
				</Button>
			</Dialog.Content>
		</Dialog.Root>
		{#each sidebarOptions as { icon, title, href, contents }}
			{#if contents.length > 0}
				<Collapsible.Root
					onOpenChange={onTranscriptDropdownOpenChange}
					bind:open={transcriptDropdownOpen}
				>
					<Collapsible.Trigger
						class="flex gap-[22px] p-3 w-full items-center hover:bg-black/5 dark:hover:bg-white/10 rounded-md"
					>
						<svelte:component this={icon} class="w-7 h-7" />
						{#if transcriptDropdownOpen}
							<p class="text-lg capitalize md:inline-block pr-5">{title}</p>
							<CaretUp class="hidden md:inline-block w-6 h-6 ml-auto" />
						{:else}
							<p class="text-lg capitalize hidden md:inline-block pr-5">{title}</p>
							<CaretSort class="hidden md:inline-block w-6 h-6 ml-auto" />
						{/if}
					</Collapsible.Trigger>
					<Collapsible.Content class="overflow-y-auto max-h-80">
						{#each contents as { title, href }}
							<a
								on:click={closeTranscriptDropdown}
								{href}
								class="flex gap-[22px] ml-12 p-2 w-fit items-center hover:bg-black/5 dark:hover:bg-white/10 rounded-md"
							>
								<p class="text-base capitalize">{title}</p>
							</a>
						{/each}
					</Collapsible.Content>
				</Collapsible.Root>
			{:else}
				<a
					{href}
					class="flex gap-[22px] p-3 w-full items-center hover:bg-black/5 dark:hover:bg-white/10 rounded-md"
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

	<Popover.Root>
		<Popover.Trigger class="mb-3 md:p-2 hover:bg-black/5 dark:hover:bg-white/10 rounded-md">
			<div class="flex gap-2 items-center justify-center">
				<Avatar.Root>
					<Avatar.Image src="./favicon.png" alt="user" />
					<Avatar.Fallback>US</Avatar.Fallback>
				</Avatar.Root>
				{#if transcriptDropdownOpen}
					<div class="text-start md:inline-block hover:bg-white/19 rounded-md">
						<p class="font-medium capitalize">Aravind Unnikrishnan</p>
						<p class="text-gray-500">aravindmathradan@gmail.com</p>
					</div>
				{:else}
					<div class="text-start hidden md:inline-block hover:bg-white/19 rounded-md">
						<p class="font-medium capitalize">Aravind Unnikrishnan</p>
						<p class="text-gray-500">aravindmathradan@gmail.com</p>
					</div>
				{/if}
			</div>
		</Popover.Trigger>
		<Popover.Content class="flex flex-col w-40 md:w-[280px] p-2">
			<a href="/" class="flex gap-[22px] p-3 w-full items-center">
				<Person class="w-5 h-5" />
				<p>Account</p>
			</a>
			<span class="w-36 md:w-60 self-center border-b-2"></span>
			<a href="/" class="flex gap-[22px] p-3 w-full items-center">
				<Exit class="w-5 h-5" />
				<p>Logout</p>
			</a>
		</Popover.Content>
	</Popover.Root>
</div>
