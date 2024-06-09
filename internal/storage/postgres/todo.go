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

func (s *DB) SaveToDo(ctx context.Context, temp models.Model) (todo models.Model, err error) {
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

func (s *DB) RemoveTodo(ctx context.Context, id int) (message string, err error) {
	const op = "postgres.todo.RemoveTodo"

	stmt, err := s.Db.Prepare("DELETE FROM todo WHERE id = $1")

	defer stmt.Close()

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", storage.ErrToDoNotFound
		}
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return fmt.Sprintf("Success deleted todo id: %d", id), nil
}
