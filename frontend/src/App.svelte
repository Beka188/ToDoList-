<script>
    import { AddTask, GetAllTasks, ToggleTaskCompletion, DeleteTask  } from "../wailsjs/go/main/TaskManager.js";
    import Modal from './Modal.svelte';

    let tasks = [];
    let taskTitle = '';
    let priority = 1;


    let deadline = ''

    let isModalOpen = false;

    function init() {
        fetchTasks();
    }

    async function fetchTasks() {
        tasks = await GetAllTasks();
    }

    async function addTask() {
        if (taskTitle.trim() !== '') {
            await AddTask(taskTitle, deadline, priority);
            taskTitle = '';
            deadline = ''
            priority = 1
            await fetchTasks();
            closeModal();
        }
    }

    function openModal() {
        isModalOpen = true;
    }

    function closeModal() {
        isModalOpen = false;
    }

    async function toggleTaskCompletion(taskID) {
        await ToggleTaskCompletion(taskID);
        await fetchTasks();
    }

    async function deleteTask(taskID) {
        await DeleteTask(taskID);
        await fetchTasks();
    }

    init();
</script>


<Modal isOpen={isModalOpen} onClose={closeModal}>
    <h2>Add New Task</h2>
    <input
            bind:value={taskTitle}
            placeholder="Enter task title"
            class="input"
    />
    <label for="deadline">Deadline:</label>
    <input type="date" id="deadline" bind:value={deadline} />

    <label for="priority">Priority:</label>
    <select id="priority" bind:value={priority}>
        <option value="">Select priority</option>
        <option value="low">Low</option>
        <option value="medium">Medium</option>
        <option value="high">High</option>
    </select>

    <button class="btn" on:click={addTask}>Add Task</button>
</Modal>

<style>
    table {
        width: 50%;
        margin: auto;
        border-collapse: collapse;
    }

    .task-title {
        color: #333;
        font-weight: bold;
        font-size: 16px;
        margin-right: 10px;
    }

    .task-description {
        font-size: 14px;
        color: #888;
        margin-top: 2px;
    }

    .task-list {
        background-color: white;
        border-collapse: collapse;
        width: 50%;
        margin: 0 auto;
        border: 1px solid #ddd;
        box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.1);
    }

    .task-list tr {
        border-bottom: 1px solid gray;
    }

    .task-list td {
        padding: 10px;
    }

    .checkbox {
        border-radius: 50%;
        width: 20px;
        height: 20px;
        border: 2px solid gray;
        cursor: pointer;
        margin-right: 10px;
    }

    .task-row {
        display: flex;
        align-items: center;
        justify-content: space-between;
        padding: 10px;
        border-bottom: 1px solid gray;
    }

    .task {
        display: flex;
        align-items: center;
        justify-content: flex-start; /* Align items to the start */

    }

    .task-details {
        display: flex;
        flex-direction: column;
        margin-left: 10px;
    }

    .priority {
        margin-right: 10px;
    }

    .delete-btn {
        color: white;
        background-color: red;
        border: none;
        border-radius: 4px;
        padding: 5px 10px;
        cursor: pointer;
        margin-left: auto; /* Align to the right */
    }

    .priority i {
        font-size: 18px;
        margin-right: 8px;
    }

    .priority .text-gray-400 {
        color: #a0a0a0;
    }

    .priority .text-green-500 {
        color: #48bb78;
    }

    .priority .text-yellow-500 {
        color: #ecc94b;
    }


</style>
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.0/css/all.min.css" />

<h3>To Do App</h3>
<div class="input-box" id="input">
    <button class="btn" on:click={openModal}>Add Task</button>
</div>

<table class="task-list">
    <tbody>
    {#each tasks as task (task.id)}
        <tr class="task-row">
            <td>
                <input
                        type="checkbox"
                        class="checkbox"
                        checked={task.completed}
                        on:change={() => toggleTaskCompletion(task.id)}
                />
            </td>
            <td class="task">
                <div class="task-details">
                    <div class="task-title">{task.title}</div>
                    <div class="task-description">{task.description}</div>
                </div>
                <div class="priority">
                    {#if task.priority === 1}
                        <i class="fas fa-arrow-down text-gray-400"></i>
                    {:else if task.priority === 2}
                        <i class="fas fa-arrow-up text-green-500"></i>
                    {:else}
                        <i class="fas fa-exclamation-triangle text-yellow-500"></i>
                    {/if}
                </div>
            </td>
            <td>
                <button class="delete-btn" on:click={() => deleteTask(task.id)}>
                    Delete
                </button>
            </td>
        </tr>
    {/each}
    </tbody>
</table>
