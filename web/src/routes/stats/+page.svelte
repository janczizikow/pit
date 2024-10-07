<script lang="ts">
	import Heading from '$lib/Heading.svelte';
	import Text from '$lib/Text.svelte';
	import { createQuery } from '@tanstack/svelte-query';
	import { writable } from 'svelte/store';
	import { getInt } from '$lib/utils';
	import PieChart from '$lib/PieChart.svelte';
	import { type GetSeasonStatistics200Response, StatisticsApi } from '$lib/api';

	const seasonId = writable(6);
	$: query = createQuery<GetSeasonStatistics200Response>({
		queryKey: ['seasonStatistics', { id: $seasonId }],
		queryFn: () => new StatisticsApi().getSeasonStatistics({ id: $seasonId })
	});
	const onChangeSeason = (season: string) => {
		seasonId.set(getInt(season, 6));
	};
</script>

<svelte:head>
	<title>Diablo 4 Pit - Statistics</title>
</svelte:head>
<div class="container">
	<Heading>Statistics</Heading>
	<Text>Aggregate and per class statistics for a given season.</Text>
	<select
		name="season"
		class="season-selector"
		aria-label="Season"
		value={`${$seasonId}`}
		on:change={(e) => {
			onChangeSeason(e.currentTarget.value);
		}}
	>
		<option value="4" selected={$seasonId === 4}>Season 4</option>
		<option value="5" selected={$seasonId === 5}>Season 5</option>
		<option value="6" selected={$seasonId === 6}>Season 6</option>
	</select>

	<h2>Totals</h2>
	<div class="grid">
		<div class="card">
			<h4 class="card-meta">Total Submission</h4>
			<span class="card-title">{$query.data?.totals.totalSubmissions || 0}</span>
		</div>
		<div class="card">
			<h4 class="card-meta">Unique Players</h4>
			<span class="card-title">{$query.data?.totals.uniquePlayerCount || 0}</span>
		</div>
		<div class="card">
			<h4 class="card-meta">Max tier</h4>
			<span class="card-title">{$query.data?.totals.maxTier || 0}</span>
		</div>
		<div class="card">
			<h4 class="card-meta">Average tier</h4>
			<span class="card-title">{$query.data?.totals.averageTier || 0}</span>
		</div>
	</div>
	<h2>Per class</h2>
	<PieChart data={$query.data?.data || []} />
	<div class="grid-classes">
		{#if $query.isLoading}
			<div>Loading...</div>
		{:else if $query.isError}
			<div>Error</div>
		{:else if $query.isSuccess}
			{#each $query.data?.data as dataPoint}
				<div class="card">
					<h4 class="card-meta">{dataPoint._class?.toUpperCase()}</h4>
					<div class="wrap">
						<span class="class-meta">Total Submissions: {dataPoint.totalSubmissions}</span>
						<span class="class-meta">Unique Players: {dataPoint.uniquePlayerCount}</span>
					</div>
					<div class="wrap">
						<span class="class-meta">Max Tier: {dataPoint.maxTier}</span>
						<span class="class-meta">Avg Tier: {dataPoint.averageTier}</span>
					</div>
					<div class="wrap">
						<span class="class-meta">% of total: {dataPoint.percentageTotal}</span>
						<span class="class-meta">% of unique: {dataPoint.percentageUnique}</span>
					</div>
				</div>
			{/each}
		{/if}
	</div>
</div>

<style>
	.container {
		padding: 0 16px;
		margin: 0 auto;
		width: 100%;
		max-width: 1200px;
	}

	.season-selector {
		margin: 0 auto 24px auto;
		max-width: 120px;
	}

	.grid {
		margin-top: 32px;
		display: grid;
		grid-template-rows: 1fr;
		gap: 8px;
	}

	@media only screen and (min-width: 562px) {
		.grid {
			grid-template-columns: repeat(2, 1fr);
		}
	}

	@media only screen and (min-width: 768px) {
		.grid {
			grid-template-columns: repeat(4, 1fr);
		}
	}

	.grid-classes {
		margin-top: 24px;
		display: grid;
		gap: 8px;
	}

	@media only screen and (min-width: 768px) {
		.grid-classes {
			grid-template-columns: repeat(2, 1fr);
		}
	}

	@media only screen and (min-width: 992px) {
		.grid-classes {
			grid-template-columns: repeat(3, 1fr);
		}
	}

	.card {
		padding: 16px;
		margin-bottom: 24px;
		width: 100%;
		background-color: hsl(210, 14%, 7%);
		border: 1px solid hsl(210, 14%, 13%);
		border-radius: 12px;
		box-shadow: 0px 4px 6px rgba(0, 0, 0, 0.2);
		overflow: hidden;
	}

	.card-meta {
		text-align: center;
		margin-top: 0;
		letter-spacing: 0;
		color: hsl(215, 15%, 75%);
		font-weight: 700;
		font-size: 16px;
		line-height: 1.5;
	}

	.card-title {
		display: block;
		text-align: center;
		font-size: 32px;
		line-height: 1.5;
		color: #fff;
		letter-spacing: 0;
		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
		margin-bottom: 0.35em;
		font-weight: 700;
	}

	.class-meta {
		color: #fff;
	}

	.wrap {
		margin-bottom: 8px;
		display: flex;
		align-items: center;
		justify-content: space-between;
	}
</style>
