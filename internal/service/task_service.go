package service

import (
	"database/sql"
	"fmt"
)

type TaskService struct {
	DB *sql.DB
}

type Task struct {
	Id   int    `json:"task_id"`
	Name string `json:"task_name"`
}

func NewTaskService(db *sql.DB) *TaskService {
	return &TaskService{DB: db}
}

func (s *TaskService) GetTasks() ([]Task, error) {
	rows, err := s.DB.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var task Task
		err = rows.Scan(&task.Id, &task.Name)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s *TaskService) CreateTask(taskName string) (*Task, error) {
	row := s.DB.QueryRow("INSERT INTO tasks(task_name) VALUES ($1) RETURNING task_id", taskName)

	var taskId int
	err := row.Scan(&taskId)
	if err != nil {
		return nil, err
	}

	fmt.Println(taskId, taskName)

	return &Task{taskId, taskName}, nil
}
