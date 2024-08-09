<script>
    import { AddTask, GetAllTasks, ToggleTaskCompletion, DeleteTask  } from "../wailsjs/go/main/TaskManager.js";
    import AddTaskModal from './AddTaskModal.svelte';
    import DeleteConfirmationModal from './DeleteConfirmationModal.svelte';
    import {onMount} from "svelte";
    let isDeleteModalOpen = false;
    let isAddTaskModalOpen = false;
    let taskToDelete = null;

    let tasks = [];
    let taskTitle = '';
    let priority = 'medium';

    let filteredTasks = [];
    let sortCriteria = 'none';
    let showCompleted = false;

    let deadline = ''


    let filter = 'active'; // 'all', 'active', or 'completed'

    onMount(() => {
        const today = new Date();
        const tomorrow = new Date(today);
        tomorrow.setDate(today.getDate() + 1);
        deadline = tomorrow.toISOString().split('T')[0];
    });

    function setFilter(newFilter) {
        filter = newFilter;
        filterAndSortTasks();
    }

    function init() {
        fetchTasks();
    }



    async function addTask({ title, deadline, priority }) {
        // await DeleteTask(taskToDelete.id);
        // await fetchTasks();
        // closeDeleteModal();
        console.log("ADDING NEW TASK");
        if (title.trim() !== '') {
            await AddTask(title, deadline, priority);
            taskTitle = '';
            // onMount(() => {
            //     const today = new Date();
            //     const tomorrow = new Date(today);
            //     tomorrow.setDate(today.getDate() + 1);
            //     deadline = tomorrow.toISOString().split('T')[0];
            // });
            // priority = 'medium'
            await fetchTasks();
            closeAddTaskModal();
        }
    }

    function openAddTaskModal() {
        console.log("Opening add task")
        isAddTaskModalOpen = true;
    }

    function closeAddTaskModal() {
        isAddTaskModalOpen = false;
    }

    async function toggleTaskCompletion(taskID) {
        await ToggleTaskCompletion(taskID);
        await fetchTasks();
    }

    async function deleteTask(taskID) {
        await DeleteTask(taskID);
        await fetchTasks();
    }

    function openDeleteModal(task) {
        taskToDelete = task;
        isDeleteModalOpen = true;
    }

    function closeDeleteModal() {
        isDeleteModalOpen = false;
        taskToDelete = null;
    }


    async function confirmDelete() {
        if (taskToDelete) {
            await DeleteTask(taskToDelete.id);
            await fetchTasks();
            closeDeleteModal();
        }
    }
    async function fetchTasks() {
        tasks = await GetAllTasks();
        filterAndSortTasks();
    }

    function filterAndSortTasks() {
        if (filter === 'all') {
            filteredTasks = [...tasks];
        } else if (filter === 'active') {
            filteredTasks = tasks.filter(task => !task.completed);
        } else if (filter === 'completed') {
            filteredTasks = tasks.filter(task => task.completed);
        }else{
            filteredTasks = tasks.filter(task => showCompleted || !task.completed);

            if (sortCriteria === 'priority') {
                filteredTasks.sort((a, b) => b.priority - a.priority);
            } else if (sortCriteria === 'deadline') {
                filteredTasks.sort((a, b) => new Date(a.deadline) - new Date(b.deadline));
            }
        }


    }

    // function clearCompleted() {
    //     tasks = tasks.filter(task => !task.completed);
    //     filterAndSortTasks();
    // }

    // function toggleShowCompleted() {
    //     showCompleted = !showCompleted;
    //     filterAndSortTasks();
    // }
    //
    // function setSortCriteria(criteria) {
    //     sortCriteria = criteria;
    //     filterAndSortTasks();
    // }

    init();
</script>

<DeleteConfirmationModal
        isOpen={isDeleteModalOpen}
        onClose={closeDeleteModal}
        onConfirm={confirmDelete}
        taskTitle={taskToDelete?.title || ''}
/>

<AddTaskModal
        isOpen={isAddTaskModalOpen}
        onClose={closeAddTaskModal}
        onAdd={addTask}
    />

