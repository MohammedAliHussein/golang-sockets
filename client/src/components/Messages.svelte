<script>
    import { afterUpdate, beforeUpdate, onMount } from "svelte";
    import Message from "./Message.svelte";

    export let messages = [];
    export let newMessage = null;

    let ready = false;

    $: handleNewMessage(newMessage);

    function handleNewMessage() {
        if (messages.length > 0) {
            messages.unshift(newMessage);
            messages = messages;
        }
    }

    onMount(() => {
        ready = true;
    });
</script>

{#if ready}
    <div class="messages" on:newMessage={handleNewMessage}>
        {#each messages as message}
            <Message {...message}/> 
        {/each}
    </div>
{/if}

<style>
    .messages {
        width: 100%;
        height: 88%;
        overflow: scroll;
        display: flex;
        flex-direction: column-reverse;
        align-items:flex-end;
    }

    ::-webkit-scrollbar {
        width: 0;  
        background: transparent;  
    }
</style>