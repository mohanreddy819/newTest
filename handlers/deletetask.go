package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"todo/database"

	"github.com/gorilla/mux"
)

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Only Delete Method Allowed", http.StatusMethodNotAllowed)
		return
	}
	// task_id_string := r.URL.Query().Get("task_id")
	vars := mux.Vars(r)
	UserID, strErr := strconv.Atoi(vars["user_id"])
	if strErr != nil {
		http.Error(w, "there was an error parsing the data try again..", http.StatusBadRequest)
		return
	}
	TaskID, strTaskErr := strconv.Atoi(vars["task_id"])
	if strTaskErr != nil {
		http.Error(w, "there was an error parsing the data try again..", http.StatusBadRequest)
		return
	}
	checkUser := "SELECT EXISTS(SELECT 1 FROM users WHERE id=?)"
	var conf bool
	confErr := database.DB.QueryRow(checkUser, UserID).Scan(&conf)
	if confErr != nil {
		http.Error(w, "The database error..", http.StatusInternalServerError)
		return
	}
	if !conf {
		http.Error(w, "The does not exists...", http.StatusNotFound)
		return
	}

	// task_id, errConv := strconv.Atoi(task_id_string)
	// if errConv != nil {
	// 	http.Error(w, "Bad ID request", http.StatusBadRequest)
	// 	return
	// }

	query := "DELETE FROM todos WHERE id=?"
	_, sqErr := database.DB.Exec(query, TaskID)
	if sqErr != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		fmt.Println(sqErr)
		return
	}
	fmt.Println("task delete success..")
}
