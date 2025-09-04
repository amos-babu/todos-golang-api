package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/amos-babu/todo-app/models"
	"github.com/gorilla/mux"
)

var todos []models.Todo

func HandleGetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos = []models.Todo{
		{
			Id:          1,
			Name:        "Learn Golang",
			Description: "ResponseWriter",
			CreatedAt:   time.Now(),
		},
		{
			Id:          2,
			Name:        "Test the golang api",
			Description: "Second task",
			CreatedAt:   time.Now().Add(1 * time.Hour),
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

func HandleGetTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID"})
		return
	}

	for _, todo := range todos {
		if todo.Id == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(todo)
			fmt.Println(todo)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Todo With ID Not Found"})
}

func HandleCreateTodo(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Server Error"})
		return
	}

	var todo models.Todo
	if err := json.Unmarshal(reqBody, &todo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid Data"})
		return
	}

	todo.Id = len(todos) + 1
	todo.CreatedAt = time.Now()
	todos = append(todos, todo)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"success": "Created successfully",
		"todo":    todos,
	})
}

func HandleUpdateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Server Error"})
		return
	}

	var updatedTodo models.Todo
	if err := json.Unmarshal(reqBody, &updatedTodo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid Data"})
		return
	}
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID"})
		return
	}

	for i := range todos {
		if todos[i].Id == id {
			todos[i].Name = updatedTodo.Name
			todos[i].Description = updatedTodo.Description

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]any{
				"success": "Created successfully",
				"todo":    todos,
			})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Todo With ID Not Found"})
}

func HandleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID"})
		return
	}

	for i, todo := range todos {
		if todo.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]any{
				"success": "Created successfully",
				"todo":    todos,
			})
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Todo With ID Not Found"})
}
