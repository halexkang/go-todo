package main

import (
	"go-todo/db"
	"go-todo/routes"
)

func main() {
	db.NewDatabase()
	routes.SetupAndRun()
}
