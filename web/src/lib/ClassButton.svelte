<script>
	export let onClick;
	export let type;
	export let selected;
	let hovered = false;
	function onHover() {
		hovered = true;
	}
	function onBlur() {
		hovered = false;
	}
	let src = `classes/base/${type}.webp`;
	let hoveredSRC = `classes/hover/${type}.webp`;
</script>

{#if selected || hovered}
	<button
		on:click={() => onClick?.(type)}
		on:mouseover={onHover}
		on:focus={onHover}
		on:mouseout={onBlur}
		on:blur={onBlur}
		style="background-image: url({hoveredSRC})"
	></button>
{:else}
	<button
		on:click={() => onClick?.(type)}
		on:mouseover={onHover}
		on:focus={onHover}
		on:mouseout={onBlur}
		on:blur={onBlur}
		style="background-image: url({src})"
	></button>
{/if}

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
