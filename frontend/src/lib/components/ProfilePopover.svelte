<script lang="ts">
	import * as Popover from "$lib/components/ui/popover";
	import * as Avatar from "$lib/components/ui/avatar";
	import { Exit, Person } from "radix-icons-svelte";

	export let userName: string | undefined;
	export let userEmail: string | undefined;
	export let transcriptDropdownOpen: boolean;
	export let avatarOnly: boolean;

	const profilePictureFallbackLetters = userName
		?.split(" ")
		.reduce((out, curr) => out + curr.charAt(0), "")
		.slice(0, 2)
		.toUpperCase();
</script>

<Popover.Root>
	<Popover.Trigger class={$$restProps.class || ""}>
		<div class="flex gap-2 items-center justify-center">
			<Avatar.Root>
				<!-- <Avatar.Image src="./favicon.png" alt="user" /> -->
				<Avatar.Fallback class="bg-accent"
					><p class="text-lg text-accent-foreground">
						{profilePictureFallbackLetters}
					</p></Avatar.Fallback
				>
			</Avatar.Root>
			{#if transcriptDropdownOpen}
				<div
					class={avatarOnly
						? "text-start hidden hover:bg-white/19 rounded-md"
						: "text-start md:inline-block hover:bg-white/19 rounded-md"}
				>
					<p class="font-medium capitalize">{userName}</p>
					<p class="text-gray-500">{userEmail}</p>
				</div>
			{:else}
				<div
					class={avatarOnly
						? "text-start hidden hover:bg-white/19 rounded-md"
						: "text-start md:inline-block hover:bg-white/19 rounded-md"}
				>
					<p class="font-medium capitalize">{userName}</p>
					<p class="text-gray-500">{userEmail}</p>
				</div>
			{/if}
		</div>
	</Popover.Trigger>
	<Popover.Content class="flex flex-col w-fit p-2">
		<a
			href="/"
			class="flex gap-[22px] p-3 w-full items-center justify-between hover:text-accent transition-colors duration-300 ease-out"
		>
			<Person class="w-5 h-5" />
			<p>Account</p>
		</a>
		<span class="w-11/12 self-center border-b-2"></span>
		<form action="/logout" method="POST">
			<button
				type="submit"
				class="flex gap-[22px] p-3 w-full items-center justify-between hover:text-accent transition-colors duration-300 ease-out"
			>
				<Exit class="w-5 h-5" />
				<p>Logout</p>
			</button>
		</form>
	</Popover.Content>
</Popover.Root>
