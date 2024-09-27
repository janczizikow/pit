<script lang="ts">
	export let type: string;
	export let selected: boolean;
	export let onSelectClass: (cls: string) => void;
</script>

<button class={selected ? `${type}-selected` : type} on:click={() => onSelectClass(type)}>
	<span class="hidden">{type}</span></button
>

<style>
	button {
		all: unset;
		cursor: pointer;
		height: 64px;
		width: 64px;
		background-size: cover;
		transition:
			opacity 350ms,
			filter 350ms;
	}

	@media only screen and (max-width: 800px) {
		button {
			height: 56px;
			width: 56px;
		}
	}

	.hidden {
		display: none;
		visibility: hidden;
	}

	.barbarian {
		background-image: url('/classes/base/BARBARIAN.webp');
	}
	.barbarian-selected {
		position: relative;
		background-image: url('/classes/hover/BARBARIAN.webp');
		filter: brightness(1.6);
	}
	.barbarian:hover {
		background-image: url('/classes/hover/BARBARIAN.webp');
		filter: brightness(1.6);
	}
	.druid {
		background-image: url('/classes/base/DRUID.webp');
	}
	.druid-selected {
		background-image: url('/classes/hover/DRUID.webp');
		filter: brightness(1.6);
	}
	.druid:hover {
		background-image: url('/classes/hover/DRUID.webp');
		filter: brightness(1.6);
	}
	.necromancer {
		background-image: url('/classes/base/NECROMANCER.webp');
	}
	.necromancer-selected {
		background-image: url('/classes/hover/NECROMANCER.webp');
		filter: brightness(1.6);
	}
	.necromancer:hover {
		background-image: url('/classes/hover/NECROMANCER.webp');
		filter: brightness(1.6);
	}
	.rogue {
		background-image: url('/classes/base/ROGUE.webp');
	}
	.rogue-selected {
		background-image: url('/classes/hover/ROGUE.webp');
		filter: brightness(1.6);
	}
	.rogue:hover {
		background-image: url('/classes/hover/ROGUE.webp');
		filter: brightness(1.6);
	}
	.sorcerer {
		background-image: url('/classes/base/SORCERER.webp');
	}
	.sorcerer-selected {
		background-image: url('/classes/hover/SORCERER.webp');
		filter: brightness(1.6);
	}
	.sorcerer:hover {
		background-image: url('/classes/hover/SORCERER.webp');
		filter: brightness(1.6);
	}
	.spiritborn {
		background-image: url('/classes/base/SPIRITBORN.webp');
	}
	.spiritborn-selected {
		background-image: url('/classes/hover/SPIRITBORN.webp');
		filter: brightness(1.6);
	}
	.spiritborn:hover {
		background-image: url('/classes/hover/SPIRITBORN.webp');
		filter: brightness(1.6);
	}
</style>
