<script lang="ts">
	import Button from './Button.svelte';
	import Input from './Input.svelte';
	import Select from './Select.svelte';
	import type { NewSubmission } from './types';

	export let onSubmit: (data: NewSubmission) => void;
	let form: NewSubmission;
	let isSubmitting = false;
</script>

<form class="container" on:submit|preventDefault={() => onSubmit(form)}>
	<Input placeholder="Name" name="name" value={form?.name ?? ''} on:change={(e) => {}} required />
	<Select name="class" value={form?.class ?? ''} required>
		<option value="barbarian">Barbarian</option>
		<option value="druid">Druid</option>
		<option value="necromancer">Necromancer</option>
		<option value="rogue">Rogue</option>
		<option value="sorcerer">Sorcerer</option>
	</Select>
	<Select name="mode" value={form?.mode ?? ''} required>
		<option value="softcore">Softcore</option>
		<option value="hardcore">Hardcore</option>
	</Select>
	<Input placeholder="Tier" name="tier" value={form?.tier ?? ''} required />
	<Input placeholder="Time" name="time" value={form?.duration ?? ''} required />
	<Input placeholder="Video" name="video" value={form?.video ?? ''} required />
	<Input placeholder="Build" name="build" value={form?.build ?? ''} />
	<Button disabled={isSubmitting}>Submit</Button>
</form>

<style>
	.container {
		margin: 0 auto;
		max-width: var(--container-width);
	}
</style>
