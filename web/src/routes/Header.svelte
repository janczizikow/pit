<script>
	// import { page } from '$app/stores';
	import logo from '$lib/assets/diablo.svg';
	import Button from '$lib/Button.svelte';
</script>

<header>
	<a href="/">
		<img src={logo} alt="Diablo IV" />
	</a>

	<Button>Submit</Button>
</header>

<style>
	header {
		position: fixed;
		top: 16px;
		left: 16px;
		right: 16px;
		padding: 16px;
		height: 72px;
		display: flex;
		align-items: center;
		justify-content: space-between;
		background-color: rgba(36, 37, 38, 0.95);
		border-radius: 8px;
		z-index: 1;
	}

	@supports ((-webkit-backdrop-filter: blur(20px)) or (backdrop-filter: blur(20px))) {
		header {
			backdrop-filter: blur(20px);
			background: rgba(36, 37, 38, 0.8);
		}
	}
</style>
