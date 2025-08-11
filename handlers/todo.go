package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo/database"
)

type Todo struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only PostMthod Allowed.", http.StatusMethodNotAllowed)
		return
	}
	// decoding the json request
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "Bad request ", http.StatusBadRequest)
		return
	}
	// inserting into sql database
	query := "INSERT INTO todos(title, completed) VALUES (?,?,?)"
	_, insertErr := database.DB.Exec(query, todo.Title, todo.Completed)
	if insertErr != nil {
		http.Error(w, "Internal service error..", http.StatusInternalServerError)
		return
	}
	fmt.Println("data has been inserted..")
}
