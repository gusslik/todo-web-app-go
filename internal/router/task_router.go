package router

import (
	"database/sql"
	"todo-web-app-go/internal/handler"

	"github.com/gorilla/mux"
)

type TaskRouter struct {
	DB *sql.DB
}

func NewTaskRouter(db *sql.DB) RouteModule {
	return &TaskRouter{DB: db}
}

func (t *TaskRouter) RegisterRoutes(r *mux.Router) {
	taskHandler := handler.NewTaskHandler(t.DB)

	userRouter := r.PathPrefix("/api/tasks").Subrouter()
	userRouter.HandleFunc("", taskHandler.GetTasks).Methods("GET")
	userRouter.HandleFunc("", taskHandler.CreateTask).Methods("POST")
}
