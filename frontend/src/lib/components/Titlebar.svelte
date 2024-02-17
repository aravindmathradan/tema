<script lang="ts">
	import ProfilePopover from "./ProfilePopover.svelte";
	import ThemeSwitcher from "./ThemeSwitcher.svelte";

	export let user: any;
</script>

<div class="flex items-center justify-between border border-l-0 w-full h-18 p-2 md:p-0">
	<div class="text-lg ml-4">Transcript Name</div>
	<div class="flex">
		<ThemeSwitcher />
		<span class="md:border-l"></span>
		<ProfilePopover
			userName={user?.name}
			userEmail={user?.email}
			transcriptDropdownOpen={false}
			avatarOnly={false}
			class="md:p-2 md:px-3 hover:bg-secondary/55 transition-colors duration-300 ease-out"
		/>
	</div>
</div>
