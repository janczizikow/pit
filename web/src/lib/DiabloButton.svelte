<script lang="ts">
	export let onClick: (() => void) | null = null;
</script>

<button on:click={onClick} {...$$restProps}>
	<slot />
</button>

<style>
	button {
		position: relative;
		text-transform: uppercase;
		text-shadow: 3px 5px 5px rgba(0, 0, 0, 0.5);
		color: #e5e0c8;
		filter: drop-shadow(0px 0px 33px rgba(0, 0, 0, 0.66))
			drop-shadow(0px 4px 15px rgba(0, 0, 0, 0.66));
		background-color: transparent;
		transition: color 350ms;
		outline: none;
		border: none;
		min-width: 92px;
		font-family: var(--font-accent);
	}

	button:not(:disabled) {
		cursor: pointer;
	}

	button::part(host-anchor) {
		backdrop-filter: none;
	}

	button[disabled] {
		filter: saturate(0) contrast(0.8);
		cursor: not-allowed;
		backdrop-filter: none;
	}

	/* @media (min-width: 960px) {
		button {
			font-size: 20px;
			line-height: 24.3px;
		}
	} */

	button::before,
	button::after {
		top: 0;
		left: 0;
		right: 0;
		bottom: 0;
		content: '';
		position: absolute;
		z-index: -1;
		transition:
			transform 350ms,
			opacity 350ms,
			color 350ms;
	}
	button:active:focus {
		transform: translateY(3px);
	}

	button::after {
		opacity: 0;
	}

	button:not([disabled]):hover {
		color: #f6f8f0;
	}

	button:not([disabled]):hover::after {
		opacity: 1;
	}

	button[type='default']::before,
	button[type='default']::after {
		background-repeat: no-repeat, no-repeat, repeat-x;
		background-position:
			left center,
			right center,
			center center;
		background-size: auto 100%;
	}

	button[type='default']::before {
		background-image: url(/buttons/d4-button-secondary-left-base.webp),
			url(/buttons/d4-button-secondary-right-base.webp),
			url(/buttons/d4-button-secondary-tile-base.webp);
	}

	button[type='default']::after {
		background-image: url(/buttons/d4-button-secondary-left-hover.webp),
			url(/buttons/d4-button-secondary-right-hover.webp),
			url(/buttons/d4-button-secondary-tile-hover.webp);
	}

	button[type='primary'] {
		padding: 1.8rem 1.6rem;
		font-size: 1.2rem;
	}

	/* @media (min-width: 1250px) {
		button[type='primary'] {
			font-size: 1.5rem;
		}
	}

	@media (max-width: 1630px) {
		button[type='primary'] {
			font-size: 1.2rem;
		}
	}

	@media (max-width: 1390px) {
		button[type='primary'] {
			font-size: 1rem;
		}
	}

	@media (max-width: 1210px) {
		button[type='primary'] {
			font-size: 0.8rem;
		}
	}

	@media (max-width: 960px) {
		button[type='primary'] {
			font-size: 0.8rem;
		}
	}

	@media (max-width: 762px) {
		button[type='primary'] {
			font-size: 0.8rem;
		}
	} */

	button[type='primary']::before,
	button[type='primary']::after {
		background-repeat: no-repeat;
		background-position:
			left center,
			right center,
			center center,
			10px center,
			calc(100% - 10px) center,
			center center;
		background-size:
			auto 100%,
			auto 100%,
			auto 100%,
			auto calc(100% - 18px),
			auto calc(100% - 18px),
			82% calc(100% - 18px);
	}

	@media (min-width: 960px) {
		button[type='primary']::before,
		button[type='primary']::after {
			background-size:
				auto 100%,
				auto 100%,
				auto 100%,
				auto calc(100% - 20px),
				auto calc(100% - 20px),
				82% calc(100% - 20px);
		}
	}

	button[type='primary']::before {
		background-image: url(/buttons/d4-button-filigree-left.webp),
			url(/buttons/d4-button-filigree-right.webp), url(/buttons/d4-button-filigree-center.webp),
			url(/buttons/d4-button-primary-left-base.webp),
			url(/buttons/d4-button-primary-right-base.webp),
			url(/buttons/d4-button-primary-tile-base.webp);
	}

	button[type='primary']::after {
		background-image: url(/buttons/d4-button-filigree-left.webp),
			url(/buttons/d4-button-filigree-right.webp), url(/buttons/d4-button-filigree-center.webp),
			url(/buttons/d4-button-primary-left-hover.webp),
			url(/buttons/d4-button-primary-right-hover.webp),
			url(/buttons/d4-button-primary-tile-hover.webp);
	}
</style>
