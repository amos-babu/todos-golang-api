package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/amos-babu/todo-app/database"
	"github.com/amos-babu/todo-app/handlers"
	"github.com/amos-babu/todo-app/repository"
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

	// fmt.Println(dbUsername, dbName, dbPassword)

	db, err := database.ConnectToDB(dbUsername, dbName, dbPassword)
	if err != nil {
		log.Fatal("Error Connecting to the database: ", err)
	}

	defer db.Close()

	fmt.Println("✅ Connected to the database successfully")

	r := mux.NewRouter()

	todoRepo := &repository.TodoRepository{DB: db}
	todoHandler := &handlers.TodoHandler{Repo: todoRepo}

	//Route Health Check
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status": "ok"}`))
	}).Methods("GET")

	//Todo Routes
	r.HandleFunc("/todos", todoHandler.HandleGetAllTodos).Methods("GET")
	r.HandleFunc("/todos/{id}", todoHandler.HandleGetTodo).Methods("GET")
	r.HandleFunc("/todos/{id}", todoHandler.HandleUpdateTodo).Methods("PUT")
	r.HandleFunc("/todos", todoHandler.HandleCreateTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", todoHandler.HandleDeleteTodo).Methods("DELETE")

	server := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}
