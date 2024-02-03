<script lang="ts">
	import {
		CardStack,
		CaretDown,
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

	export let userName: string | undefined;
	export let userEmail: string | undefined;

	const profilePictureFallbackLetters = userName
		?.split(" ")
		.reduce((out, curr) => out + curr.charAt(0), "")
		.slice(0, 2)
		.toUpperCase();

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
	<div class="flex flex-col flex-1 mt-3 min-h-0 overflow-y-auto">
		<Dialog.Root>
			<Dialog.Trigger
				class="flex gap-[22px] mb-3 p-3 w-full items-center bg-primary hover:bg-primary/90 rounded-md shadow-sm hover:shadow-none shadow-primary/50 active:translate-y-[0.5px]  transition-colors duration-300 ease-out"
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
				{#each projects as { title, href }}
					<a
						{href}
						class="flex gap-[22px] p-3 w-full items-center hover:bg-secondary/55 border-b-2 transition-colors duration-300 ease-out"
					>
						<p class="text-lg capitalize ml-2">{title}</p>
					</a>
				{/each}
				<Button
					class="mt-8 bg-primary p-5 hover:bg-primary/90 shadow-sm hover:shadow-none shadow-primary/50 active:translate-y-[0.5px] transition-colors duration-300 ease-out"
				>
					<p class="text-lg font-medium text-primary-foreground">Create new project</p>
				</Button>
			</Dialog.Content>
		</Dialog.Root>
		{#each sidebarOptions as { icon, title, href, contents }}
			{#if contents.length > 0}
				<Collapsible.Root
					onOpenChange={onTranscriptDropdownOpenChange}
					bind:open={transcriptDropdownOpen}
					class="flex flex-col min-h-0 mb-3"
				>
					<Collapsible.Trigger
						class="flex gap-[22px] p-3 w-full items-center justify-between hover:bg-secondary/55 rounded-md transition-colors duration-300 ease-out"
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
					<Collapsible.Content class="overflow-y-auto border-y-[1px]">
						{#each contents as { title, href }}
							<a
								on:click={closeTranscriptDropdown}
								{href}
								class="flex gap-[22px] p-2 w-full items-center hover:text-accent transition-colors duration-300 ease-out"
							>
								<Reader class="ml-12 w-4 h-4" />
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

	<Popover.Root>
		<Popover.Trigger
			class="mb-3 md:p-2 hover:bg-secondary/55 rounded-md transition-colors duration-300 ease-out"
		>
			<div class="flex gap-2 items-center justify-center">
				<Avatar.Root>
					<!-- <Avatar.Image src="./favicon.png" alt="user" /> -->
					<Avatar.Fallback>{profilePictureFallbackLetters}</Avatar.Fallback>
				</Avatar.Root>
				{#if transcriptDropdownOpen}
					<div class="text-start md:inline-block hover:bg-white/19 rounded-md">
						<p class="font-medium capitalize">{userName}</p>
						<p class="text-gray-500">{userEmail}</p>
					</div>
				{:else}
					<div class="text-start hidden md:inline-block hover:bg-white/19 rounded-md">
						<p class="font-medium capitalize">{userName}</p>
						<p class="text-gray-500">{userEmail}</p>
					</div>
				{/if}
			</div>
		</Popover.Trigger>
		<Popover.Content class="flex flex-col w-40 md:w-[280px] p-2">
			<a
				href="/"
				class="flex gap-[22px] p-3 w-full items-center hover:text-accent transition-colors duration-300 ease-out"
			>
				<Person class="w-5 h-5" />
				<p>Account</p>
			</a>
			<span class="w-36 md:w-60 self-center border-b-2"></span>
			<form action="/logout" method="POST">
				<button
					type="submit"
					class="flex gap-[22px] p-3 w-full items-center hover:text-accent transition-colors duration-300 ease-out"
				>
					<Exit class="w-5 h-5" />
					<p>Logout</p>
				</button>
			</form>
		</Popover.Content>
	</Popover.Root>
</div>
