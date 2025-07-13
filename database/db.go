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

	CreateTables()

}

func CreateTables() {
	UserTable := `CREATE TABLE IF NOT EXISTS users(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE);`

	TodoTable := `CREATE TABLE IF NOT EXISTS todos(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_id INTEGER NOT NULL,
	title TEXT NOT NULL,
	completed BOOLEAN NOT NULL DEFAULT 0,
	FOREIGN KEY(user_id) REFERENCES users(id));`

	// command for creating the tables
	_, Err := DB.Exec(UserTable)
	if Err != nil {
		fmt.Println(Err)
		panic("user table not created error..")
	}
	_, TodoTableErr := DB.Exec(TodoTable)
	if TodoTableErr != nil {
		fmt.Println(TodoTableErr)
		panic("todos table not created error..")
	}

}
