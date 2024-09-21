<script lang="ts">
	import Button from '$lib/Button.svelte';
	import type { NewSubmission } from '$lib/types';
	import Label from '$lib/Label.svelte';
	import HelperText from '$lib/HelperText.svelte';
	import { validator } from '@felte/validator-zod';
	import * as zod from 'zod';
	import { createForm } from 'felte';

	export let onSubmit: (data: NewSubmission) => Promise<boolean>;
	const MODES = ['softcore', 'hardcore'] as const;
	const CLASSES = ['barbarian', 'druid', 'necromancer', 'rogue', 'sorcerer'] as const;
	const schema = zod.object({
		name: zod.string().min(1),
		class: zod.enum(CLASSES),
		mode: zod.enum(MODES),
		tier: zod.coerce.number().gt(0).lte(200),
		duration: zod
			.string()
			.min(1)
			// TODO: fix this regex, right now it allows entering time greater than 15:00
			.regex(/^(?:[0-1][0-5]:[0-5][0-9]|(?:0[1-9]|1[0-4]):[0-5][0-9]|00:[0-5][0-9])$/, {
				message: 'Invalid time, must be in format minutes:seconds, eg.: 12:43 or 03:01'
			}),
		video: zod.string().url().min(1),
		build: zod.string().optional()
	});
	const { form, isSubmitting, reset, errors, isValid } = createForm<zod.infer<typeof schema>>({
		onSubmit: async (values) => {
			const shouldReset = await onSubmit(parseData(values));
			if (shouldReset) {
				reset();
				alert('Submission successful');
			}
		},
		extend: validator({ schema })
	});
	function toSeconds(str: string) {
		const [m, s] = str.split(':');
		return parseInt(m) * 60 + parseInt(s);
	}
	function parseData(raw: zod.infer<typeof schema>): NewSubmission {
		return { ...raw, duration: toSeconds(raw.duration) };
	}
</script>

<form class="container" use:form>
	<div class="field">
		<Label for="name">Name*</Label>
		<input name="name" required class={$errors.name ? 'input-error' : ''} />
		<HelperText>{$errors.name?.[0] || ''}</HelperText>
	</div>
	<div class="field">
		<Label for="class">Class*</Label>
		<select name="class" required class={$errors.class ? 'input-error' : ''}>
			<option value="barbarian">Barbarian</option>
			<option value="druid">Druid</option>
			<option value="necromancer">Necromancer</option>
			<option value="rogue">Rogue</option>
			<option value="sorcerer">Sorcerer</option>
		</select>
		<HelperText>{$errors.class?.[0] || ''}</HelperText>
	</div>
	<div class="field">
		<Label for="mode">Mode*</Label>
		<select name="mode" required class={$errors.mode ? 'input-error' : ''}>
			<option value="softcore">Softcore</option>
			<option value="hardcore">Hardcore</option>
		</select>
		<HelperText>{$errors.mode?.[0] || ''}</HelperText>
	</div>
	<div class="field">
		<Label for="tier">Tier*</Label>
		<input
			name="tier"
			type="number"
			min="1"
			max="200"
			required
			placeholder="150"
			class={$errors.tier ? 'input-error' : ''}
		/>
		<HelperText>{$errors.tier?.[0] || ''}</HelperText>
	</div>
	<div class="field">
		<Label for="duration">Time*</Label>
		<input
			name="duration"
			required
			placeholder="12:43"
			class={$errors.duration ? 'input-error' : ''}
		/>
		<HelperText>{$errors.duration?.[0] || ''}</HelperText>
	</div>
	<div class="field">
		<Label for="video">Video*</Label>
		<input
			name="video"
			required
			class={$errors.video ? 'input-error' : ''}
			placeholder="https://youtube.com/watch?v=...."
		/>
		<HelperText>{$errors.video?.[0] || ''}</HelperText>
	</div>
	<div class="field">
		<Label for="build">Build</Label>
		<input name="build" class={$errors.build ? 'input-error' : ''} />
		<HelperText>{$errors.build?.[0] || ''}</HelperText>
	</div>
	<Button type="submit" disabled={$isSubmitting || !$isValid}>Submit</Button>
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

	.input-error {
		border-color: #ff4b58;
	}

	input:not(.input-error):hover {
		border-color: #7abfff;
		background-color: #171920;
	}

	input:not(.input-error):focus {
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

	.input-error {
		border-color: #ff4b58;
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

	select:not(.input-error):hover {
		border-color: #7abfff;
	}

	select:not(.input-error):focus {
		border-color: #148eff;
	}
</style>
