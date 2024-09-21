<script lang="ts">
	import Alert from '$lib/Alert.svelte';
	import type { APIError } from '$lib/types';

	export let error: APIError | null;
</script>

{#if error}
	<Alert title="Error" message="Something went wrong" />
{/if}
