<script lang="ts">
	enum AlertTypes {
		WARNING = 'warning',
		ERROR = 'error',
		SUCCESS = 'success'
	}
	const type: AlertTypes = AlertTypes.ERROR;
	export let title = '';
	export let message: string;
</script>

<div class="alert" role="alert" tabindex="-1">
	<div class="message">
		<div class="title">{title}</div>
		{message}
	</div>
</div>

<style>
	.alert {
		display: flex;
		font-size: 16px;
		color: #ff4b58;
		background-color: rgba(255, 75, 88, 0.1);
		border-radius: 4px;
		margin-bottom: 16px;
	}

	.title {
		font-size: 18px;
		font-weight: 500;
	}

	.message {
		padding: 8px;
	}
</style>
