package main

import (
	"net/http"
	"todo-web-app-go/internal/router"
)

func main() {
	http.ListenAndServe("localhost:8000", router.Setup())
}
