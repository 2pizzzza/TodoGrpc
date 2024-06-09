package todo

import (
	"ToDoAppGrpc/internal/domain/models"
	"ToDoAppGrpc/internal/lib/logger/sl"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"time"
)

func (a *Todo) Create(ctx context.Context, title, description string, complete bool) (res models.Model, err error) {

	const op = "todo.Create"

	log := a.log.With(
		slog.String("op: ", op),
	)

	log.Info("create todo")

	tempTodo := models.Model{
		Title:       title,
		Description: description,
		Completed:   complete,
		CreatedAt:   timestamppb.New(time.Now()).AsTime(),
	}

	todo, err := a.todoService.Save(ctx, tempTodo)
	if err != nil {
		log.Error("failed create todo", sl.Err(err))

		return models.Model{}, fmt.Errorf("%s: %w", op, err)
	}
	log.Info("todo create")

	return todo, nil
}
