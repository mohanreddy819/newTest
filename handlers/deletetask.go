package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo/database"

	"github.com/gorilla/mux"
)

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["task_id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec("DELETE FROM todo WHERE id = ?", taskID)
	if err != nil {
		http.Error(w, "Database delete failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Todo deleted"})
}
