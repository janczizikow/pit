<script lang="ts">
	import type { Submission } from '$lib/types';
	export let data: Submission[];

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
			<th>
				<span>Rank</span>
			</th>
			<th>
				<span>Player</span>
			</th>
			<th>
				<span>Class</span>
			</th>
			<th>
				<span>Tier</span>
			</th>
			<th>
				<span>Time</span>
			</th>
			<th>
				<span>Video</span>
			</th>
			<th>
				<span>Date</span>
			</th>
		</tr>
	</thead>
	<tbody>
		{#each data as submission, i}
			<tr>
				<td>{i + 1}</td>
				<td>{submission.name}</td>
				<td>{submission.class}</td>
				<td>{submission.tier}</td>
				<td>{formatSeconds(submission.duration)}</td>
				<td>
					<a href={submission.video} target="_blank" rel="noopener noreferrer">{submission.video}</a
					>
				</td>
				<td>{formatDate(submission.created_at)}</td>
			</tr>
		{/each}
	</tbody>
</table>

<style>
	table {
		width: 100%;
		border-collapse: collapse;
		border-spacing: 0;
		color: var(--text-default);
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
</style>
