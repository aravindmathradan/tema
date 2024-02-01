<script lang="ts">
	import type { PageData } from "./$types";
	import LoginForm from "./LoginForm.svelte";
	export let data: PageData;
</script>

<div class="flex flex-col gap-[22px] items-center justify-center h-screen w-screen border">
	<LoginForm form={data.form} />
	<div>
		Don't have an account?
		<a href="/signup" class="text-accent">Sign up</a>
	</div>
</div>
