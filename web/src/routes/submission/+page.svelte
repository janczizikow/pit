<script lang="ts">
	import Heading from '$lib/Heading.svelte';
	import Text from '$lib/Text.svelte';
	import SubmissionForm from '$lib/SubmissionForm.svelte';
	import { writable } from 'svelte/store';
	import ErrorMessage from '$lib/ErrorMessage.svelte';
	import Modal from '$lib/Modal.svelte';
	import { createQuery } from '@tanstack/svelte-query';
	import {
		type Season,
		type NewSubmission,
		type APIError,
		SeasonApi,
		SubmissionApi
	} from '$lib/api';
	let seasonId = 6;
	let showModal = false;
	let error = writable<APIError | null>(null);

	const query = createQuery<Season>({
		queryKey: ['currentSeason'],
		queryFn: () => new SeasonApi().getCurrentSeason()
	});
	$: if ($query?.isSuccess) {
		seasonId = $query.data.id;
	}

	const handleSubmit = async (data: NewSubmission) => {
		try {
			error.set(null);
			await new SubmissionApi().createSubmission({
				id: seasonId,
				newSubmission: data
			});
			showModal = true;
			return true;
		} catch (err) {
			error.set(err as APIError);
			return false;
		}
	};
</script>

<svelte:head>
	<title>Diablo 4 Pit - Submission</title>
	<meta
		name="description"
		content="unofficial Diablo 4 Seasonal Pit Leaderboard. Submit your pit run and see how you rank!"
	/>
</svelte:head>
<div class="container">
	<Heading>Submission</Heading>
	<Text>
		Submit a video as proof of a successful pit run. The video will be verified and added to the
		leaderboard.
	</Text>
	<ErrorMessage error={$error} />
	<SubmissionForm {seasonId} isBusy={$query.isLoading} onSubmit={handleSubmit} />
	<Modal bind:showModal>
		<h2 slot="header">Submission successful</h2>
		<p class="text">
			Thank you for your submission! Once it's verified it will show in the leaderboard.
		</p>
	</Modal>
</div>

<style>
	.container {
		padding: 0 16px;
		margin: 0 auto;
		width: 100%;
		max-width: var(--container-width);
	}

	h2 {
		font-family: 'Poppins', Helvetica, Arial, sans-serif;
	}
	.text {
		font-size: 16px;
		text-shadow: 3px 5px 5px rgba(0, 0, 0, 0.5);
		color: var(--text-default);
	}
</style>
