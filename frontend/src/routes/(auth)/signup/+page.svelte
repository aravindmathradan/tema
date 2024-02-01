<script lang="ts">
	import type { PageData } from "./$types";
	import SignupForm from "./SignupForm.svelte";
	export let data: PageData;
</script>

<div class="flex flex-col gap-[22px] items-center justify-center h-screen w-screen border">
	<SignupForm form={data.form} />
	<div>
		Already have an account?
		<a href="/login" class="text-accent">Login here</a>
	</div>
</div>
