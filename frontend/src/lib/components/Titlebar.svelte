<script lang="ts">
	import { Switch } from "$lib/components/ui/switch";
	import { Moon, Sun } from "radix-icons-svelte";

	let darkMode = true;
	function themeToggled(checked: boolean) {
		darkMode = checked;

		darkMode
			? document.documentElement.classList.add("dark")
			: document.documentElement.classList.remove("dark");
	}
</script>

<div class="flex items-center justify-between border border-l-0 p-3 w-full">
	<div class="text-lg">Transcript Name</div>
	<div class="flex gap-2 items-center">
		<label for="theme-switch">
			{#if darkMode}
				<Moon class="w-5 h-5" />
			{:else}
				<Sun class="w-5 h-5" />
			{/if}
		</label>
		<Switch checked={darkMode} id="theme-switch" class="ml-1" onCheckedChange={themeToggled} />
	</div>
</div>
