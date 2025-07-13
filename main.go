package main

import (
	"fmt"
	"net/http"
	"todo/database"
	"todo/handlers"
)

func main() {
	fmt.Println("starting...")

	database.ConnectDB()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("api running...")
	})

	http.HandleFunc("/signupuser", handlers.SignUpUser)
	http.HandleFunc("/login", handlers.LoginUser)
	http.HandleFunc("/createtodo", handlers.CreateTodo)
	http.HandleFunc("/fetchtodo", handlers.GetTodo)
	http.HandleFunc("/updatetask", handlers.UpdateTodo)
	http.HandleFunc("/deletetask", handlers.DeleteTask)

	fmt.Println("starting server...")
	http.ListenAndServe(":8080", nil)

}
