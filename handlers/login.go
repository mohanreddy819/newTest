package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo/database"
)

type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only Post Method Aloowed", http.StatusMethodNotAllowed)
		return
	}

	// since i will be recieving an an json ecoded format i must decode it first
	// and then parse them seperatrly put them in an query and use Exec()

	var loginInput login
	err := json.NewDecoder(r.Body).Decode(&loginInput)
	if err != nil {
		http.Error(w, "bad request data ", http.StatusBadRequest)
		return
	}

	// queryToCheck:= "SELECT password FROM users WHERE username=?"

	/*Exec is only meant for insert , update, delete and fails for reading the data as it returns just sql type
	to read data we use queryrow for single row answer and query fro multiple row*/
	check := database.DB.QueryRow("SELECT password FROM users WHERE username=?", loginInput.Username)
	// if err1 != nil {
	// 	fmt.Println("error checking login cred..")
	// 	return
	// }

	/* queryrow also reutrns a pointer that particular row so we have use scan to actually retrieve the data
	like when use queryrow it reuturns a box to open that box and see the data we need to use scan.
	*/
	var pass string
	fetcherr := check.Scan(&pass)
	if fetcherr != nil {
		fmt.Println("data does not exsist")
		http.Error(w, "wrong data", http.StatusBadRequest)
		return
	}
	// now check returned password and user entered password
	if pass == loginInput.Password {
		fmt.Printf("User %s has been validated..", loginInput.Username)
		sql_id := database.DB.QueryRow("SELECT id FROM users WHERE username=?", loginInput.Username)
		var id int
		id_fetching := sql_id.Scan(&id)
		fmt.Println(id_fetching)
		fmt.Printf("User ID is %d\n", id)

	}
}
