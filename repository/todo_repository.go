package repository

import (
	"database/sql"

	"github.com/amos-babu/todo-app/models"
)

type TodoRepository struct {
	DB *sql.DB
}

func (r *TodoRepository) GetAll() ([]models.Todo, error) {
	return nil, nil
}

func (r *TodoRepository) CreateTodo() (models.Todo, error) {
	return models.Todo{}, nil
}

func (r *TodoRepository) GetTodoById(id int) (models.Todo, error) {
	return models.Todo{}, nil
}

func (r *TodoRepository) UpdateTodo(id int) (models.Todo, error) {
	return models.Todo{}, nil
}

func (r *TodoRepository) DeleteTodo(id int) error {
	return nil
}
