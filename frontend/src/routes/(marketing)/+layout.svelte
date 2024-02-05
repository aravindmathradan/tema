<script lang="ts">
	import Navbar from "$lib/components/Navbar.svelte";

	export let data;

	$: profile = data.profile;
</script>

<div class="flex flex-col border mx-auto w-full h-full">
	<Navbar {profile} />
	<slot />
</div>
