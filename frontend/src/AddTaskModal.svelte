<script>
    import { onMount } from "svelte";
    export let isOpen = false;
    export let onClose = () => {};
    export let onAdd = () => {};

    let title = '';
    let deadline = '';
    let priority = 'medium';

    // onMount(() => {
    $: if (isOpen) {
        const today = new Date();
        const tomorrow = new Date(today);
        tomorrow.setDate(today.getDate() + 1);
        deadline = tomorrow.toISOString().split('T')[0];
    }
    // });

    function closeModal() {
        title = '';
        deadline = '';
        priority = 'medium';

        onClose();
    }

    function handleAdd() {
        onAdd({ title, deadline, priority });
        if (title.trim() !== ''){
            closeModal();
        }
    }
</script>

{#if isOpen}
    <div class="modal-overlay" on:click={closeModal}>
        <div class="modal-content" on:click|stopPropagation>
            <div class="form-container">
                <form class="form-horizontal">
                    <div class="form-group">
                        <label for="title">Task:</label>
                        <input type="text" id="title" bind:value={title} />
                    </div>

                    <div class="form-group">
                        <label for="deadline">Deadline:</label>
                        <input type="date" id="deadline" bind:value={deadline} />
                    </div>

                    <div class="form-group">
                        <label for="priority">Priority:</label>
                        <select id="priority" bind:value={priority}>
                            <option value="low">Low</option>
                            <option value="medium">Medium</option>
                            <option value="high">High</option>
                        </select>
                    </div>
                </form>
            </div>

            <div class="modal-footer">
                <button class="close-btn" on:click={closeModal}>Close</button>
                <button class="add-btn" on:click={handleAdd}>Add</button>
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
        background: rgba(0, 0, 0, 0.6);
        display: flex;
        align-items: center;
        justify-content: center;
        z-index: 1000;
    }

    .modal-content {
        color: #333;
        background-color: white;
        padding: 20px;
        border-radius: 5px;
        max-width: 500px;
        width: 100%;
        box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
        position: relative;
    }

    .form-container {
        margin-bottom: 60px; /* Space for the footer */
    }

    .form-horizontal {
        display: flex;
        gap: 10px; /* Space between each form group */
    }

    .form-group {
        flex: 1; /* Allows each group to grow equally */
        display: flex;
        flex-direction: column;
    }

    form label {
        margin-bottom: 5px;
        font-size: 14px;
    }

    form input,
    form select {
        padding: 8px;
        box-sizing: border-box;
    }

    .modal-footer {
        position: absolute;
        bottom: 20px;
        left: 0;
        width: 100%;
        display: flex;
        justify-content: space-between;
        padding: 0 20px;
    }

    .close-btn,
    .add-btn {
        padding: 10px 20px;
        border: none;
        border-radius: 3px;
        cursor: pointer;
        width: 100px;
    }

    .close-btn {
        background: #ff6347;
        color: white;
    }

    .close-btn:hover {
        background: #ff4500;
    }

    .add-btn {
        background: #4caf50;
        color: white;
        margin-right: 40px;
    }

    .add-btn:hover {

        background: #45a049;
    }
</style>
