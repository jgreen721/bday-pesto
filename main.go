package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var (
	todos = []Todo{
		{Id: 1, Title: "todo-1", Description: "here is your first todo"},
		{Id: 2, Title: "todo-2", Description: "here is your second todo"},
		{Id: 3, Title: "todo-3", Description: "here is your third todo"},
	}
)

func main() {

	http.HandleFunc("/", welcomeGoApi)
	http.HandleFunc("/todos", getTodos)
	http.HandleFunc("/todo/{id}", getTodo)
	// http.HandleFunc("/", welcomeGoApi)

	fmt.Println("Listening in on port 3456")
	http.ListenAndServe(":3456", nil)
}

func welcomeGoApi(w http.ResponseWriter, r *http.Request) {

	fmt.Println("--GO API was pinged...")
	fmt.Fprintf(w, "Welcome to your GO API")

}

func getTodos(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("%s\n", r.URL)
	json.NewEncoder(w).Encode(todos)
}

func getTodo(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("%s\n", r.Method)
	json.NewEncoder(w).Encode(todos)
}
