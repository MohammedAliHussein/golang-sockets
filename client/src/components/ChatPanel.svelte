<script>
    import { onMount } from "svelte";
    import Connected from "./Connected.svelte";
    import Messages from "./Messages.svelte";
    import MessageInput from "./MessageInput.svelte";

	let ready = false;

    let connectedCount = 0;
    let messages = [];

    onMount(async () => {
        ready = true;
        //retreive all past messages
        //display them
        //open socket connection to send/receive new messages
    });
</script>

{#if ready}
    <div class="chat-panel">
        <Connected connectedCount={connectedCount}/>
        <Messages messages={messages}/>
        <MessageInput />
    </div>
{/if}

<style>
    .chat-panel {
        width: 20%;
        height: 60%;
        border: 1px solid rgba(255, 255, 255, 0.05);
    }
</style>