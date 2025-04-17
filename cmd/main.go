package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"todo-web-app-go/internal/db"
	"todo-web-app-go/internal/router"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	var (
		host     = os.Getenv("DBHOST")
		port     = os.Getenv("DBPORT")
		user     = os.Getenv("DBUSER")
		password = os.Getenv("DBPASSWORD")
		dbname   = os.Getenv("DBNAME")
	)

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	taskDB := db.OpenDBConnection("postgres", dataSourceName)
	defer taskDB.Close()

	routeModules := []router.RouteModule{router.NewTaskRouter(taskDB)}
	router := router.Setup(routeModules)

	fmt.Println("Server is starting...")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
