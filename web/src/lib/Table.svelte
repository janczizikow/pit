<script lang="ts">
	import type { Submission } from '$lib/types';
	export let data: Submission[];
	export let skip: number;

	function formatSeconds(seconds: number) {
		const minutes = Math.floor(seconds / 60);
		const remainingSeconds = Math.floor(seconds % 60);
		const formattedSeconds = remainingSeconds.toString().padStart(2, '0');
		return `${minutes}:${formattedSeconds}`;
	}
	function formatDate(date: string) {
		return new Date(date).toLocaleDateString();
	}
</script>

<table>
	<thead>
		<tr>
			<th class="rank">
				<span>Rank</span>
			</th>
			<th class="player">
				<span>Player</span>
			</th>
			<th class="class">
				<span>Class</span>
			</th>
			<th class="tier">
				<span>Tier</span>
			</th>
			<th class="time">
				<span>Time</span>
			</th>
			<th class="link-column">
				<span>Video</span>
			</th>
			<th class="link-column">
				<span>Build</span>
			</th>
			<th>
				<span>Date</span>
			</th>
		</tr>
	</thead>
	<tbody>
		{#each data as submission, i}
			<tr>
				<td class="rank">{i + 1 + skip}</td>
				<td class="player" title={submission.name}>{submission.name}</td>
				<td class="class">{submission.class}</td>
				<td class="tier">{submission.tier}</td>
				<td class="time">{formatSeconds(submission.duration)}</td>
				<td class="link-column">
					<a href={submission.video} target="_blank" rel="noopener noreferrer">{submission.video}</a
					>
				</td>
				<td class="link-column">
					{#if submission.build}
						<a href={submission.build} target="_blank" rel="noopener noreferrer"
							>{submission.build}</a
						>
					{:else}
						-
					{/if}
				</td>
				<td>{formatDate(submission.created_at)}</td>
			</tr>
		{/each}
	</tbody>
</table>

<style>
	table {
		margin: 0 auto;
		max-width: 720px;
		width: 100%;
		border-collapse: collapse;
		border-spacing: 0;
		text-align: left;
		color: hsla(0, 0%, 100%, 0.8);
		/* color: var(--text-default); */
		background-color: #000;
	}

	table thead th {
		white-space: nowrap;
	}

	th {
		text-align: left;
	}

	table thead th span {
		display: block;
		padding: 0 15px;
		color: var(--heading-default);
		text-transform: uppercase;
		font-size: 14px;
		font-weight: var(--fs-800);
		text-shadow: 3px 5px 5px rgba(0, 0, 0, 0.5);
		height: 30px;
		line-height: 30px;
		background: #300a00 url('/table-header.png') 0 100% repeat-x;
		border-left: 1px solid #582c19;
		border-right: 1px solid #390b00;
		border-bottom: 1px solid #68381f;
		border-top: 1px solid #62351f;
		zoom: 1;
	}

	table tbody td,
	table tfoot td {
		padding: 15px;
		border-bottom: 1px solid #28241d;
		white-space: nowrap;
		overflow: hidden;
		text-overflow: ellipsis;
	}

	.rank {
		width: 64px;
		max-width: 64px;
	}

	.player {
		width: 156px;
		max-width: 156px;
	}

	.class {
		width: 156px;
		max-width: 156px;
		text-transform: capitalize;
	}

	.tier {
		width: 156px;
		max-width: 156px;
	}

	.time {
		width: 156px;
		max-width: 156px;
	}

	.link-column {
		width: 96px;
		max-width: 96px;
	}
</style>
