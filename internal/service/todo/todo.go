package todo

import (
	"ToDoAppGrpc/internal/domain/models"
	"golang.org/x/net/context"
	"log/slog"
)

type Todo struct {
	log         *slog.Logger
	todoService T
}

type T interface {
	Save(ctx context.Context, mod models.Model) (todo models.Model, err error)
	GetById(ctx context.Context, id int) (models.Model, error)
	Remove(ctx context.Context, id int) (message string, err error)
}

func New(
	log *slog.Logger,
	todoService T,
) *Todo {

	return &Todo{
		log:         log,
		todoService: todoService,
	}
}
