<script lang="ts">
	import { MagnifyingGlass } from "radix-icons-svelte";
</script>

<div
	class="rounded-md h-fit flex items-center px-2 bg-gray-200 focus-within:bg-gray-100 dark:bg-gray-800 dark:focus-within:bg-gray-700"
>
	<MagnifyingGlass class="w-5 h-5" />
	<input
		type="text"
		placeholder="Search"
		class="outline-none border-none bg-transparent p-2 w-full"
	/>
</div>
