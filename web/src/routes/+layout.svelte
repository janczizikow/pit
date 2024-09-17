<script>
	import '../app.css';
	import Header from './Header.svelte';
	import Footer from './Footer.svelte';
</script>

<div class="app">
	<Header />

	<main>
		<slot />
	</main>

	<Footer />
</div>

<style>
	.app {
		min-height: 100vh;
		display: flex;
		flex-direction: column;
	}

	main {
		flex: 1 0 auto;
	}
</style>
