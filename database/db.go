package database

import (
	"database/sql"
	"fmt"

	_ "github.com/glebarez/sqlite"
	// _ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	// DB, err = sql.Open("sqlite", "./todo.db")
	DB, err = sql.Open("sqlite", "./todo.db")
	if err != nil {
		fmt.Println("error while connecting to DB ", err)
	}

	fmt.Println("DB connected successfully..")

	// CreateTables()

}

// no need of this as we will use flyway to create the tables
// func CreateTables() {
// TodoTable := `CREATE TABLE IF NOT EXISTS todos(
// id INTEGER PRIMARY KEY AUTOINCREMENT,
// user_id INTEGER NOT NULL,
// title TEXT NOT NULL,
// completed BOOLEAN NOT NULL DEFAULT 0,
// FOREIGN KEY(user_id) REFERENCES users(id));`
// _, TodoTableErr := DB.Exec(TodoTable)
// if TodoTableErr != nil {
// 	fmt.Println(TodoTableErr)
// 	panic("todos table not created error..")
// }

// }
