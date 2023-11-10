package routes

import (
	"fmt"
	"go-todo/todo"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func sendTodos(w http.ResponseWriter) {
	todos, err := todo.GetAllTodos()
	if err != nil {
		fmt.Printf("Error:%s", err)
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err = tmpl.ExecuteTemplate(w, "todos", todos)
	if err != nil {
		fmt.Printf("Error:%s", err)
	}
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("Error:%s", err)
	}
	err = todo.CreateTodo(r.FormValue("todoStr"))
	if err != nil {
		fmt.Printf("Error:%s", err)
	}
	sendTodos(w)
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := todo.GetAllTodos()
	if err != nil {
		fmt.Printf("Error:%s", err)
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err = tmpl.Execute(w, todos)
	if err != nil {
		fmt.Println("Could not execute template", err)
	}

}

func updateDone(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		fmt.Println("Could not parse id", err)
	}
	err = todo.UpdateDone(id)
	if err != nil {
		fmt.Println("Could not update todo", err)
	}
	sendTodos(w)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		fmt.Println("C2ould not parse id", err)
	}
	err = todo.Delete(id)
	if err != nil {
		fmt.Println("Could not delete", err)
	}
	sendTodos(w)
}

func Run() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", getTodos)
	mux.HandleFunc("/todo/{id}", updateDone).Methods("PUT")
	mux.HandleFunc("/todo/{id}", deleteTodo).Methods("DELETE")
	mux.HandleFunc("/create", createTodo).Methods("POST")
	log.Fatal(http.ListenAndServe(":5001", mux))

}
