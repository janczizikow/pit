<script lang="ts">
	import img from '$lib/assets/hc.png';
	export let selected: boolean;
	export let onToggleHC: () => void;
</script>

<button on:click={onToggleHC} class={selected ? 'enabled' : null}>
	<img src={img} alt="hardcore" />
</button>

<style>
	button {
		all: unset;
		cursor: pointer;
		transition-duration: 0.15s;
		transition-property: all;
		transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
		filter: grayscale(100%);
	}

	button:hover {
		filter: grayscale(0%);
	}

	button img {
		width: 32px;
		height: auto;
	}

	.enabled {
		filter: grayscale(0%);
	}
</style>
