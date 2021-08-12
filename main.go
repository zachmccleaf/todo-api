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

	r.HandleFunc("/api/todos", controllers.GetTodos).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/todos/{id}", controllers.GetTodo).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/todos", controllers.CreateTodos).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/todos/{id}", controllers.UpdateTodos).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/todos/{id}", controllers.DeleteTodos).Methods("DELETE", "OPTIONS")

	config := helper.GetConfiguration()
	log.Fatal(http.ListenAndServe(config.Port, r))

}
