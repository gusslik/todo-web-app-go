package service

import (
	"database/sql"
	"net/http"
	custom_error "todo-web-app-go/internal/error"

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
	if taskName == "" {
		return nil, custom_error.NewApiError(http.StatusUnprocessableEntity, "request body doesn't have task name")
	}
	row := s.DB.QueryRow("INSERT INTO tasks(task_name) VALUES ($1) RETURNING task_id", taskName)

	var taskId int
	err := row.Scan(&taskId)
	if err != nil {
		return nil, err
	}

	return &Task{taskId, taskName}, nil
}

func (s *TaskService) UpdateTask(taskName string, taskId int) (*Task, error) {

	if taskId <= 0 {
		return nil, custom_error.NewApiError(http.StatusUnprocessableEntity, fmt.Sprintf("wrong task id %d", taskId))
	}

	foundTask := s.DB.QueryRow("SELECT * FROM tasks where task_id = $1", taskId)
	if foundTask.Scan() == sql.ErrNoRows {
		return nil, custom_error.NewApiError(http.StatusNotFound, fmt.Sprintf("task with id %d doesn't exist", taskId))
	}

	_, err := s.DB.Query("UPDATE tasks SET task_name = $1 WHERE task_id = $2", taskName, taskId)
	if err != nil {
		return nil, err
	}

	return &Task{taskId, taskName}, nil
}

func (s *TaskService) DeleteTask(taskId int) error {
	_, err := s.DB.Query("DELETE FROM tasks WHERE task_id = $1", taskId)
	if err != nil {
		return err
	}

	return nil
}
