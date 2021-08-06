package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zachmccleaf/todo-api/controllers"
	"github.com/zachmccleaf/todo-api/helper"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/todos", controllers.GetTodos).Methods("GET")
	r.HandleFunc("/api/todos/{id}", controllers.GetTodo).Methods("GET")
	r.HandleFunc("/api/todos", controllers.CreateTodos).Methods("POST")
	r.HandleFunc("/api/todos/{id}", controllers.UpdateTodos).Methods("PUT")
	r.HandleFunc("/api/todos/{id}", controllers.UpdateTodos).Methods("DELETE")

	config := helper.GetConfiguration()
	log.Fatal(http.ListenAndServe(config.Port, r))

}
