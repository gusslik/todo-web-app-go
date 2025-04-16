package main

import (
	"fmt"
	"log"
	"net/http"
	"todo-web-app-go/internal/router"
)

func main() {
	fmt.Println("Server is starting...")
	log.Fatal(http.ListenAndServe("localhost:8000", router.Setup()))
}
