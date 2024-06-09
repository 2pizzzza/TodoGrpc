package todo

import (
	"ToDoAppGrpc/internal/domain/models"
	"ToDoAppGrpc/internal/todo_v1"
	"context"
	"google.golang.org/grpc"
)

type Todo interface {
	Create(ctx context.Context, title, description string, complete bool) (res models.Model, err error)
	GetTodoById(ctx context.Context, id int) (todo models.Model, err error)
	RemoveTodo(ctx context.Context, id int) (message string, err error)
	ChangeTodo(ctx context.Context, reqTodo models.Model) (models.Model, error)
	GetAllTodo(ctx context.Context) (todos []*models.Model, err error)
}

type ServerAPI struct {
	todo_v1.UnimplementedTodoV1Server
	todo Todo
}

func Register(gRPC *grpc.Server, todo Todo) {
	todo_v1.RegisterTodoV1Server(gRPC, &ServerAPI{todo: todo})
}
