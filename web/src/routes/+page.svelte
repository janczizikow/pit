<script lang="ts">
	import ClassButton from '$lib/ClassButton.svelte';
	import Table from '$lib/Table.svelte';
	import video from '$lib/assets/hero.webm';
	import type { SubmissionsResponse } from '$lib/types';
	import { writable } from 'svelte/store';

	const PAGE_SIZE = 50;
	let page = 1;
	let data = writable<SubmissionsResponse>({ data: [], metadata: {} });
	let selected = '';

	$: fetch(
		`http://localhost:8080/api/v1/submissions?page=${page}&size=${PAGE_SIZE}&class=${selected.toLowerCase()}&sort=-tier,duration`
	)
		.then((r) => r.json())
		.then((json) => {
			data.set(json);
		});

	function onSelect(cls: string) {
		page = 1;
		if (selected === cls) {
			selected = '';
		} else {
			selected = cls;
		}
	}
</script>

<div class="page">
	<div class="video-container">
		<video autoplay loop muted>
			<source src={video} type="video/webm" />
		</video>
		<h1 class="heading">Solo Pit Ladder</h1>
		<p>
			Best community seasonal pit leaderboard. Rankings are determined by the highest tier level
			achieved and the lowest completion time.
		</p>
		<div class="classes">
			<ClassButton type="BARBARIAN" selected={selected === 'BARBARIAN'} onClick={onSelect} />
			<ClassButton type="DRUID" selected={selected === 'DRUID'} onClick={onSelect} />
			<ClassButton type="NECROMANCER" selected={selected === 'NECROMANCER'} onClick={onSelect} />
			<ClassButton type="ROGUE" selected={selected === 'ROGUE'} onClick={onSelect} />
			<ClassButton type="SORCERER" selected={selected === 'SORCERER'} onClick={onSelect} />
		</div>
		<Table data={$data.data} />
	</div>

	{#each Array($data.metadata?.last_page).fill(null) as _, i}
		<button
			on:click={() => {
				page = i + 1;
			}}>{i + 1}</button
		>
	{/each}
</div>

<style>
	.video-container {
		position: relative;
		width: calc(-16px + 100vw);
		min-height: 100dvh;
		box-shadow: inset 0em -1em 1em rgba(0, 0, 0, 0.8);
	}

	video {
		position: absolute;
		width: 100%;
		height: calc(1px + 100vh);
		object-fit: cover;
		z-index: -1;
	}

	.heading {
		padding-top: 180px;
		margin-bottom: 16px;
		font-weight: var(--fw-400);
		font-family: var(--font-accent);
		font-size: 43px;
		line-height: 1.05;
		text-transform: uppercase;
		text-align: center;
		text-shadow: 3px 5px 5px rgba(0, 0, 0, 0.5);
		color: var(--heading-default);
	}

	p {
		margin: 0 auto;
		max-width: 480px;
		margin-bottom: 32px;
		text-align: center;
		font-family: var(--font-default);
		font-weight: 500;
		font-size: 16px;
		text-shadow: 3px 5px 5px rgba(0, 0, 0, 0.5);
		color: var(--text-default);
	}

	.classes {
		display: flex;
		align-items: center;
		justify-content: center;
		gap: 24px;
		margin-bottom: 64px;
	}
</style>
