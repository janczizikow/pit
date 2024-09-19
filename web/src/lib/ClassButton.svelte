<script lang="ts">
	import { page } from '$app/stores';
	import { afterUpdate, onMount } from 'svelte';
	import { writable } from 'svelte/store';

	export let type: string;
	export let selected: boolean;

	let url = writable('/');
	const src = `classes/base/${type}.webp`;
	const hoveredSRC = `classes/hover/${type}.webp`;

	function syncURLState() {
		const link = new URL($page.url);
		if (selected) {
			link.searchParams.delete('class');
			link.searchParams.delete('page');
		} else {
			link.searchParams.set('class', type);
			link.searchParams.set('page', '1');
		}
		url.set(link.toString());
	}

	onMount(() => {
		syncURLState();
	});

	afterUpdate(() => {
		syncURLState();
	});

	let hovered = false;
	function onHover() {
		hovered = true;
	}
	function onBlur() {
		hovered = false;
	}
</script>

<a
	href={$url}
	on:mouseover={selected ? null : onHover}
	on:focus={selected ? null : onHover}
	on:mouseout={selected ? null : onBlur}
	on:blur={selected ? null : onBlur}
	style="background-image: url({selected || hovered ? hoveredSRC : src})"
>
	<span class="hidden">{type}</span></a
>

<style>
	a {
		all: unset;
		cursor: pointer;
		height: 64px;
		width: 64px;
		background-size: cover;
	}

	@media only screen and (max-width: 800px) {
		a {
			height: 56px;
			width: 56px;
		}
	}

	.hidden {
		display: none;
	}
</style>
