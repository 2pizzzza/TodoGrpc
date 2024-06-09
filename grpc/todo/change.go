package todo

import (
	"ToDoAppGrpc/internal/domain/models"
	"ToDoAppGrpc/internal/todo_v1"
	"context"
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *ServerAPI) ChangeTodo(
	ctx context.Context, req *todo_v1.ChangeTodoRequest) (res *todo_v1.TodoResponse, err error) {

	v, err := protovalidate.New()
	if err != nil {
		log.Fatalln("error protovalidate", err)
	}

	if err := v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	todo, err := s.todo.ChangeTodo(ctx, models.Model{
		ID:          int(req.Id),
		Title:       req.Title,
		Description: req.Description,
		Completed:   req.Completed,
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &todo_v1.TodoResponse{
		Todo: &todo_v1.Todo{
			Id:          int64(todo.ID),
			Title:       todo.Title,
			Description: todo.Description,
			Completed:   todo.Completed,
		},
	}, nil
}
