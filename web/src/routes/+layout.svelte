<script>
	import '../app.css';
	import Header from './Header.svelte';
	import Footer from './Footer.svelte';
	import { browser } from '$app/environment';
	import { QueryClient, QueryClientProvider } from '@tanstack/svelte-query';
	const queryClient = new QueryClient({
		defaultOptions: {
			queries: {
				retry: false,
				enabled: browser,
				refetchOnWindowFocus: false
			}
		}
	});
</script>

<div class="app">
	<Header />

	<main>
		<QueryClientProvider client={queryClient}>
			<slot />
		</QueryClientProvider>
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