<style>
    table {
        width: 50%;
        margin: auto;
        border-collapse: collapse;
        border-radius: 20px;
        background-color: #333;
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
        /*border: 1px solid #ddd;*/
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

    .task-row:last-child {
        border-bottom: none;
    }

    .task {
        display: flex;
        align-items: center;
        justify-content: flex-start;

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
        margin-left: auto;
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

    .app-header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 10px 20px;
        /*background-color: #f0f0f0;*/
        margin-bottom: 20px;
    }

    .app-title {
        margin: 0;
        font-size: 24px;
        color: #f0f0f0;
    }

    .add-task-btn {
        background-color: #4CAF50;
        color: white;
        border: none;
        padding: 10px 20px;
        text-align: center;
        text-decoration: none;
        display: inline-block;
        font-size: 16px;
        margin: 4px 2px;
        cursor: pointer;
        border-radius: 4px;
    }
     .completed {
         text-decoration: line-through;
         color: gray;
     }
    .button-group {
        display: flex;
        gap: 10px;
    }

    .btn {
        background-color: #4CAF50;
        color: white;
        border: none;
        padding: 10px 15px;
        text-align: center;
        text-decoration: none;
        display: inline-block;
        font-size: 14px;
        margin: 4px 2px;
        cursor: pointer;
        border-radius: 4px;
    }

    .add-task-btn {
        background-color: #2196F3;
    }
    .task-container {
        background-color: #333;
        border-radius: 8px;
        padding: 10px;
        margin-bottom: 10px;
    }

    .task-row {
        display: flex;
        align-items: center;
        padding: 10px;
        border-bottom: 1px solid #444;
    }

    .task-row:last-child {
        border-bottom: none;
    }

    .task-title {
        flex-grow: 1;
        color: #f0f0f0;
    }

    .delete-btn {
        background: none;
        border: none;
        color: #888;
        cursor: pointer;
    }

    .filter-bar {
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 10px;
        color: #888;
    }

    .filter-buttons {
        display: flex;
        gap: 10px;
    }

    .filter-btn {
        background: none;
        border: none;
        color: #888;
        cursor: pointer;
    }

    .filter-btn.active {
        color: #4CAF50;
    }

    .clear-completed {
        background: none;
        border: none;
        color: #888;
        cursor: pointer;
    }

</style>
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.0/css/all.min.css" />

<div class="app-header">
    <h1 class="app-title">To Do App</h1>
<!--    <div class="button-group">-->
<!--        <button class="btn" on:click={() => setSortCriteria('none')}>See All</button>-->
<!--        <button class="btn" on:click={() => setSortCriteria('priority')}>Sort by Priority</button>-->
<!--        <button class="btn" on:click={() => setSortCriteria('deadline')}>Sort by Deadline</button>-->
<!--        <button class="btn" on:click={toggleShowCompleted}>-->
<!--            {showCompleted ? 'Hide Completed' : 'Show Completed'}-->
<!--        </button>-->
<!--    <button class="delete-btn" on:click={() => openDeleteModal(task)}>×</button>-->

    <button class="btn add-task-btn" on:click={ openAddTaskModal}>Add Task</button>
<!--    </div>-->
</div>

<div class="task-container">
    {#each filteredTasks as task (task.id)}
        <div class="task-row">
            <input
                    type="checkbox"
                    class="checkbox"
                    checked={task.completed}
                    on:change={() => toggleTaskCompletion(task.id)}
            />

            <div class="task">
                <div class="task-details">
                    <span class="task-title {task.completed ? 'completed' : ''}">{task.title}</span>
                </div>
                <div class="priority">
                    {#if task.priority === 0}
                        <i class="fas fa-arrow-down text-gray-400"></i>
                    {:else if task.priority === 2}
                        <i class="fas fa-arrow-up text-green-500"></i>
                    {:else}
                        <i class="fas fa-exclamation-triangle text-yellow-500"></i>
                    {/if}
                </div>
            </div>

<!--            <span class="task-title {task.completed ? 'completed' : ''}">{task.title}</span>-->
            <button class="delete-btn" on:click={() => openDeleteModal(task)}>×</button>
        </div>
    {/each}

    <div class="filter-bar">
        <span>{filteredTasks.length} tasks </span>
        <div class="filter-buttons">
            <button class="filter-btn {filter === 'all' ? 'active' : ''}" on:click={() => setFilter('all')}>All</button>
            <button class="filter-btn {filter === 'active' ? 'active' : ''}" on:click={() => setFilter('active')}>Active</button>
            <button class="filter-btn {filter === 'completed' ? 'active' : ''}" on:click={() => setFilter('completed')}>Completed</button>
        </div>
<!--        <button class="clear-completed" on:click={clearCompleted}>Clear Completed</button>-->
    </div>
</div>
