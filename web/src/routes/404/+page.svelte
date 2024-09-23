<script lang="ts">
	import Button from '$lib/Button.svelte';
	import Heading from '$lib/Heading.svelte';
	import Text from '$lib/Text.svelte';
</script>

<svelte:head>
	<title>Diablo 4 Pit - Not found</title>
</svelte:head>
<div class="container">
	<Heading>Page Not Found</Heading>
	<Text>The page you are looking for doesn't seem to exist...</Text>
	<div class="btn">
		<a href="/" class="link">
			<Button>Go home</Button>
		</a>
	</div>
</div>

<style>
	.container {
		padding: 0 16px;
		margin: 0 auto;
		width: 100%;
		max-width: var(--container-width);
	}

	.btn {
		display: flex;
		justify-content: center;
	}

	.link {
		text-decoration: none;
	}
</style>
