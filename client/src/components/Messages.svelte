<script>
    import { onMount } from "svelte";
    import Message from "./Message.svelte";

    export let messages = [];

    let ready = false;

    onMount(() => {
        ready = true;
    });
</script>

{#if ready}
    <div class="messages">
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