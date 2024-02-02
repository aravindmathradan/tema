<script lang="ts">
	import CodeExplorer from "$lib/components/CodeExplorer.svelte";
	import Sidebar from "$lib/components/Sidebar.svelte";
	import Titlebar from "$lib/components/Titlebar.svelte";
</script>

<div class="flex border mx-auto w-full h-full">
	<Sidebar />
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
