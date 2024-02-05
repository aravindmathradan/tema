<script lang="ts">
	import { Frame } from "radix-icons-svelte";
	import CodeSearchbar from "./CodeSearchbar.svelte";
</script>

<div
	class="border border-t-0 hidden md:flex md:flex-[0.5] px-3 md:px-2 py-2 space-y-4 flex-col sticky z-10"
>
	<div class="flex gap-[22px] h-fit p-2">
		<Frame class="w-7 h-7" />
		<p class="text-xl font-medium">Code Explorer</p>
	</div>
	<CodeSearchbar />
</div>
