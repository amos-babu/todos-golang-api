package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/amos-babu/todo-app/database"
	"github.com/amos-babu/todo-app/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading .env files: ", err)
	}

	dbUsername := os.Getenv("DATABASE_USERNAME")
	dbName := os.Getenv("DATABASE_NAME")
	dbPassword := os.Getenv("DATABASE_PASSWORD")

	db, err := database.ConnectToDB(dbUsername, dbName, dbPassword)
	if err != nil {
		log.Fatal("Error Connecting to the database: ", err)
	}

	defer db.Close()

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
