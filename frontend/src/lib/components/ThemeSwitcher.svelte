<script lang="ts">
	import { Switch } from "$lib/components/ui/switch";
	import { Moon, Sun } from "radix-icons-svelte";

	let darkMode = false;
	function themeToggled(checked: boolean) {
		darkMode = !checked;

		darkMode
			? document.documentElement.classList.add("dark")
			: document.documentElement.classList.remove("dark");
	}
</script>

<div class="flex gap-2 items-center">
	<label for="theme-switch" class="mr-1">
		{#if darkMode}
			<Moon class="w-5 h-5" />
		{:else}
			<Sun class="w-5 h-5" />
		{/if}
	</label>
	<Switch checked={!darkMode} id="theme-switch" onCheckedChange={themeToggled} />
</div>
