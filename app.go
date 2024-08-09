package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
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
	db  *sql.DB
	mu  sync.Mutex
	ctx context.Context
}

//type TaskManagerData struct {
//	Next  int
//	Tasks map[int]Task
//}

//// App struct
//type App struct {
//	ctx context.Context
//}

func NewTaskManager() *TaskManager {
	fmt.Println("Opening db 1")

	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("Opening db 2")
	tm := &TaskManager{
		db:  db,
		mu:  sync.Mutex{},
		ctx: context.Background(),
	}
	fmt.Println("Opening db 3")

	_ = tm.createSchema()
	fmt.Println("Opening db 4")

	return tm
}

func (tm *TaskManager) startup(ctx context.Context) {
	//err := tm.createSchema()
	//if err != nil {
	//	return
	//}
	tm.ctx = ctx
}

func (tm *TaskManager) createSchema() error {
	query := `
    CREATE TABLE IF NOT EXISTS tasks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        description TEXT,
        completed BOOLEAN,
        created_at DATETIME,
        deadline DATETIME,
        priority INTEGER
    );
    `
	fmt.Println("Opening db 5")
	_, err := tm.db.Exec(query)
	fmt.Println("Opening db 6")
	return err
}

func (tm *TaskManager) GetAllTasks() ([]Task, error) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	rows, err := tm.db.Query(`SELECT id, title, description, completed, created_at, deadline, priority FROM tasks;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.Deadline, &task.Priority)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	fmt.Println(tasks)
	return tasks, nil
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
	query := `INSERT INTO tasks (title, description, completed, created_at, deadline, priority)
              VALUES (?, ?, ?, ?, ?, ?);`
	_, err = tm.db.Exec(query, title, "", false, time.Now(), parsedTime, priorityInt)
	if err != nil {
		fmt.Println("Error inserting task:", err)
		return false
	}
	fmt.Println("Added task:", title)
	return true
}

func (tm *TaskManager) DeleteTask(id int) bool {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	query := `DELETE FROM tasks WHERE id = ?;`
	_, err := tm.db.Exec(query, id)
	if err != nil {
		fmt.Println("Error deleting task:", err)
		return false
	}

	return true
}

func (tm *TaskManager) ToggleTaskCompletion(taskID int) bool {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	query := `UPDATE tasks SET completed = NOT completed WHERE id = ?;`
	_, err := tm.db.Exec(query, taskID)
	if err != nil {
		fmt.Println("Error toggling completion:", err)
		return false
	}

	return true
}

func (tm *TaskManager) Close() error {
	return tm.db.Close()
}

//func (tm *TaskManager) saveTasks() error {
//
//	data := TaskManagerData{
//		Next:  tm.Next,
//		Tasks: tm.tasks,
//	}
//
//	jsonData, err := json.Marshal(data)
//	if err != nil {
//		fmt.Println("Error marshalling tasks.json")
//		return err
//	}
//
//	return os.WriteFile(taskFile, jsonData, 0644)
//}
//
//func (tm *TaskManager) loadTasks() error {
//	if _, err := os.Stat(taskFile); os.IsNotExist(err) {
//		return nil
//	}
//
//	data, err := os.ReadFile(taskFile)
//	if err != nil {
//		return err
//	}
//
//	var loadedData TaskManagerData
//	if err := json.Unmarshal(data, &loadedData); err != nil {
//		return err
//	}
//
//	tm.Next = loadedData.Next
//	tm.tasks = loadedData.Tasks
//	return nil
//}
