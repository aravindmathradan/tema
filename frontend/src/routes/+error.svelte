<script lang="ts">
	import { page } from "$app/stores";
	import { Home } from "radix-icons-svelte";
</script>

<div class="flex items-center justify-center w-full h-full">
	<div class="text-center">
		{#if $page.status === 404}
			<h1 class="text-5xl font-medium mb-10">You must be lost.</h1>
		{:else if $page.status === 500}
			<h1 class="text-3xl font-medium">Woops! We weren't expecting that.</h1>
			<p class="mt-4 mb-12 text-xl text-muted-foreground">We're investigating the issue.</p>
		{/if}
		<a
			href="/"
			class="font-medium text-xl text-accent underline underline-offset-4 hover:font-bold transition-all duration-300 ease-in-out"
		>
			Go Home
		</a>
	</div>
</div>
