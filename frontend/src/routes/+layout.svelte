<script lang="ts">
	import { Toaster } from "svelte-sonner";
	import "../app.css";
	import { theme } from "$lib/stores/theme";
</script>

<svelte:head>
	<script>
		localStorage.theme === "dark"
			? document.documentElement.classList.add("dark")
			: document.documentElement.classList.remove("dark");
	</script>
</svelte:head>

<Toaster theme={$theme} richColors />
<slot />
