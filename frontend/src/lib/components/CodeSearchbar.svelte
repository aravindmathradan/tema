<script lang="ts">
	import { MagnifyingGlass } from "radix-icons-svelte";
	import { Input } from "./ui/input";
</script>

<div class="rounded-md h-fit flex items-center px-2 bg-secondary/85 focus-within:bg-secondary/75">
	<MagnifyingGlass class="w-5 h-5" />
	<Input
		type="text"
		placeholder="Search"
		class="text-md font-normal outline-none border-none focus-visible:ring-0 p-2"
	/>
</div>
