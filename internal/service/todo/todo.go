package todo

import (
	"ToDoAppGrpc/internal/domain/models"
	"golang.org/x/net/context"
	"log/slog"
)

type Todo struct {
	log        *slog.Logger
	todoCreate CreateTodo
}

type CreateTodo interface {
	SaveToDo(ctx context.Context, mod models.Model) (todo models.Model, err error)
}

func New(
	log *slog.Logger,
	todoCreate CreateTodo,
) *Todo {

	return &Todo{
		log:        log,
		todoCreate: todoCreate,
	}
}
