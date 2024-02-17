<script lang="ts">
	import Navbar from "$lib/components/Navbar.svelte";

	export let data;
</script>

<div class="flex flex-col border mx-auto w-full h-full">
	<Navbar user={data.user} />
	<slot />
</div>
