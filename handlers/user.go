package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo/database"
)

type signupform struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func SignUpUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only Post Method allowed", http.StatusMethodNotAllowed)
		return
	}

	var input signupform

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil || input.Username == "" || input.Password == "" {
		http.Error(w, "Data not in proper format", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO users (username,password,email) VALUES(?,?,?)"

	_, insertErr := database.DB.Exec(query, input.Username, input.Password, input.Email)
	if insertErr != nil {
		http.Error(w, "problem during insertion of data ", http.StatusInternalServerError)
		return
	}

	fmt.Printf("user is created is %s", input.Username)

}
