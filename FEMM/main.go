package main

import (
	"net/http"
	"fmt"
)
func handlePing (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}
func main() {
	// create a new server
	server := http.NewServeMux()
	server.HandleFunc("/ping", handlePing)

	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	// start the server
	fmt.Println("Server started at http://localhost:8080")
	defer fmt.Println("Server stopped")
	err := http.ListenAndServe(":8080", server)
	if err != nil {
		fmt.Println(err)
	}
}