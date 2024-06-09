package postgres

import (
	"ToDoAppGrpc/internal/domain/models"
	"ToDoAppGrpc/internal/storage"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

func (s *DB) Save(ctx context.Context, temp models.Model) (todo models.Model, err error) {
	const op = "storage/postgres.CreateToDo"

	var id int
	err = s.Db.QueryRow(
		"INSERT INTO todo (title, description, completed) VALUES ($1, $2, $3) RETURNING id\n",
		temp.Title, temp.Description, temp.Completed).Scan(&id)
	if err != nil {
		log.Printf("failed to SQL create: %v", err)
		return models.Model{}, err
	}

	todo, err = s.GetById(ctx, int(id))

	if err != nil {
		return models.Model{}, fmt.Errorf("%s, %w", op, err)
	}

	return todo, nil
}

func (s *DB) GetById(ctx context.Context, id int) (models.Model, error) {
	const op = "storage/postgres.GetById"
	stmt, err := s.Db.Prepare("SELECT * FROM todo WHERE id = $1")

	defer stmt.Close()

	if err != nil {
		return models.Model{}, fmt.Errorf("%s: %w", op, err)
	}

	var tempToDo = models.Model{}

	var (
		title       string
		description string
		completed   bool
		created_at  time.Time
	)

	err = stmt.QueryRow(id).Scan(&id, &title, &description, &completed, &created_at)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Model{}, storage.ErrToDoNotFound
		}
		return models.Model{}, fmt.Errorf("%s: %w", op, err)
	}
	tempToDo.ID = id
	tempToDo.Title = title
	tempToDo.Description = description
	tempToDo.Completed = completed
	tempToDo.CreatedAt = created_at

	return tempToDo, nil
}

func (s *DB) Remove(ctx context.Context, id int) (message string, err error) {
	const op = "postgres.todo.RemoveTodo"

	stmt, err := s.Db.Prepare("DELETE FROM todo WHERE id = $1")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	if rowsAffected == 0 {
		return "", storage.ErrToDoNotFound
	}

	return fmt.Sprintf("Successfully deleted todo id: %d", id), nil
}

func (s *DB) Change(ctx context.Context, reqTodo models.Model) (models.Model, error) {
	const op = "storage.postgres.todo.Change"

	stmt, err := s.Db.Prepare("UPDATE todo SET title = $2, description = $3, completed = $4 WHERE id = $1")
	if err != nil {
		return models.Model{}, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(reqTodo.ID, reqTodo.Title, reqTodo.Description, reqTodo.Completed)

	if err != nil {
		return models.Model{}, fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return models.Model{}, fmt.Errorf("%s: %w", op, err)
	}

	if rowsAffected == 0 {
		return models.Model{}, storage.ErrToDoNotFound
	}

	return reqTodo, nil
}
