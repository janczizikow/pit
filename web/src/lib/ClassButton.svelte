<script lang="ts">
	export let onClick;
	export let type;
	export let selected;
	const src = `classes/base/${type}.webp`;
	const hoveredSRC = `classes/hover/${type}.webp`;
	let hovered = false;
	function onHover() {
		hovered = true;
	}
	function onBlur() {
		hovered = false;
	}
</script>

<button
	on:click={() => onClick?.(type)}
	on:mouseover={onHover}
	on:focus={onHover}
	on:mouseout={onBlur}
	on:blur={onBlur}
	style="background-image: url({selected || hovered ? hoveredSRC : src})"
></button>

<style>
	button {
		all: unset;
		cursor: pointer;
		height: 64px;
		width: 64px;
		background-size: cover;
	}

	@media only screen and (max-width: 800px) {
		button {
			height: 56px;
			width: 56px;
		}
	}
</style>
