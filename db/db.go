package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	TodoDB *sql.DB
)

func NewDatabase() {
	var err error
	TodoDB, err = sql.Open("postgres", "postgresql://root:postgres@localhost:5432/go-todo?sslmode=disable")
	if err != nil {
		fmt.Println("Could not connect to db", err)
	}
	err = TodoDB.Ping()
	if err != nil {
		fmt.Println("Could not ping db", err)
	}
	fmt.Println("connected to db")
}
