// fetching todo based on user ID and and checking if the user exists and and then fetching the tasks

package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"todo/database"

	"github.com/gorilla/mux"
)

func FetchtasksByuser(w http.ResponseWriter, r *http.Request) {
	// get userID
	vars := mux.Vars(r)
	UserID, _ := strconv.Atoi(vars["user_id"])
	// TaskID, _ := strconv.Atoi(vars["task_id"])

	// check forr user_id existstence.
	query1 := "SELECT EXISTS (SELECT 1 FROM users WHERE id = ?);"
	// userConf := database.DB.QueryRow(query1, UserID)
	var exists bool
	err := database.DB.QueryRow(query1, UserID).Scan(&exists)
	// var exists bool
	// err := userConf.Scan(&exists)
	if err != nil {
		http.Error(w, "database error..", http.StatusBadRequest)
		return
	}
	if !exists {
		http.Error(w, "user does not exists or db error ", http.StatusBadRequest)
		return
	}

	fmt.Println("success", exists)

}
