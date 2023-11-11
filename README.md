# Go Todolist

## Description
A simple todolist made with Go, HTMX, Posgresql. 

Additional tools used: 
- golang-migrate
- Tailwind
- Docker
- gorilla/mux

It displays a list of todos ordered by:
  1) whether it is done or not
  2) created date



### Project Layout
`/db` : database connection, migrations, db variable exposed as `TodoDB`

`/todo`: `Todo` database model, CRUD operation on `Todo` model

`/routes`: registers backend paths for CRUD operation

`Makefile`: makefile for the shell commands

## How to run
```
// connect to db, create db, create table
$ make init
$ make createdb
$ make migrateup
// run go
$ go run main.go
```

## TODO:
- multiple users / user authentication
- edit todos
- reorder todos

## Image
![todo](/img/todo.png)
