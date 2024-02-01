<script lang="ts">
	import { MagnifyingGlass } from "radix-icons-svelte";
	import { Input } from "./ui/input";
</script>

<div class="rounded-md h-fit flex items-center px-2 bg-secondary focus-within:bg-secondary/80">
	<MagnifyingGlass class="w-5 h-5" />
	<Input
		type="text"
		placeholder="Search"
		class="outline-none border-none focus-visible:ring-0 p-2"
	/>
</div>
