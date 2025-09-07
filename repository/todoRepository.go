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
	query := `SELECT id, name, description, createdAt FROM todos;`
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

	return todos, nil
}

func (r *TodoRepository) CreateTodo(t *models.Todo) error {
	query := `INSERT INTO todos (name, description) VALUES (?, ?);`
	result, err := r.DB.Exec(query, t.Name, t.Description)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	t.Id = int(id)

	return nil
}

func (r *TodoRepository) GetTodoById(id int) (models.Todo, error) {
	var t models.Todo
	query := `SELECT id, name, description, createdAt FROM todos WHERE id = ?;`
	err := r.DB.QueryRow(query, id).Scan(&t.Id, &t.Name, &t.Description, &t.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return t, fmt.Errorf("No todo found with that id")
		} else {
			return t, err
		}
	}

	return t, nil
}

func (r *TodoRepository) UpdateTodo(t *models.Todo, id int) error {
	query := `UPDATE todos SET name = ?, description = ? WHERE id = ?;`
	result, err := r.DB.Exec(query, t.Name, t.Description, id)
	if err != nil {
		return err
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowCount == 0 {
		return fmt.Errorf("No todo updated with that id")
	}

	return nil
}

func (r *TodoRepository) DeleteTodo(id int) error {
	query := `DELETE FROM todos WHERE id = ?;`
	result, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no todo found with id %d", id)
	}

	return nil
}
