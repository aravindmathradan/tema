<script lang="ts">
	import type { PageData } from "./$types";
	import ActivateForm from "./ActivateForm.svelte";
	export let data: PageData;
</script>

<div class="text-lg text-primary">
	{#if data.user}
		<p class="text-lg text-primary font-medium">Hi {data.user.name}</p>
	{/if}
	<p>We have sent an activation code to your email. Paste it below and click on activate.</p>
</div>
<ActivateForm form={data.form} />
