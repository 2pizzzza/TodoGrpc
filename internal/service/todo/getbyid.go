package todo

import (
	"ToDoAppGrpc/internal/domain/models"
	"ToDoAppGrpc/internal/lib/logger/sl"
	"context"
	"fmt"
	"log/slog"
)

func (a *Todo) GetTodoById(ctx context.Context, id int) (todo models.Model, err error) {
	const op = "service.todo.GetToDoById"

	log := a.log.With(
		slog.String("op: ", op),
	)

	todo, err = a.todoService.GetById(ctx, id)

	if err != nil {
		log.Error("Dont get todo by id", sl.Err(err))

		return models.Model{}, fmt.Errorf("%s: %w", op, err)
	}

	return todo, nil
}
