<script lang="ts">
	import { Enter, OpenInNewWindow } from "radix-icons-svelte";
	import ProfilePopover from "./ProfilePopover.svelte";
	import ThemeSwitcher from "./ThemeSwitcher.svelte";

	export let profile: any;
</script>

<div class="flex items-center justify-between border w-full p-3">
	<div class="text-4xl font-bold text-primary px-3">Tema</div>

	<div class="flex items-center h-full">
		<ThemeSwitcher />
		{#if profile}
			<a
				href="/app"
				class="flex items-center justify-center gap-2 leading-normal text-lg font-medium bg-primary text-primary-foreground ml-4 px-4 py-2 rounded-md hover:bg-primary/90 active:translate-y-[1px] active:translate-x-[1px] transition-all duration-300 ease-out"
			>
				Open App
				<OpenInNewWindow class="w-5 h-5" />
			</a>
			<span class="border-l mr-3 ml-5 h-full"></span>
		{:else}
			<a
				href="/app"
				class="flex items-center justify-center gap-2 leading-normal text-lg font-medium bg-primary text-primary-foreground mx-4 px-4 py-2 rounded-md hover:bg-primary/90 active:translate-y-[1px] active:translate-x-[1px] transition-all duration-300 ease-out"
			>
				Login
				<Enter class="w-5 h-5" />
			</a>
		{/if}
		{#if profile}
			<ProfilePopover
				userName={profile?.name}
				userEmail={profile?.email}
				transcriptDropdownOpen={false}
				avatarOnly={true}
				class="md:p-2 hover:bg-secondary/55 rounded-md transition-colors duration-300 ease-out"
			/>
		{/if}
	</div>
</div>
