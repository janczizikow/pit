<script lang="ts">
	import DiabloButton from '$lib/DiabloButton.svelte';
	import type { PaginationMetadata } from '$lib/api';

	export let metadata: PaginationMetadata;
	export let onChangePage: (page: number) => void;
</script>

<div class="pagination">
	<DiabloButton
		type="primary"
		disabled={metadata?.currentPage === 1 || !metadata.currentPage}
		onClick={() => {
			if (metadata?.currentPage) {
				onChangePage(metadata.currentPage - 1);
			}
		}}
	>
		<b>&lt;</b>
	</DiabloButton>
	<DiabloButton type="primary">{metadata?.currentPage || 1}</DiabloButton>
	<DiabloButton
		type="primary"
		disabled={metadata?.lastPage === metadata?.currentPage}
		onClick={() => {
			if (metadata?.currentPage) {
				onChangePage(metadata.currentPage + 1);
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
