package todo

import (
	"ToDoAppGrpc/internal/lib/logger/sl"
	"context"
	"fmt"
	"log/slog"
)

func (a *Todo) RemoveTodo(ctx context.Context, id int) (message string, err error) {
	const op = "service.todo.RemoveTodo"

	log := a.log.With(
		slog.String("op: ", op),
	)

	message, err = a.todoService.Remove(ctx, id)
	if err != nil {
		log.Error("Dont get todo by id", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, err)
	}

	return message, nil
}
