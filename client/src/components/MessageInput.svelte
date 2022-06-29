<script>
    import { onMount } from "svelte/internal";

    import moment from "moment";

    export let connection = null;
    export let name = "";

    let ready = false;
    let messageText = "";

    function handleSend() {
        connection.send(JSON.stringify({
                "Sender": `${name}`,
                "Message": `${messageText}`,
                "Time": `${moment().format('MMMM Do YYYY, h:mm:ss a')}` 
            })
        );
    }

    onMount(() => {
        ready = true;
    });
</script>

{#if ready}
    <div class="message-input">
        <input type="text" placeholder="Say something" bind:value={messageText}>
        <button on:click|preventDefault={handleSend}>Send</button>
    </div>
{/if}

<style>
    .message-input {
        width: 100%;
        height: 7%;
        display: flex;
        justify-content: space-between;
        align-items: center;
        background: none;
    }

    input {
        color: white;
        outline: none;
        border: 1px solid rgba(255, 255, 255, 0.091);
        text-align: center;
        padding: 3px 13px;
        transition: 0.2s cubic-bezier(0, 0.55, 0.45, 1);
        font-size: 12px;
        height: 100%;
        width: 75%;
        font-size: 11px;
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
        height: 100%;
        width: 25%;
	}

	button:hover {
		background-color: rgba(255, 255, 255, 0.059);
		/* color: black; */
	}
</style>