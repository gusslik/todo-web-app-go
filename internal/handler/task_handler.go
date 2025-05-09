package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	custom_error "todo-web-app-go/internal/error"
	"todo-web-app-go/internal/service"
)

type TaskHandler struct {
	DB *sql.DB
}

type RequestData struct {
	Task_id   int    `json:"task_id"`
	Task_name string `json:"task_name"`
}

func NewTaskHandler(db *sql.DB) *TaskHandler {
	return &TaskHandler{DB: db}
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	service := service.NewTaskService(h.DB)

	tasks, err := service.GetTasks()
	if err != nil {
		custom_error.ErrorHandler(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	service := service.NewTaskService(h.DB)

	var data RequestData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		custom_error.ErrorHandler(w, err)
		return
	}

	task, err := service.CreateTask(data.Task_name)
	if err != nil {
		custom_error.ErrorHandler(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	service := service.NewTaskService(h.DB)
	var data RequestData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		custom_error.ErrorHandler(w, err)
		return
	}
	task, err := service.UpdateTask(data.Task_name, data.Task_id)
	if err != nil {
		custom_error.ErrorHandler(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	service := service.NewTaskService(h.DB)
	var data RequestData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		custom_error.ErrorHandler(w, err)
		return
	}
	err = service.DeleteTask(data.Task_id)
	if err != nil {
		custom_error.ErrorHandler(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
