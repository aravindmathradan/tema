<script lang="ts">
	import CodeExplorer from "$lib/components/CodeExplorer.svelte";
	import Sidebar from "$lib/components/Sidebar.svelte";
	import Titlebar from "$lib/components/Titlebar.svelte";

	export let data;

	$: profile = data.profile;
</script>

<div class="flex border mx-auto w-full h-full">
	<Sidebar />
	<div class="flex flex-1 flex-col">
		<Titlebar {profile} />
		<div class="flex flex-1">
			<main class="flex-1">
				<slot />
			</main>
			<CodeExplorer />
		</div>
	</div>
</div>
