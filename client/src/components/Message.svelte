<script>
    import { onMount } from "svelte";

    export let displayName;
    export let messageTime;
    export let message;

    let ready = false;

    onMount(() => {
        ready = true
    });
</script>

{#if ready}
    <div class="message">
        <div class="name-and-time">
            <h6>{displayName} [{messageTime}]</h6>
        </div>
        <h6 class="messageText">{message}</h6>
    </div>
{/if}

<style>
    * {
        background: none;
    }

    .message {
        width: 100%;
        height: fit-content;
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
        border-top: 1px solid rgba(255, 255, 255, 0.05);
        padding: 3px 0px;
    }

    .name-and-time {
        width: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
    }

    h6 {
        color: rgba(255, 255, 255, 0.701);
        margin: 1px;
        font-weight: 300;
    }
</style>