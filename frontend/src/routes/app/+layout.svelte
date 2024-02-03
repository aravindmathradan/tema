<script lang="ts">
	import CodeExplorer from "$lib/components/CodeExplorer.svelte";
	import Sidebar from "$lib/components/Sidebar.svelte";
	import Titlebar from "$lib/components/Titlebar.svelte";
	import type { LayoutData } from "./$types";

	export let data: LayoutData;
</script>

<div class="flex border mx-auto w-full h-full">
	<Sidebar userName={data.user?.name} userEmail={data.user?.email} />
	<div class="flex flex-1 flex-col">
		<Titlebar />
		<div class="flex flex-1">
			<main class="flex-1">
				<slot />
			</main>
			<CodeExplorer />
		</div>
	</div>
</div>
