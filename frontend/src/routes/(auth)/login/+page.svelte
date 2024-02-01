<script lang="ts">
	import type { PageData } from "./$types";
	import LoginForm from "./LoginForm.svelte";
	export let data: PageData;
</script>

<div class="mb-5 text-5xl text-left font-light">
	Login to <span class="font-bold">Tema</span>
</div>
<LoginForm form={data.form} />
<div>
	Don't have an account?
	<a href="/signup" class="text-accent font-medium">Sign up</a>
</div>
