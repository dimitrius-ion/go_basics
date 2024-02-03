package main

import (
	"net/http"
	"fmt"
	"text/template"
	"github.com/dimitrius-ion/go_basics/femm/data"
	"github.com/dimitrius-ion/go_basics/femm/api"
	"github.com/gin-gonic/gin"
)
func handlePing (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func handleTemplate (w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("./templates/index.tmpl")
	if err != nil {
		fmt.Println(err)
	}else{
		html.Execute(w, data.GetAll())
	}
}

func main() {
	r:= gin.Default()
	r. GET("/ping", func(c *gin. Context) {
		c.JSON (200, gin. H{
			"message": "pong",
		})
	})
	go r.Run() 
	// create a new server
	server := http.NewServeMux()
	server.HandleFunc("/ping", handlePing)
	
	//template
	server.HandleFunc("/template", handleTemplate)
	
	//api
	// https://gin-gonic.com/ < like express for node
	server.HandleFunc("/api/data", api.Get)
	server.HandleFunc("/api/data/add", api.Post)

	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	// start the server
	fmt.Println("Server started at http://localhost:8080")
	defer fmt.Println("Server stopped")
	err := http.ListenAndServe(":8000", server)
	if err != nil {
		fmt.Println(err)
	}

	
}