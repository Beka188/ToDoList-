package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"sync"
	"time"
)

const taskFile = "tasks.json"

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	Deadline    time.Time `json:"deadline"`
	Priority    int       `json:"priority"`
}

type TaskManager struct {
	tasks map[int]Task
	Next  int `json:"next"`
	mu    sync.Mutex
	ctx   context.Context
}

type TaskManagerData struct {
	Next  int
	Tasks map[int]Task
}

func NewTaskManager() *TaskManager {
	tm := &TaskManager{
		tasks: make(map[int]Task),
		Next:  0,
		mu:    sync.Mutex{},
		ctx:   context.Background(),
	}

	_ = tm.loadTasks()

	return tm

}

func (tm *TaskManager) startup(ctx context.Context) {
	tm.ctx = ctx
}

func (tm *TaskManager) GetAllTasks() []Task {
	tasks := make([]Task, len(tm.tasks))
	idx := 0
	for _, task := range tm.tasks {
		tasks[idx] = task
		idx++
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].CreatedAt.Before(tasks[j].CreatedAt)
	})

	return tasks
}

func (tm *TaskManager) AddTask(title string, deadline string, priority string) bool {
	priorityInt := 1
	switch priority {
	case "low":
		priorityInt = 0
	case "medium":
		priorityInt = 1
	case "high":
		priorityInt = 2
	}
	layout := "2006-01-02"
	parsedTime, err := time.Parse(layout, deadline)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return false
	}
	fmt.Printf("%T\n%v", parsedTime, parsedTime)
	tm.mu.Lock()
	defer tm.mu.Unlock()
	tm.tasks[tm.Next] = Task{ID: tm.Next, Title: title, CreatedAt: time.Now(), Deadline: parsedTime, Priority: priorityInt}
	fmt.Println(tm.Next)
	tm.Next++
	_ = tm.saveTasks()
	fmt.Println(tm.GetAllTasks())

	return true
}

func (tm *TaskManager) DeleteTask(id int) bool {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	_, exists := tm.tasks[id]
	if !exists {
		return false
	}

	delete(tm.tasks, id)
	_ = tm.saveTasks()
	return true
}

func (tm *TaskManager) ToggleTaskCompletion(taskID int) bool {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	task, exists := tm.tasks[taskID]
	if !exists {
		return false
	}

	task.Completed = !task.Completed
	tm.tasks[taskID] = task
	_ = tm.saveTasks()
	return true
}

func (tm *TaskManager) saveTasks() error {

	data := TaskManagerData{
		Next:  tm.Next,
		Tasks: tm.tasks,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling tasks.json")
		return err
	}

	return os.WriteFile(taskFile, jsonData, 0644)
}

func (tm *TaskManager) loadTasks() error {
	if _, err := os.Stat(taskFile); os.IsNotExist(err) {
		return nil
	}

	data, err := os.ReadFile(taskFile)
	if err != nil {
		return err
	}

	var loadedData TaskManagerData
	if err := json.Unmarshal(data, &loadedData); err != nil {
		return err
	}

	tm.Next = loadedData.Next
	tm.tasks = loadedData.Tasks
	return nil
}
