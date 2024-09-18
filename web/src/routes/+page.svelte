<script lang="ts">
	import ClassButton from '$lib/ClassButton.svelte';
	import Table from '$lib/Table.svelte';
	import video from '$lib/assets/hero.webm';
	import { onDestroy, onMount } from 'svelte';
	import { createSubmissionsStore, PAGE_SIZE } from '../lib/store';
	import Pagination from '$lib/Pagination.svelte';
	import type { Unsubscriber } from 'svelte/store';
	import Heading from '$lib/Heading.svelte';

	const { listSubmissions, selected, selectClass, changePage, data, query } =
		createSubmissionsStore();
	let unsubscribe: Unsubscriber;

	onMount(() => {
		unsubscribe = query.subscribe((q) => {
			listSubmissions({ ...q, classQuery: q.class.toLowerCase() });
		});
	});

	onDestroy(() => {
		unsubscribe?.();
	});
</script>

<div>
	<div class="video-container">
		<video autoplay loop muted>
			<source src={video} type="video/webm" />
		</video>
		<Heading>Solo Pit Ladder</Heading>
		<p>
			Best community seasonal pit leaderboard. Rankings are determined by the highest tier level
			achieved and the lowest completion time.
		</p>
		<div class="classes">
			<ClassButton type="BARBARIAN" selected={$selected === 'BARBARIAN'} onClick={selectClass} />
			<ClassButton type="DRUID" selected={$selected === 'DRUID'} onClick={selectClass} />
			<ClassButton
				type="NECROMANCER"
				selected={$selected === 'NECROMANCER'}
				onClick={selectClass}
			/>
			<ClassButton type="ROGUE" selected={$selected === 'ROGUE'} onClick={selectClass} />
			<ClassButton type="SORCERER" selected={$selected === 'SORCERER'} onClick={selectClass} />
		</div>
		<Pagination metadata={$data.metadata} onChangePage={changePage} />
		<Table data={$data.data} skip={($query.page - 1) * PAGE_SIZE} />
	</div>
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
