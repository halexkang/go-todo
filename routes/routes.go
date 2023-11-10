package routes

import (
	"fmt"
	"go-todo/db"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func sendTodos(w http.ResponseWriter) {
	todos, err := db.GetAllTodos()
	if err != nil {
		fmt.Errorf("Error:%s", &err)
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err = tmpl.ExecuteTemplate(w, "Todos", todos)
	if err != nil {
		fmt.Errorf("Error:%s", &err)
	}
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("Error:%s", err)
	}
	err = db.CreateTodo(r.FormValue("todo"))
	if err != nil {
		fmt.Printf("Error:%s", err)
	}
	sendTodos(w)
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := db.GetAllTodos()
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
	err = db.UpdateDone(id)
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
	err = db.Delete(id)
	if err != nil {
		fmt.Println("Could not delete", err)
	}
	sendTodos(w)
}

func SetupAndRun() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", getTodos)
	mux.HandleFunc("/todo/{id}", updateDone).Methods("PUT")
	mux.HandleFunc("/todo/{id}", deleteTodo).Methods("DELETE")
	mux.HandleFunc("/create", createTodo).Methods("POST")
	log.Fatal(http.ListenAndServe(":5001", mux))

}
