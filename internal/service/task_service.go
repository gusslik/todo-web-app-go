package service

import (
	"database/sql"
	"fmt"
)

type TaskService struct {
	DB *sql.DB
}

func NewTaskService(db *sql.DB) *TaskService {
	return &TaskService{DB: db}
}

func (s *TaskService) GetTasks() {
	fmt.Println("Task's been created successfully")
}
