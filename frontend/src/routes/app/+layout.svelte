<script lang="ts">
	import { afterNavigate, invalidate } from "$app/navigation";
	import CodeExplorer from "$lib/components/CodeExplorer.svelte";
	import Sidebar from "$lib/components/Sidebar.svelte";
	import Titlebar from "$lib/components/Titlebar.svelte";
	import type { LayoutServerData } from "./$types";

	export let data: LayoutServerData;

	afterNavigate(async () => {
		await invalidate("app:projectsFetch");
	});
</script>

<div class="flex border mx-auto w-full h-full">
	<Sidebar projects={data.projects} createProjectForm={data.createProjectForm} />
	<div class="flex flex-1 flex-col">
		<Titlebar user={data.user} />
		<div class="flex flex-1">
			<main class="flex-1">
				<slot />
			</main>
			<CodeExplorer />
		</div>
	</div>
</div>
