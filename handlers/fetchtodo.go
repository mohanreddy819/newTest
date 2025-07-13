package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo/database"
)

type TodoData struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only Get Method Allowed", http.StatusMethodNotAllowed)
		return
	}

	ID := r.URL.Query().Get("user_id")
	userID, er := strconv.Atoi(ID)
	if er != nil {
		http.Error(w, "bad request id", http.StatusBadRequest)
		return
	}

	query := "SELECT * FROM todos WHERE user_id=?"
	rows, err := database.DB.Query(query, userID)
	if err != nil {
		http.Error(w, "Bad Request data", http.StatusBadRequest)
		return
	}
	var todoData []TodoData
	for rows.Next() {
		var todo TodoData
		err := rows.Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Completed)
		if err != nil {
			http.Error(w, "Error scanning todo", http.StatusInternalServerError)
			continue
		}
		todoData = append(todoData, todo)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&todoData)
}
