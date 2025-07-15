package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todo/database"

	"github.com/gorilla/mux"
)

type updatetodo struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Only Put Method Allowed..", http.StatusMethodNotAllowed)
		return
	}
	// task_id_str := r.URL.Query().Get("task_ID")
	// task_id, _ := strconv.Atoi(task_id_str)
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
	var data updatetodo
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println("error decoding the data", err)
		return
	}
	// after getting the task_id we then query with task_id
	query := "UPDATE todos SET title=? and completed=? WHERE id=?"
	_, insert_err := database.DB.Exec(query, data.Title, data.Completed, TaskID)
	if insert_err != nil {
		http.Error(w, "insert data error", http.StatusInternalServerError)
		return
	}
	fmt.Println("Data updated succesfully..")

}
