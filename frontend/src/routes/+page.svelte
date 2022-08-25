<script lang="ts">
	interface File {
		size: number;
		path: string;
	}
	let limit = 10;
	let prevLimit = 10; // Used to save last value in case an invalid character is typed
	let errorMSG = '';
	let path = '.';
	let loading = false;
	let fileData: File[] = [];
	const displayBytes = (bytes: number) => {
		let sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
		if (bytes == 0) return '0 Byte';
		let i = Math.floor(Math.log(bytes) / Math.log(1024));
		return Math.round(bytes / Math.pow(1024, i)) + ' ' + sizes[i];
	};

	const submitReq = () => {
		errorMSG = '';
		loading = true;
		fetch('http://localhost:8000/large-files', {
			method: 'POST',
			headers: {
				Accept: 'application/json',
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ limit: limit, path: path })
		})
			.then((response) => response.json())
			.then((data) => {
				if (data['error']) {
					errorMSG = data['error'];
				} else {
					fileData = data;
				}
			})
			.catch((error) => {
				errorMSG = error;
			})
			.finally(() => {
				loading = false;
			});
	};
	const submitDel = (delFile: File) => {
		errorMSG = '';
		fetch('http://localhost:8000/delete-file', {
			method: 'DELETE',
			headers: {
				Accept: 'application/json',
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ size: delFile.size, path: delFile.path })
		})
			.then((response) => response.json())
			.then((data) => {
				if (data['error']) {
					errorMSG = data['error'];
				} else {
					const index = fileData.indexOf(delFile);
					console.log(delFile);
					console.log(fileData.indexOf(delFile));
					if (index > -1) {
						fileData.splice(index, 1);
						fileData = fileData;
					}
				}
			})
			.catch((error) => {
				errorMSG = error;
			});
	};
	const validator = (node: any, value: number) => {
		return {
			update(value: string) {
				limit = value === null || limit < node.min ? prevLimit : parseInt(value);
				prevLimit = limit;
			}
		};
	};
	const clearError = () => {
		errorMSG = '';
	};
</script>

{#if errorMSG !== ''}
	<div class="modal">
		<div class="backdrop" on:click={clearError} />

		<div class="content-wrapper">
			<div>
				<h1>Error Message</h1>
				<div />
			</div>

			<div class="content">
				<p>{errorMSG}</p>
			</div>
		</div>
	</div>
{/if}

{#if !loading}
	<div class="submit">
		<div>
			<label for="limit"> Maximum Number of Returned Files</label>

			<input id="limit" type="number" use:validator={limit} bind:value={limit} min="1" />
		</div>
		<label for="path"> Where to Start Search</label>

		<input id="path" type="text" bind:value={path} />
		<button on:click={submitReq}> Find Files </button>
	</div>
	{#each fileData as file}
		<div class="results">
			<li>
				<input class="path" type="text" disabled value={file.path} />
				<input class="size" type="text" disabled value={displayBytes(file.size)} />
				<button
					on:click={() => {
						submitDel(file);
					}}>Delete</button
				>
			</li>
		</div>
	{/each}
{/if}
{#if loading}
	<div class="loading">
		<p>Loading...</p>
	</div>
{/if}

<style>
	.results input.path {
		width: 75%;
		padding: 12px 20px;
		margin: 8px 0;
		display: inline-block;
		border: 1px solid #ccc;
		border-radius: 4px;
		box-sizing: border-box;
	}
	.results input.size {
		width: 10%;
		padding: 12px 20px;
		margin: 8px 0;
		display: inline-block;
		border: 1px solid #ccc;
		border-radius: 4px;
		box-sizing: border-box;
	}
	.results button {
		width: 14%;
		background-color: #e10e0e;
		color: white;
		padding: 14px 20px;
		border: none;
		border-radius: 4px;
		cursor: pointer;
	}
	.submit input {
		width: 100%;
		padding: 12px 20px;
		margin: 8px 0;
		display: inline-block;
		border: 1px solid #ccc;
		border-radius: 4px;
		box-sizing: border-box;
	}

	.submit button {
		width: 100%;
		background-color: #4caf50;
		color: white;
		padding: 14px 20px;
		margin: 8px 0;
		border: none;
		border-radius: 4px;
		cursor: pointer;
	}

	.submit button:hover {
		background-color: #45a049;
		cursor: pointer;
	}

	.submit {
		border-radius: 5px;
		background-color: #f2f2f2;
		padding: 20px;
	}

	.loading {
		/* animation properties */
		animation-name: bounce;
		animation-duration: 4s;
		animation-iteration-count: infinite;
		animation-timing-function: linear;

		/* other properties */

		width: 300px;
		height: 100px;
		border-radius: 10px;
		margin-top: 10rem;
		margin-left: auto;
		margin-right: auto;
		overflow: hidden;
		text-align: center;
		justify-content: center;
		display: flex;
	}
	.loading p {
		color: white;
		font-size: 2rem;
		top: 0;
		bottom: 0;
		margin: auto;
	}

	@keyframes bounce {
		0% {
			background-color: #1a2b56;
			width: 30rem;
			height: 0px;
			transform: rotate(0deg);
		}
		50% {
			background-color: #ff5e00;
			width: 12rem;
			height: 30rem;
			transform: rotate(360deg);
		}
		100% {
			background-color: #1a2b56;
			width: 30rem;
			height: 0px;
			transform: rotate(720deg);
		}
	}
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
