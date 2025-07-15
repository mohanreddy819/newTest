package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todo/database"

	"github.com/gorilla/mux"
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

	// ID := r.URL.Query().Get("user_id")
	// userID, er := strconv.Atoi(ID)
	// if er != nil {
	// 	http.Error(w, "bad request id", http.StatusBadRequest)
	// 	return
	// }
	// writing mux logic
	vars := mux.Vars(r)
	UserID, strErr := strconv.Atoi(vars["user_id"])
	if strErr != nil {
		http.Error(w, "there was an error parsing the data try again..", http.StatusBadRequest)
		return
	}
	// check for user
	// TaskID:=strconv.Atoi(vars["task_id"])
	query1 := "SELECT EXISTS(SELECT 1 FROM users WHERE id=?)"
	var checkUser bool
	checkErr := database.DB.QueryRow(query1, UserID).Scan(&checkUser)
	if checkErr != nil {
		http.Error(w, "there was an error checking userIDtry again..", http.StatusBadRequest)
		return
	}
	fmt.Println(checkUser)
	if !checkUser {
		http.Error(w, "the user does not exists...", http.StatusNotFound)
		return
	}
	query := "SELECT * FROM todos WHERE user_id=?"
	rows, err := database.DB.Query(query, UserID)
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
	fmt.Printf(" User %v has been validated...\n", UserID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&todoData)
}
