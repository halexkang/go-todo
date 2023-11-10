package todo

import (
	"fmt"
	"go-todo/db"
)

// look https://go.dev/doc/database/querying for SQL query examples
type Todo struct {
	Id      uint64 `json:"id"`
	TodoStr string `json:"todoStr"`
	Done    bool   `json:"done"`
}

func CreateTodo(todo string) error {
	stmt := `insert into todos(todo, done) values($1, $2);`
	_, err := db.TodoDB.Query(stmt, todo, false)
	if err != nil {
		fmt.Println("hey man its broken")
	}
	return err

}

func GetAllTodos() ([]Todo, error) {
	stmt := `select id, todo, done from todos order by done, ts;`
	rows, err := db.TodoDB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []Todo{}
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.Id, &todo.TodoStr, &todo.Done); err != nil {
			return todos, err
		}
		todos = append(todos, todo)
	}
	if err = rows.Err(); err != nil {
		return todos, err
	}
	return todos, nil
}

func GetTodoById(id uint64) (Todo, error) {
	todo := Todo{}
	todo.Id = id
	stmt := `select todo, done from todos where id=$1;`
	var todoStr string
	var done bool
	if err := db.TodoDB.QueryRow(stmt, id).Scan(&todoStr, &done); err != nil {
		return todo, err
	}
	todo.TodoStr = todoStr
	todo.Done = done
	return todo, nil
}

func UpdateDone(id uint64) error {
	todo, err := GetTodoById(id)
	if err != nil {
		return err
	}
	stmt := `update todos set done=$2 where id=$1;`
	_, err = db.TodoDB.Query(stmt, id, !todo.Done)
	return err
}

func Delete(id uint64) error {
	stmt := `delete from todos where id=$1;`
	_, err := db.TodoDB.Exec(stmt, id)
	return err
}
