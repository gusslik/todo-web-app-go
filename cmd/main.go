package main

import (
	"fmt"
	"log"
	"net/http"
	"todo-web-app-go/internal/config"
	"todo-web-app-go/internal/db"
	"todo-web-app-go/internal/router"
)

func main() {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.Dbname)
	taskDB := db.OpenDBConnection("postgres", dataSourceName)
	defer taskDB.Close()

	routeModules := []router.RouteModule{router.NewTaskRouter(taskDB)}
	router := router.Setup(routeModules)

	fmt.Println("Server is starting...")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
