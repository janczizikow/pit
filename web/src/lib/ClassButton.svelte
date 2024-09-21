<script lang="ts">
	export let type: string;
	export let selected: boolean;
	export let onSelectClass: (cls: string) => void;

	const src = `/classes/base/${type.toUpperCase()}.webp`;
	const hoveredSRC = `/classes/hover/${type.toUpperCase()}.webp`;

	let hovered = false;
	function onHover() {
		hovered = true;
	}
	function onBlur() {
		hovered = false;
	}
</script>

<button
	on:mouseover={selected ? null : onHover}
	on:focus={selected ? null : onHover}
	on:mouseout={selected ? null : onBlur}
	on:blur={selected ? null : onBlur}
	on:click={() => onSelectClass(type)}
	style="background-image: url({selected || hovered ? hoveredSRC : src})"
>
	<img src={selected || hovered ? src : hoveredSRC} alt="" class="hidden" />
	<span class="hidden">{type}</span></button
>

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

	.hidden {
		display: none;
	}
</style>
