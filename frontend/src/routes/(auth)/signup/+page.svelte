<script lang="ts">
	import type { PageData } from "./$types";
	import SignupForm from "./SignupForm.svelte";
	export let data: PageData;
</script>

<div class="mb-5 text-7xl lg:text-8xl text-left font-light flex-1">
	Join <span class="md:block md:text-8xl lg:text-9xl font-bold">Tema</span>
</div>
<div class="flex flex-col flex-1 gap-5">
	<SignupForm form={data.form} />
	<div class="text-right">
		Have an account?
		<a href="/login" class="text-accent font-medium">Login</a>
	</div>
</div>
