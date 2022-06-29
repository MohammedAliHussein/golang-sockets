<script>
    import { onMount } from "svelte/internal";

    export let connectedCount = 0;

    let ready = false;

    onMount(() => {
        ready = true
    });
</script>

{#if ready}
    <div class="connected">
        {#if connectedCount > 0}
            <div class="green-circle"></div>
        {:else}
            <div class="red-circle"></div>
        {/if}
        <h4>{connectedCount} Online</h4>
    </div>
{/if}

<style>
    .connected {
        width: 100%;
        height: 5%;
        display: flex;
        justify-content: center;
        align-items: center;
        border-bottom: 1px solid rgba(255, 255, 255, 0.05);
    }

    h4 {
        color: white;
        font-size: 12px;
    }

    .green-circle, .red-circle {
        width: 5px;
        height: 5px;
        border-radius: 100%;
        margin-right: 5px;
    }

    .green-circle {
        background-color: greenyellow;
    }

    .red-circle {
        background-color: red;
    }
</style>