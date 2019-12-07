package main

import (
	"./server"
	"fmt"
	"net/http"
)

var (
	Router *http.ServeMux
)

//main.go
func main() {
	Router = http.NewServeMux()
	server.RegisterRouter(Router)

	fmt.Println("Now the application expose port: http://localhost:8080")
	http.ListenAndServe(":8080", Router)
}
