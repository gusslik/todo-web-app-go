package handler

import (
	"database/sql"
	"encoding/json"
	"log"
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

	tasks := service.GetTasks()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(tasks)
	if err != nil {
		log.Fatal(err)
	}

}
