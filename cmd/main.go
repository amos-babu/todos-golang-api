package main

import (
	"log"
	"net/http"
	"time"

	"github.com/amos-babu/todo-app/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HandleGetAllTodos)
	r.HandleFunc("/todo/{id}", handlers.HandleGetTodo)
	r.HandleFunc("/todo/{id}", handlers.HandleUpdateTodo).Methods("PUT")
	r.HandleFunc("/todo", handlers.HandleCreateTodo).Methods("POST")
	r.HandleFunc("/todo/{id}", handlers.HandleDeleteTodo).Methods("DELETE")

	server := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
