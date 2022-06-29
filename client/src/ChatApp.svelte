<script>
	import ChatPanel from "./components/ChatPanel.svelte";
	import { onMount } from "svelte/internal";

	let ready = false;
	let name = "";

	function handleConfirmName() {

	}	

	onMount(() => {
		ready = true;
	});
</script>

{#if ready}
	<div class="main">
		<h3>[type your display name]</h3>
		<div class="top-section">
			<input placeholder="Display Name" bind:value={name}/>
			<button on:click|preventDefault={handleConfirmName}>Confirm</button>
		</div>
		<ChatPanel />
	</div>
{/if}

<style>
	.main {
		width: 100%;
		height: 100vh;
		display: flex;
		justify-content: center;
		align-items: center;
		flex-direction: column;
	}

    input {
        text-align: center;
        border: none;
        outline: none;
        color: white;
        padding: 5px 8px;
        border: 1px solid rgba(255, 255, 255, 0.089);
		margin: 1em;
		font-size: 11px;
    }

	h3 {
		color: white;
	}

	button {
        color: white;
		outline: none;
		background: none;
		border: 1px solid white;
		cursor: pointer;
		padding: 1px 8px;
		transition: 0.2s cubic-bezier(0, 0.55, 0.45, 1);
		font-size: 12px;
        height: fit-content;
        width: fit-content;
        padding: 3px 13px;
	}

	button:hover {
		background-color: rgba(255, 255, 255, 0.059);
		/* color: black; */
	}
</style>
