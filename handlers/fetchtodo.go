package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo/database"

	"github.com/gorilla/mux"
)

type todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func GetTodoByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["task_id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var t todo
	err = database.DB.QueryRow("SELECT id, title, completed FROM todo WHERE id = ?", taskID).
		Scan(&t.ID, &t.Title, &t.Completed)
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(t)
}
