<script lang="ts">
	import { arc, pie } from 'd3-shape';
	import { scaleOrdinal } from 'd3-scale';
	import { ascending, sort } from 'd3-array';
	import type { SeasonStatistics } from '$lib/types';

	let el;
	let chartData: SeasonStatistics[];

	export let width = 300;
	let height = 340;
	export let data: SeasonStatistics[] = [];

	// sort the data so we will place the biggest slices first starting from 12 o'clock position.
	$: if (data) {
		chartData = sort(data, (a, b) => ascending(a.class, b.class));
	}

	const colorScale = scaleOrdinal()
		.domain(
			['barbarian', 'druid', 'necromancer', 'rogue', 'sorcerer', 'spiritborn'].sort(ascending)
		)
		.range(['#C69B6D', '#FF7C0A', '#8788EE', '#FFF468', '#3FC7EB', '#00FF98']);

	const pieGenerator = pie<SeasonStatistics>()
		.value((d) => d.percentage_total || 0)
		.sort(null);

	$: pieData = pieGenerator(chartData);

	const arcGenerator = arc<unknown>()
		// 60% of the full radius gives a nice appearance .innerRadius((0.6) => (0.6 * height) / 2.4)
		// To see a pie chart, just change this to zero
		.innerRadius(0)
		// Outer radius is less than the full radius because our labels will sit outside of the donut
		.outerRadius((0.85 * height) / 2.2)
		.padRadius(40)
		.cornerRadius(0);

	const labelArcs = arc<unknown>()
		.innerRadius((0.96 * height) / 2)
		.outerRadius((0.96 * height) / 2);
</script>

<div class="chart">
	<div class="svg-container" bind:clientWidth={width}>
		{#if width}
			<svg viewBox={`0 0 ${width} ${height + 40}`} class="chart">
				<g bind:this={el} transform="translate({width / 2 - 5} {height / 2 + 20})">
					{#each pieData as d, i (d.data.class)}
						<path
							class={`${i}`}
							pointer-events="all"
							cursor="pointer"
							d={arcGenerator(d)}
							{...{
								/* @ts-ignore */
							}}
							fill={colorScale(d.data.class)}
						/>
						<!-- labels -->
						<text
							x="0"
							y="0"
							text-anchor="middle"
							font-size="0.8em"
							class="text"
							transform="translate({labelArcs.centroid(d).join(' ')})"
							>{d.data.class}
						</text>
						<text
							x="0"
							y="1.2em"
							text-anchor="middle"
							font-size="0.8em"
							font-weight="700"
							class="text"
							transform="translate({labelArcs.centroid(d).join(' ')})"
							>{d.data.percentage_total + ' %'}
						</text>
					{/each}
				</g>
			</svg>
		{/if}
	</div>
</div>

<style>
	.chart {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: space-around;
	}

	@media (min-width: 860px) {
		.chart {
			flex-direction: row;
		}
	}

	.svg-container {
		width: 100%;
	}

	.text {
		fill: #f3f4f6;
	}
</style>
