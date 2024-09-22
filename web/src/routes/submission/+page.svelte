<script lang="ts">
	import Heading from '$lib/Heading.svelte';
	import Text from '$lib/Text.svelte';
	import SubmissionForm from '$lib/SubmissionForm.svelte';
	import type { APIError, NewSubmission } from '$lib/types';
	import { writable } from 'svelte/store';
	import ErrorMessage from '$lib/ErrorMessage.svelte';

	let error = writable<APIError | null>(null);

	const handleSubmit = async (data: NewSubmission) => {
		try {
			error.set(null);
			const res = await fetch('/api/v1/seasons/5/submissions', {
				method: 'POST',
				body: JSON.stringify(data)
			});
			const json = await res.json();
			if (res.status >= 300) {
				throw json;
			}
			return true;
		} catch (err) {
			error.set(err as APIError);
			return false;
		}
	};
</script>

<svelte:head>
	<title>Diablo 4 Pit - Submission</title>
</svelte:head>
<div class="container">
	<Heading>Submission</Heading>
	<Text>
		Submit a video as proof of a successful pit run. The video will be verified and added to the
		leaderboard.
	</Text>
	<ErrorMessage error={$error} />
	<SubmissionForm onSubmit={handleSubmit} />
</div>

<style>
	.container {
		padding: 0 16px;
		margin: 0 auto;
		width: 100%;
		max-width: var(--container-width);
	}
</style>
