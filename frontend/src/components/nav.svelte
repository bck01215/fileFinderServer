<script>
	import { page } from '$app/stores';
</script>

<ul>
	<li>
		<a rel="external" class:active={$page.url.pathname === '/'} href="/">Home</a>
	</li>
	<li>
		<a rel="external" class:active={$page.url.pathname === '/help'} href="/help">Help</a>
	</li>
</ul>

<style>
	ul {
		list-style-type: none;
		margin: 0;
		padding: 0;
		overflow: hidden;
		background-color: #1a2b56;
		border-bottom: solid 2px #ff5e00;
		box-shadow: 0px 0px 5px #ff5e00;
		height: 50px;
	}

	li {
		float: left;
		background-color: transparent;
	}

	li a {
		display: block;
		color: #8f8f8f;
		text-align: center;
		padding: 14px 16px;
		text-decoration: none;
		border-right: solid 0.2px transparent;
	}

	li a:hover {
		color: #ffffff;
		transition-duration: 0.2s;
		text-shadow: 0 0 1px #bebebe, 0 0 5px #bebebe;
	}

	.active {
		color: #ffffff;
		transition-duration: 1s;
		font-weight: bold;
	}
	.active:hover {
		color: #ffffff;
		transition-duration: 0.5s;
	}
</style>
