<script lang="ts">
	import { afterNavigate, goto } from '$app/navigation';
	import { page } from '$app/stores';
	import ClassButton from '$lib/ClassButton.svelte';
	import Table from '$lib/Table.svelte';
	import video from '$lib/assets/hero.webm';
	import { createSubmissionsStore, PAGE_SIZE } from '$lib/store';
	import Pagination from '$lib/Pagination.svelte';
	import Heading from '$lib/Heading.svelte';

	const { listSubmissions, data, query } = createSubmissionsStore({
		page: 1,
		class: ''
	});
	const onChangeClass = async (cls: string) => {
		const link = new URL($page.url);
		if ($query.class === cls) {
			link.searchParams.delete('class');
			link.searchParams.delete('page');
		} else {
			link.searchParams.set('class', cls);
			link.searchParams.set('page', '1');
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
			const url = $page.url;
			url.searchParams.set('page', `${page}`);
			await goto(url, { noScroll: true });
			window.scrollTo({
				top: 200,
				left: 0,
				behavior: 'smooth'
			});
		}
	};

	afterNavigate(() => {
		const p = $page.url.searchParams.get('page') || '1';
		const classQuery = $page.url.searchParams.get('class') || '';
		let page = parseInt(p);
		if (page <= 0 || isNaN(page) || !Number.isFinite(page)) {
			page = 1;
		}
		query.set({
			page,
			class: classQuery
		});
		listSubmissions({
			page,
			classQuery
		});
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
			<ClassButton
				type="barbarian"
				selected={$query.class === 'barbarian'}
				onSelectClass={onChangeClass}
			/>
			<ClassButton type="druid" selected={$query.class === 'druid'} onSelectClass={onChangeClass} />
			<ClassButton
				type="necromancer"
				selected={$query.class === 'necromancer'}
				onSelectClass={onChangeClass}
			/>
			<ClassButton type="rogue" selected={$query.class === 'rogue'} onSelectClass={onChangeClass} />
			<ClassButton
				type="sorcerer"
				selected={$query.class === 'sorcerer'}
				onSelectClass={onChangeClass}
			/>
		</div>
		<Pagination metadata={$data.metadata} {onChangePage} />
		<Table data={$data.data} skip={($query.page - 1) * PAGE_SIZE} />
		<Pagination metadata={$data.metadata} {onChangePage} />
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
	}
</style>
