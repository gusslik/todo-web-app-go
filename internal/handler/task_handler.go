package handler

import (
	"database/sql"
	"net/http"
	"todo-web-app-go/internal/service"
)

type TaskHandler struct {
	DB *sql.DB
}

func NewTaskHandler(db *sql.DB) *TaskHandler {
	return &TaskHandler{DB: db}
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	service := service.NewTaskService(h.DB)

	service.GetTasks()
}
