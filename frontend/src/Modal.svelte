<script>
    export let isOpen = false;
    export let onClose = () => {
    };

    function closeModal() {
        onClose();
    }
</script>

{#if isOpen}
    <div class="modal-overlay" on:click={closeModal}>
        <div class="modal-content" on:click|stopPropagation>
            <slot></slot>
            <button class="close-btn" on:click={closeModal}>Close</button>
        </div>
    </div>
{/if}

<style>
    .modal-overlay {
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        background: rgba(0, 0, 0, 0.6);
        display: flex;
        align-items: center;
        justify-content: center;
        z-index: 1000;
    }

    .modal-content {
        background: white;
        padding: 20px;
        border-radius: 5px;
        max-width: 500px;
        width: 100%;
        box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
    }

    .close-btn {
        margin-top: 10px;
        padding: 5px 10px;
        border: none;
        background: #ff6347;
        color: white;
        border-radius: 3px;
        cursor: pointer;
    }

    .close-btn:hover {
        background: #ff4500;
    }
</style>