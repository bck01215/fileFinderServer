<script lang="ts">
</script>

<div class="modal">
	<div class="backdrop" />

	<div class="content-wrapper">
		<div class="close">
			<slot name="closeChar">X</slot>
		</div>
		<slot name="header">
			<!-- fallback -->
			<div />
		</slot>

		<div class="content">
			<slot name="content" />
		</div>
	</div>
</div>

<style>
	div.modal {
		position: fixed;
		top: 0;
		left: 0;
		width: 100%;
		height: 100vh;

		display: flex;
		justify-content: center;
		align-items: center;
	}
	div.backdrop {
		position: absolute;
		width: 100%;
		height: 100%;
		background-color: rgba(0, 0, 0, 0.4);
	}
	div.content-wrapper {
		z-index: 10;
		max-width: 70vw;
		border-radius: 0.3rem;
		background-color: #191919;
		color: #ff5e00;
		overflow: hidden;
		padding: 10px;
	}
	.content {
		font-weight: bold;
		text-transform: capitalize;
	}
	div.content {
		max-height: 50vh;
		overflow: auto;
	}
</style>
