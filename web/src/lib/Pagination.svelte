<script lang="ts">
	import DiabloButton from '$lib/DiabloButton.svelte';
	import type { PaginationMetadata } from '$lib/types';

	export let metadata: PaginationMetadata;
	export let onChangePage: (page: number) => void;
</script>

<div class="pagination">
	<DiabloButton
		type="primary"
		disabled={metadata?.current_page === 1 || !metadata.current_page}
		onClick={() => {
			if (metadata?.current_page) {
				onChangePage(metadata.current_page - 1);
			}
		}}
	>
		<b>&lt;</b>
	</DiabloButton>
	<DiabloButton type="primary">{metadata?.current_page || 1}</DiabloButton>
	<DiabloButton
		type="primary"
		disabled={metadata?.last_page === metadata?.current_page}
		onClick={() => {
			if (metadata?.current_page) {
				onChangePage(metadata.current_page + 1);
			}
		}}><b>&gt;</b></DiabloButton
	>
</div>

<style>
	.pagination {
		display: flex;
		flex-direction: row;
		gap: 12px;
		align-items: center;
		justify-content: center;
		margin: 48px 0;
	}
</style>
