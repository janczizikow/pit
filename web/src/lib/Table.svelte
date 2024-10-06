<script lang="ts">
	import type { Submission } from '$lib/api';
	import canParseUrl from '$lib/canParseUrl';
	export let data: Submission[];
	export let skip: number;
	export let buildAsText: boolean;

	function formatSeconds(seconds: number) {
		const minutes = Math.floor(seconds / 60);
		const remainingSeconds = Math.floor(seconds % 60);
		const formattedSeconds = remainingSeconds.toString().padStart(2, '0');
		return `${minutes}:${formattedSeconds}`;
	}
	function formatLink(link: string) {
		if (link.startsWith('http')) {
			return link;
		}
		return `https://${link}`;
	}
	function isLink(link: string) {
		if (!link) {
			return false;
		}
		return canParseUrl(link) || link.includes('/');
	}
</script>

<div class="wrapper">
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
					<td class="class">{submission._class}</td>
					<td class="tier">{submission.tier}</td>
					<td class="time">{formatSeconds(submission.duration)}</td>
					<td class="link-column">
						<a href={formatLink(submission.video)} target="_blank" rel="noopener noreferrer"
							>Video</a
						>
					</td>
					<td class="link-column">
						{#if buildAsText || !isLink(submission.build)}
							<span title={submission.build}>{submission.build || '-'}</span>
						{:else}
							<a href={formatLink(submission.build)} target="_blank" rel="noopener noreferrer"
								>Build</a
							>
						{/if}
					</td>
					<td>{submission.createdAt.toLocaleDateString()}</td>
				</tr>
			{/each}
		</tbody>
	</table>
</div>

<style>
	.wrapper {
		overflow-x: auto;
	}

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

	table a {
		color: #39a9f7;
		text-decoration: none;
		transition: all 0.2s ease-out;
		text-decoration: underline;
		text-underline-position: under;
	}

	table a:hover {
		color: #fff;
		text-decoration: none;
	}

	table a:focus {
		outline: 0 none;
		outline-offset: 0;
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
