package main

import (
	"./server"
	"net/http"
)

var (
	Router *http.ServeMux
)

func main() {
	Router = http.NewServeMux()
	server.RegisterRouter(Router)
	http.ListenAndServe(":8080", Router)
}
