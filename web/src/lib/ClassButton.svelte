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
		background-image: url('/classes/hover/BARBARIAN.webp');
	}
	.barbarian:hover {
		background-image: url('/classes/hover/BARBARIAN.webp');
	}
	.druid {
		background-image: url('/classes/base/DRUID.webp');
	}
	.druid-selected {
		background-image: url('/classes/hover/DRUID.webp');
	}
	.druid:hover {
		background-image: url('/classes/hover/DRUID.webp');
	}
	.necromancer {
		background-image: url('/classes/base/NECROMANCER.webp');
	}
	.necromancer-selected {
		background-image: url('/classes/hover/NECROMANCER.webp');
	}
	.necromancer:hover {
		background-image: url('/classes/hover/NECROMANCER.webp');
	}
	.rogue {
		background-image: url('/classes/base/ROGUE.webp');
	}
	.rogue-selected {
		background-image: url('/classes/hover/ROGUE.webp');
	}
	.rogue:hover {
		background-image: url('/classes/hover/ROGUE.webp');
	}
	.sorcerer {
		background-image: url('/classes/base/SORCERER.webp');
	}
	.sorcerer-selected {
		background-image: url('/classes/hover/SORCERER.webp');
	}
	.sorcerer:hover {
		background-image: url('/classes/hover/SORCERER.webp');
	}
</style>
