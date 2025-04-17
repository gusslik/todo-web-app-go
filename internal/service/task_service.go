package service

import (
	"database/sql"
	"log"
)

type TaskService struct {
	DB *sql.DB
}

type Task struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func NewTaskService(db *sql.DB) *TaskService {
	return &TaskService{DB: db}
}

func (s *TaskService) GetTasks() []Task {
	rows, err := s.DB.Query("SELECT * FROM tasks")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var task Task
		err = rows.Scan(&task.Id, &task.Name)
		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, task)
	}

	return tasks
}
