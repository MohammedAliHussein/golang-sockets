<script>
    import { onMount } from "svelte";
    import Connected from "./Connected.svelte";
    import Messages from "./Messages.svelte";
    import MessageInput from "./MessageInput.svelte";

    export let connection = null;
    export let name = "";

	let ready = false;
    let connectedCount = 0;
    let messages = [];
    let newMessage = null;

    function setupAndListen() {
        if(connection != null) {
            connection.addEventListener("message", (event) => {
                if(messages.length === 0 && event.data.length !== 0) {
                    const pastMessages = JSON.parse("[" + (event.data) + "]");
                    for (let i = 0; i < pastMessages.length; i++) {
                        let pastMessage = {
                            displayName: pastMessages[i].sender,
                            messageTime: pastMessages[i].time,
                            message: pastMessages[i].message
                        }
                        messages.unshift(pastMessage);
                    }
                    messages = messages;
                } else {
                    let data = JSON.parse(event.data);
                    newMessage = {
                        displayName: data.sender,
                        messageTime: data.time,
                        message: data.message
                    }
                }
            });
        }
    }

    $: setupAndListen(connection);

    onMount(() => {
        ready = true;
    });
</script>

{#if ready}
    <div class="chat-panel">
        <Connected connectedCount={connectedCount}/>
        <Messages messages={messages} newMessage={newMessage}/>
        <MessageInput connection={connection} name={name}/>
    </div>
{/if}

<style>
    .chat-panel {
        width: 20%;
        height: 60%;
        border: 1px solid rgba(255, 255, 255, 0.05);
    }
</style>