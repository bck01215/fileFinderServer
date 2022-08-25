<div class="box">
	<div class="box-inner">
		<h2>Introduction</h2>
		<p>
			This app is used to track and manage large files on your filesystem. Whichever user starts the
			app sets the limited permissions of this app. So, if the user only has read access to the file
			he is trying to edit, it will return an insufficient permissions error. Currently there is no
			authentication required, meaning if you are running this app on a server with other users or a
			public port, other users can access this web app and manage the filesystem as the user running
			the app. This app will only descend into directories on the same device/filesystem as the path
			you provide.
		</p>
		<br />
		<h2>How Do I use this App?</h2>
		<p>
			Fill out the form on the <a href="/">home</a> page to search a folder and get all sub-files of
			that folder. If you don't want every file, and you only want the largest x number of files, you
			can pass in a limit to only return the largest x files.
		</p>
	</div>
</div>

<style lang="scss">
	$color-alpha: #1a2b56;
	a {
		color: #000000;
		font-weight: bold;
	}
	a:hover {
		color: #ff5e00;
	}
	.box {
		box-sizing: border-box;
		&:before,
		&:after {
			box-sizing: border-box;
		}
		position: fixed;
		margin-top: 52px;
		top: 50%;
		left: 50%;
		height: 100%;
		transform: translate3d(-50%, -50%, 0);
		background-color: rgba(#ffffff, 0.5);
		width: 100%;
		padding: 5px;
		border: 2px solid $color-alpha;
		&:before,
		&:after {
			content: '•';
			position: absolute;
			width: 14px;
			height: 14px;
			font-size: 14px;
			color: $color-alpha;
			border: 2px solid $color-alpha;
			line-height: 12px;
			top: 5px;
			text-align: center;
		}
		&:before {
			left: 5px;
		}
		&:after {
			right: 5px;
		}
		.box-inner {
			position: relative;
			border: 2px solid $color-alpha;
			padding: 40px;
			&:before,
			&:after {
				content: '•';
				position: absolute;
				width: 14px;
				height: 14px;
				font-size: 14px;
				color: $color-alpha;
				border: 2px solid $color-alpha;
				line-height: 12px;
				bottom: -2px;
				text-align: center;
			}
			&:before {
				left: -2px;
			}
			&:after {
				right: -2px;
			}
		}
	}
</style>
