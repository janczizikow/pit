<script lang="ts">
</script>

<input {...$$restProps} />

<style>
	input {
		display: block;
		width: 100%;
		padding-top: 8px;
		padding-right: 16px;
		padding-bottom: 8px;
		padding-left: 16px;
		appearance: none;
		outline: none;
		font-size: 1rem;
		font-family: var(--font-default);
		line-height: 1.5;
		color: white;
		background-color: #171920;
		border: 1px solid rgba(255, 255, 255, 0.36);
		border-radius: 4px;
		transition:
			border 0.2s ease,
			background-color 0.2s ease;
		text-overflow: ellipsis;
		white-space: nowrap;
		box-shadow:
			rgba(0, 0, 0, 0.016) 0px 2px 2px 0px,
			rgba(0, 0, 0, 0.008) 0px 0px 0px 1px;

		&:focus {
			border-color: #7abfff;
			border-color: #148eff;
		}

		&::placeholder {
			color: grey;
		}
	}
</style>
