<button {...$$restProps}>
	<slot />
</button>

<style>
	button {
		/* reset */
		background: none;
		border: none;
		color: inherit;
		display: flex;
		font-family: inherit;
		font-size: 16px;
		font-weight: inherit;
		padding: 0px;
		position: relative;
		text-decoration: none;
		/* end rest */
		min-height: 48px;
		font-size: 16px;
		align-items: center;
		border-radius: 8px;
		color: #fff;
		background-color: #a50905;
		border: 0;
		filter: brightness(100%);
		font-family: 'Poppins', Helvetica, Arial, sans-serif;
		font-weight: 600;
		justify-content: center;
		padding: 0 24px;
		text-align: center;
		transition:
			color 0.1s ease,
			background-color 0.1s ease,
			text-decoration-color 0.1s ease,
			filter 0.1s ease,
			border-color 0.1s ease;
	}

	button:not(:disabled):hover {
		cursor: pointer;
		background-color: #dd0707;
	}

	button:disabled {
		cursor: not-allowed;
	}
</style>
