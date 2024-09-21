<script lang="ts">
	import Button from '$lib/Button.svelte';
	import type { NewSubmission, NewSubmissionRaw } from '$lib/types';
	import Label from '$lib/Label.svelte';

	export let onSubmit: (data: NewSubmission) => void;
	let form: NewSubmissionRaw = {
		name: '',
		class: '',
		mode: '',
		tier: '',
		duration: '',
		video: '',
		build: ''
	};
	export let isSubmitting: boolean;
	function validate(data: NewSubmissionRaw) {
		const errs: Partial<NewSubmissionRaw> = {};
		if (!data.name) {
			data.name = 'Required';
		}
		if (!data.class) {
			data.class = 'Required';
		}
		if (!data.mode) {
			data.mode = 'Required';
		}
		if (!data.tier) {
			data.tier = 'Required';
		}
		if (!data.duration) {
			data.duration = 'Required';
		}
		if (!data.video) {
			data.video = 'Required';
		}

		return errs;
	}
	function toSeconds(str: string) {
		return 900;
	}
	function parseData(raw: NewSubmissionRaw): NewSubmission {
		return { ...raw, tier: parseInt(raw.tier, 10), duration: toSeconds(raw.duration) };
	}
</script>

<form
	class="container"
	on:submit|preventDefault={(e) => {
		onSubmit(parseData(form));
	}}
>
	<div class="field">
		<Label for="name">Name*</Label>
		<input name="name" bind:value={form.name} required />
	</div>
	<div class="field">
		<Label for="class">Class*</Label>
		<select name="class" bind:value={form.class} required>
			<option value="barbarian">Barbarian</option>
			<option value="druid">Druid</option>
			<option value="necromancer">Necromancer</option>
			<option value="rogue">Rogue</option>
			<option value="sorcerer">Sorcerer</option>
		</select>
	</div>
	<div class="field">
		<Label for="mode">Mode*</Label>
		<select name="mode" bind:value={form.mode} required>
			<option value="softcore">Softcore</option>
			<option value="hardcore">Hardcore</option>
		</select>
	</div>
	<div class="field">
		<Label for="tier">Tier*</Label>
		<input name="tier" bind:value={form.tier} required placeholder="150" />
	</div>
	<div class="field">
		<Label for="time">Time*</Label>
		<input name="time" bind:value={form.duration} required placeholder="12:43" />
	</div>
	<div class="field">
		<Label for="video">Video*</Label>
		<input name="video" bind:value={form.video} required />
	</div>
	<div class="field">
		<Label for="build">Build</Label>
		<input name="build" bind:value={form.build} />
	</div>
	<Button type="submit" disabled={isSubmitting}>Submit</Button>
</form>

<style>
	.container {
		margin: 0 auto;
		max-width: var(--container-width);
	}

	.field {
		margin-bottom: 16px;
	}

	input {
		display: block;
		width: 100%;
		padding-top: 8px;
		padding-right: 16px;
		padding-bottom: 8px;
		padding-left: 16px;
		appearance: none;
		outline: none;
		font-size: 1rem;
		font-family: var(--font-default);
		line-height: 1.5;
		color: white;
		background-color: #171920;
		border: 1px solid rgba(255, 255, 255, 0.36);
		border-radius: 4px;
		transition:
			border 0.2s ease,
			background-color 0.2s ease;
		text-overflow: ellipsis;
		white-space: nowrap;
		box-shadow:
			rgba(0, 0, 0, 0.016) 0px 2px 2px 0px,
			rgba(0, 0, 0, 0.008) 0px 0px 0px 1px;
	}

	input:hover {
		border-color: #7abfff;
		background-color: #171920;
	}

	input::focus {
		border-color: #148eff;
		background-color: #171920;
		outline: none;
	}

	input:is(:-webkit-autofill, :autofill) {
		filter: none;
		-webkit-text-fill-color: white;
		-webkit-box-shadow: 0 0 0px 40rem #171920 inset;
		border-color: #148eff;
		background-color: #171920 !important;
	}

	input:-webkit-autofill,
	input:-webkit-autofill:hover,
	input:-webkit-autofill:focus {
		filter: none;
		-webkit-text-fill-color: white;
		-webkit-box-shadow: 0 0 0px 40rem #171920 inset;
		border-color: #148eff;
		background-color: #171920 !important;
	}

	select::placeholder,
	input::placeholder {
		color: var(--text-default);
	}

	select {
		display: block;
		width: 100%;
		padding-top: 8px;
		padding-right: 16px;
		padding-bottom: 8px;
		padding-left: 16px;
		appearance: none;
		outline: none;
		font-size: 1rem;
		font-family: var(--font-default);
		line-height: 1.5;
		color: white;
		background-color: #171920;
		border: 1px solid rgba(255, 255, 255, 0.36);
		border-radius: 4px;
		transition:
			border 0.2s ease,
			background-color 0.2s ease;
		text-overflow: ellipsis;
		white-space: nowrap;
		box-shadow:
			rgba(0, 0, 0, 0.016) 0px 2px 2px 0px,
			rgba(0, 0, 0, 0.008) 0px 0px 0px 1px;
	}

	select:hover {
		border-color: #7abfff;
	}

	select:focus {
		border-color: #148eff;
	}
</style>
