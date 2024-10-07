<script lang="ts">
	import { afterNavigate, goto } from '$app/navigation';
	import { page } from '$app/stores';
	import ClassButton from '$lib/ClassButton.svelte';
	import Table from '$lib/Table.svelte';
	import Pagination from '$lib/Pagination.svelte';
	import Heading from '$lib/Heading.svelte';
	import HardcoreButton from '$lib/HardcoreButton.svelte';
	import Text from '$lib/Text.svelte';
	import { createQuery } from '@tanstack/svelte-query';
	import { writable } from 'svelte/store';
	import {
		ListSeasonSubmissions200ResponseFromJSON,
		ListSeasonSubmissionsClassEnum,
		ListSeasonSubmissionsModeEnum,
		SubmissionApi,
		type ListSeasonSubmissions200Response
	} from '$lib/api';
	import preloaded from '$lib/assets/preloaded.json';
	import { getInt } from '$lib/utils';

	const PAGE_SIZE = 50;
	const query = writable<{
		page: number;
		mode: ListSeasonSubmissionsModeEnum;
		class: string;
		season: number;
	}>({
		page: 1,
		mode: ListSeasonSubmissionsModeEnum.Softcore,
		class: '',
		season: 6
	});

	let prevData: ListSeasonSubmissions200Response =
		ListSeasonSubmissions200ResponseFromJSON(preloaded);
	$: queryResult = createQuery({
		queryKey: ['submissions', { ...$query }],
		queryFn: () =>
			new SubmissionApi().listSeasonSubmissions({
				id: $query.season,
				page: $query.page,
				size: PAGE_SIZE,
				mode: $query.mode,
				_class: $query.class as ListSeasonSubmissionsClassEnum,
				sort: '-tier,duration'
			}),
		staleTime: 60 * 1000 // 1min
		// https://github.com/TanStack/query/issues/5913
		// placeholderData: keepPreviousData,
	});
	$: if ($queryResult?.isSuccess) {
		prevData = $queryResult.data;
	}
	const onChangeClass = async (cls: string) => {
		const link = new URL($page.url);
		if ($query.class === cls) {
			link.searchParams.delete('class');
			link.searchParams.delete('page');
		} else {
			link.searchParams.set('class', cls);
			link.searchParams.delete('page');
		}
		await goto(link.toString(), { noScroll: true });
		window.scrollTo({
			top: 200,
			left: 0,
			behavior: 'smooth'
		});
	};
	const onChangePage = async (page: number) => {
		if (page > 0 && !isNaN(page) && Number.isFinite(page)) {
			const url = new URL($page.url);
			url.searchParams.set('page', `${page}`);
			await goto(url, { noScroll: true });
			window.scrollTo({
				top: 200,
				left: 0,
				behavior: 'smooth'
			});
		}
	};
	const onToggleHC = async () => {
		const url = new URL($page.url);
		if (url.searchParams.get('mode')) {
			url.searchParams.delete('mode');
		} else {
			url.searchParams.set('mode', 'hardcore');
		}
		url.searchParams.delete('page');
		await goto(url, { noScroll: true });
		window.scrollTo({
			top: 200,
			left: 0,
			behavior: 'smooth'
		});
	};
	const onChangeSeason = async (s: string) => {
		const url = new URL($page.url);
		url.searchParams.delete('page');
		url.searchParams.set('season', s);
		await goto(url, { noScroll: true });
		window.scrollTo({
			top: 200,
			left: 0,
			behavior: 'smooth'
		});
	};

	afterNavigate(() => {
		const p = $page.url.searchParams.get('page') || '1';
		const classQuery = $page.url.searchParams.get('class') || '';
		const mode = $page.url.searchParams.get('mode') || ListSeasonSubmissionsModeEnum.Softcore;
		const s = $page.url.searchParams.get('season') || '6';

		query.set({
			page: getInt(p, 1),
			class: classQuery,
			mode: mode as ListSeasonSubmissionsModeEnum,
			season: getInt(s, 6)
		});
	});
</script>

