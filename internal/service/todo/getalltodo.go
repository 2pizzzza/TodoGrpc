package todo

import (
	"ToDoAppGrpc/internal/domain/models"
	"ToDoAppGrpc/internal/lib/logger/sl"
	"context"
	"fmt"
	"log/slog"
)

func (a *Todo) GetAllTodo(ctx context.Context) (todos []*models.Model, err error) {
	const op = "service.todo.GetAllToDo"

	log := a.log.With(
		slog.String("op: ", op),
	)

	todos, err = a.todoService.GetAll(ctx)
	if err != nil {
		log.Error("failed get all todos", sl.Err(err))

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return todos, nil
}
