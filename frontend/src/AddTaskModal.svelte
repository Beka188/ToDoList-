<script>
    import {onMount} from "svelte";
    export let isOpen = false;
    export let onClose = () => {};
    export let onAdd = () => {};

    let title = '';
    let deadline = '';
    let priority = 'medium';

    onMount(() => {
        if (isOpen) {
            const today = new Date();
            const tomorrow = new Date(today);
            tomorrow.setDate(today.getDate() + 1);
            deadline = tomorrow.toISOString().split('T')[0];
        }
    });

    function closeModal() {
        title = '';
        deadline = '';
        priority = 'medium';


        onClose();
    }

    function handleAdd() {
            onAdd({ title, deadline, priority });
          closeModal();
        }

</script>

{#if isOpen}
    <div class="modal-overlay" on:click={closeModal}>
        <div class="modal-content" on:click|stopPropagation>
            <form>
                <label for="title">Title:</label>
                <input type="text" id="title" bind:value={title} />

                <label for="deadline">Deadline:</label>
                <input type="date" id="deadline" bind:value={deadline} />

                <label for="priority">Priority:</label>
                <select id="priority" bind:value={priority}>
                    <option value="low">Low</option>
                    <option value="medium">Medium</option>
                    <option value="high">High</option>
                </select>
            </form>

            <div class="modal-footer">
                <button class="close-btn" on:click={closeModal}>Close</button>
                <button class="add-btn" on:click={handleAdd}>Add</button>
            </div>
        </div>
    </div>
{/if}

<style>
    .input, .btn {
        margin-top: 10px;
        display: block;
        width: 100%;
    }
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

<!--<script>-->
<!--    export let isOpen = false;-->
<!--    export let onClose = () => {};-->
<!--    export let onAdd = () => {};-->

<!--    let title = '';-->
<!--    let deadline = '';-->
<!--    let priority = 'medium';-->

<!--    function closeModal() {-->
<!--        title = '';-->
<!--        deadline = '';-->
<!--        priority = 'medium';-->
<!--        onClose();-->
<!--    }-->

<!--    function handleAdd() {-->
<!--        onAdd({ title, deadline, priority });-->
<!--        closeModal();-->
<!--    }-->
<!--</script>-->

<!--{#if isOpen}-->
<!--    <div class="modal-overlay" on:click={closeModal}>-->
<!--        <div class="modal-content" on:click|stopPropagation>-->
<!--            <form>-->
<!--                <label for="title">Title:</label>-->
<!--                <input type="text" id="title" bind:value={title} />-->

<!--                <label for="deadline">Deadline:</label>-->
<!--                <input type="date" id="deadline" bind:value={deadline} />-->

<!--                <label for="priority">Priority:</label>-->
<!--                <select id="priority" bind:value={priority}>-->
<!--                    <option value="low">Low</option>-->
<!--                    <option value="medium">Medium</option>-->
<!--                    <option value="high">High</option>-->
<!--                </select>-->
<!--            </form>-->

<!--            <div class="modal-footer">-->
<!--                <button class="close-btn" on:click={closeModal}>Close</button>-->
<!--                <button class="add-btn" on:click={handleAdd}>Add</button>-->
<!--            </div>-->
<!--        </div>-->
<!--    </div>-->
<!--{/if}-->
