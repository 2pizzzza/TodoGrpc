package todo

import (
	"ToDoAppGrpc/internal/domain/models"
	"ToDoAppGrpc/internal/lib/logger/sl"
	"context"
	"fmt"
	"log/slog"
)

func (a *Todo) ChangeTodo(ctx context.Context, reqTodo models.Model) (models.Model, error) {
	const op = "service.todo.change.ChangeTodo"

	log := a.log.With(
		slog.String("op: ", op),
	)

	tempTodo, err := a.todoService.Change(ctx, reqTodo)
	if err != nil {
		log.Error("failed create todo", sl.Err(err))

		return models.Model{}, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("todo changed")

	return tempTodo, nil
}
