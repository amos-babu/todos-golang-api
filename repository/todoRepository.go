package repository

import (
	"database/sql"
	"fmt"

	"github.com/amos-babu/todo-app/models"
)

type TodoRepository struct {
	DB *sql.DB
}

func (r *TodoRepository) GetAll() ([]models.Todo, error) {
	query := `SELECT * FROM todos;`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var t models.Todo
		if err := rows.Scan(&t.Id, &t.Name, &t.Description, &t.CreatedAt); err != nil {
			return nil, err
		}

		todos = append(todos, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	fmt.Println(todos)

	return todos, nil
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
