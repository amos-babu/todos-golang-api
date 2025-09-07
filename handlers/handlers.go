package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/amos-babu/todo-app/models"
	"github.com/amos-babu/todo-app/repository"
	"github.com/gorilla/mux"
)

type TodoHandler struct {
	Repo *repository.TodoRepository
}

func (h *TodoHandler) HandleGetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.Repo.GetAll()
	if err != nil {
		http.Error(w, "Failed to fetch todos: ", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

func (h *TodoHandler) HandleCreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Failed to decode todos", http.StatusInternalServerError)
		return
	}

	if err := h.Repo.CreateTodo(&todo); err != nil {
		http.Error(w, "Error creating todos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"success": "Created successfully",
		"todo":    todo,
	})
}

func (h *TodoHandler) HandleGetTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID"})
		return
	}

	todo, err := h.Repo.GetTodoById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) HandleUpdateTodo(w http.ResponseWriter, r *http.Request) {
	var updatedTodo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
		http.Error(w, "Failed to decode body", http.StatusInternalServerError)
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
	// if err != nil {
	//     http.Error(w, "Invalid ID", http.StatusBadRequest)
	//     return
	// }

	if err := h.Repo.UpdateTodo(&updatedTodo, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"success": "Todo Updated successfully"})
}

func (h *TodoHandler) HandleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID"})
		return
	}

	if err := h.Repo.DeleteTodo(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"success": "Todo Deleted Successfully"})
}