<svelte:head>
	<title>Diablo 4 Pit - Leaderboard</title>
	<meta name="description" content="unofficial Diablo 4 Seasonal Pit Leaderboard." />
	<link rel="preload" as="image" href="/buttons/d4-button-filigree-center.webp" />
	<link rel="preload" as="image" href="/classes/hover/BARBARIAN.webp" />
	<link rel="preload" as="image" href="/classes/hover/DRUID.webp" />
	<link rel="preload" as="image" href="/classes/hover/NECROMANCER.webp" />
	<link rel="preload" as="image" href="/classes/hover/ROGUE.webp" />
	<link rel="preload" as="image" href="/classes/hover/SORCERER.webp" />
	<link rel="preload" as="image" href="/classes/hover/SORCERER.webp" />
</svelte:head>
<div>
	<div class="video-container">
		<video autoplay loop muted playsinline>
			<source
				src="https://res.cloudinary.com/shanlongjj/video/upload/f_auto:video,q_auto/v1/pit/yvw0yowe66ytcnpjnsbh"
				type="video/webm"
			/>
		</video>
		<div class="container">
			<Heading>Solo Pit Ladder</Heading>
			<Text>
				Best community seasonal pit leaderboard. Rankings are determined by the highest tier level
				achieved and the lowest completion time.
			</Text>
			<select
				name="season"
				class="season-selector"
				aria-label="Season"
				value={`${$query.season}`}
				on:change={(e) => {
					onChangeSeason(e.currentTarget.value);
				}}
			>
				<option value="4" selected={$query.season === 4}>Season 4</option>
				<option value="5" selected={$query.season === 5}>Season 5</option>
				<option value="6" selected={$query.season === 6}>Season 6</option>
			</select>
			<div class="flex">
				<HardcoreButton selected={$query.mode === 'hardcore'} {onToggleHC} />
				<div class="flex wrap classes">
					<ClassButton
						type={ListSeasonSubmissionsClassEnum.Barbarian}
						selected={$query.class === ListSeasonSubmissionsClassEnum.Barbarian}
						onSelectClass={onChangeClass}
					/>
					<ClassButton
						type={ListSeasonSubmissionsClassEnum.Druid}
						selected={$query.class === ListSeasonSubmissionsClassEnum.Druid}
						onSelectClass={onChangeClass}
					/>
					<ClassButton
						type={ListSeasonSubmissionsClassEnum.Necromancer}
						selected={$query.class === ListSeasonSubmissionsClassEnum.Necromancer}
						onSelectClass={onChangeClass}
					/>
					<ClassButton
						type={ListSeasonSubmissionsClassEnum.Rogue}
						selected={$query.class === ListSeasonSubmissionsClassEnum.Rogue}
						onSelectClass={onChangeClass}
					/>
					<ClassButton
						type={ListSeasonSubmissionsClassEnum.Sorcerer}
						selected={$query.class === ListSeasonSubmissionsClassEnum.Sorcerer}
						onSelectClass={onChangeClass}
					/>
					<ClassButton
						disabled={$query.season < 6}
						type={ListSeasonSubmissionsClassEnum.Spiritborn}
						selected={$query.class === ListSeasonSubmissionsClassEnum.Spiritborn}
						onSelectClass={onChangeClass}
					/>
				</div>
			</div>
		</div>
		<Pagination metadata={$queryResult.data?.metadata || prevData.metadata} {onChangePage} />
		<Table
			data={$queryResult.data?.data || prevData.data}
			skip={($query.page - 1) * PAGE_SIZE}
			buildAsText={$query.season == 4}
		/>
		<Pagination metadata={$queryResult.data?.metadata || prevData.metadata} {onChangePage} />
	</div>
</div>

<style>
	.container {
		padding: 0 16px;
		margin: 0 auto;
		width: 100%;
		max-width: var(--container-width);
	}
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

	.flex {
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.wrap {
		flex-wrap: wrap;
	}

	.season-selector {
		margin: 0 auto 24px auto;
		max-width: 120px;
	}

	.classes {
		padding-left: 32px;
		gap: 24px;
	}
</style>
