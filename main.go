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

	r.HandleFunc("/signupuser", handlers.SignUpUser).Methods("POST")
	r.HandleFunc("/login", handlers.LoginUser).Methods("POST")
	// http.HandleFunc("/createtodo", handlers.CreateTodo)
	r.HandleFunc("/fetchtodo/{user_id:[0-9]+}", handlers.GetTodo).Methods("GET")
	r.HandleFunc("/updatetask/{user_id:[0-9]+}/{task_id:[0-9]+}", handlers.UpdateTodo).Methods("PUT")
	r.HandleFunc("/deletetask/{user_id:[0-9]+}/{task_id:[0-9]+}", handlers.DeleteTask).Methods("DELETE")
	r.HandleFunc("/fetchmuxID/{user_id:[0-9]+}", handlers.FetchtasksByuser).Methods("POST")
	fmt.Println("starting server...")
	http.ListenAndServe(":8080", r)

}
