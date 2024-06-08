package todo

import (
	"ToDoAppGrpc/internal/todo_v1"
	"golang.org/x/net/context"
	"log/slog"
)

type Todo struct {
	log        *slog.Logger
	todoCreate TodoCreate
}

type TodoCreate interface {
	CreateToDo(ctx context.Context, req todo_v1.TodoResponse)
}
