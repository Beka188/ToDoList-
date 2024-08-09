<script>
    export let isOpen = false;
    export let onClose = () => {};
    export let onConfirm = () => {};
    export let taskTitle = '';
</script>


{#if isOpen}
    <div class="modal-overlay">
        <div class="modal-content">
            <h2>Confirm Deletion</h2>
            <p>Are you sure you want to delete the task "{taskTitle}"?</p>
            <div class="button-group">
                <button on:click={onClose}>Cancel</button>
                <button on:click={onConfirm} class="delete-btn">Delete</button>
            </div>
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
        background-color: rgba(0, 0, 0, 0.5);
        display: flex;
        justify-content: center;
        align-items: center;
    }

    .modal-content {
        color: #333;
        background-color: white;
        padding: 20px;
        border-radius: 5px;
        text-align: center;
    }

    .button-group {
        margin-top: 20px;
    }

    button {
        margin: 0 10px;
        padding: 5px 10px;
        border: none;
        border-radius: 4px;
        cursor: pointer;
    }

    .delete-btn {
        background-color: red;
        color: white;
    }
</style>