<script lang="ts">
	import Heading from '$lib/Heading.svelte';
	import Text from '$lib/Text.svelte';
	import SubmissionForm from '$lib/SubmissionForm.svelte';
	import type { NewSubmission } from '$lib/types';

	const handleSubmit = async (data: NewSubmission) => {
		try {
			const res = await fetch('/api/v1/submissions', {
				method: 'POST',
				body: JSON.stringify(data)
			});
			const json = await res.json();
		} catch {}
	};
</script>

<svelte:head>
	<title>Diablo 4 Pit - Submission</title>
</svelte:head>
<Heading>Submission</Heading>
<Text>
	Submit a video as proof of a successful pit run. The video will be verified and added to the
	leaderboard.
</Text>
<SubmissionForm onSubmit={handleSubmit} />
