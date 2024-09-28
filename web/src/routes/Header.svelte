<script>
	import { page } from '$app/stores';
	import logo from '$lib/assets/logo.svg';
</script>

<header>
	<a href="/" class="center">
		<img src={logo} alt="Diablo IV Pit" class="logo" />
	</a>

	<nav class="nav">
		<a href="/stats" class="link">
			<div class="nav-link-content">
				<div class="nav-link-text">Stats</div>
			</div>
		</a>
		<a href="/discord" class="link">
			<div class="nav-link-content">
				<div class="nav-link-text">Discord Bot</div>
			</div>
		</a>
	</nav>

	{#if !$page.url.pathname.includes('/submission')}
		<a href="/submission" class="button">Submit new run</a>
	{/if}
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

	.center {
		display: flex;
		align-items: center;
		justify-content: center;
	}

	@supports ((-webkit-backdrop-filter: blur(20px)) or (backdrop-filter: blur(20px))) {
		header {
			backdrop-filter: blur(20px);
			background: rgba(36, 37, 38, 0.8);
		}
	}

	.logo {
		width: 48px;
		height: 48px;
	}

	.nav {
		display: none;
		flex: 1 1 auto;
		font-family: 'Poppins', Helvetica, Arial, sans-serif;
		color: #fff;
	}
	@media only screen and (min-width: 768px) {
		.nav {
			display: flex;
		}
	}

	.nav-link-content {
		padding: 0 8px;
		align-items: center;
		border-radius: 8px;
		box-sizing: border-box;
		display: flex;
		min-height: 48px;
		overflow: hidden;
		text-align: left;
		transition: var(--transition-hover);
	}

	.link:is(:hover, :focus) .nav-link-content {
		background-color: rgba(255, 255, 255, 0.05);
	}

	.nav-link-text {
		font-size: 18px;
		font-weight: 700;
		overflow: hidden;
		text-overflow: ellipsis;
		white-space: collapse;
	}

	.link {
		background: none;
		border: none;
		color: inherit;
		display: flex;
		font-family: inherit;
		font-size: 18px;
		font-weight: inherit;
		padding: 0px;
		position: relative;
		text-decoration: none;
	}
</style>
