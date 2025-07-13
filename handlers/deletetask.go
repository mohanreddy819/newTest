package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"todo/database"
)

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Only Delete Method Allowed", http.StatusMethodNotAllowed)
		return
	}
	task_id_string := r.URL.Query().Get("task_id")

	task_id, errConv := strconv.Atoi(task_id_string)
	if errConv != nil {
		http.Error(w, "Bad ID request", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM todos WHERE id=?"
	_, sqErr := database.DB.Exec(query, task_id)
	if sqErr != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		fmt.Println(sqErr)
		return
	}
	fmt.Println("task delete success..")
}
