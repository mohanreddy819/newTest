package main

import (
	"fmt"
	"net/http"
	"todo/database"

	"todo/handlers"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("starting...")

	database.ConnectDB()
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("api running...")
	})

	// http.HandleFunc("/createtodo", handlers.CreateTodo)
	r = mux.NewRouter()
	r.HandleFunc("/todos", handlers.CreateTodo).Methods("POST")
	// r.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
	r.HandleFunc("/todos/{task_id}", handlers.GetTodoByID).Methods("GET")
	r.HandleFunc("/todos/{task_id}", handlers.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/{task_id}", handlers.DeleteTodo).Methods("DELETE")

	fmt.Println("starting server...")
	http.ListenAndServe(":8080", r)

}
