<script lang="ts">
	import Heading from '$lib/Heading.svelte';
	const url = 'https://discord.com/oauth2/authorize?client_id=1288436640758632481';
</script>

<svelte:head>
	<title>Diablo 4 Pit - Discord bot</title>
</svelte:head>
<div class="container">
	<Heading>Discord Bot</Heading>
	<p>
		"Beta" version of the discord bot that allows you to view pit rankings from your discord server.
		Type <span> <pre>/leaderboard</pre></span> command to view top players. You can further narrow down
		the search by class/mode/season etc. Just like on the website, but directly on your discord server!
	</p>

	<div class="btn">
		<a target="_blank" rel="noopener noreferrer" href={url} class="button link">
			Invite Discord Bot
		</a>
	</div>
	<div class="video-container">
		<video muted controls>
			<source
				src="https://res.cloudinary.com/shanlongjj/video/upload/f_auto:video,q_auto/v1/pit/cqfbq0ml1qvje9prunm4"
				type="video/webm"
			/>
		</video>
	</div>

	<h2 class="center">Why "Beta"?</h2>
	<p>The bot might occasionally go offline and the list of commands is subject to change.</p>
	<p>
		It's still work in progress - I didn't test it fully, but the general functionality is there!
	</p>
</div>

<style>
	.container {
		padding: 0 16px;
		margin: 0 auto;
		width: 100%;
		max-width: var(--container-width);
	}

	.video-container {
		margin: 32px 0;
		position: relative;
		width: 100%;
	}

	.video-container video {
		max-width: var(--container-width);
		height: auto;
	}

	p {
		margin-bottom: 32px;
		text-align: center;
		font-family: var(--font-default);
		font-weight: 500;
		font-size: 16px;
		text-shadow: 3px 5px 5px rgba(0, 0, 0, 0.5);
		color: var(--text-default);
	}

	pre {
		display: inline-block;
	}
	.btn {
		display: flex;
		justify-content: center;
	}

	.link {
		text-decoration: none;
	}

	.center {
		text-align: center;
	}
</style>
