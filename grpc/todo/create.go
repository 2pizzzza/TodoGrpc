package todo

import (
	"ToDoAppGrpc/internal/todo_v1"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func (s *ServerAPI) CreateTodo(
	ctx context.Context, req *todo_v1.CreateTodoRequest) (res *todo_v1.TodoResponse, err error) {

	return &todo_v1.TodoResponse{
		Todo: &todo_v1.Todo{
			Title:       req.Title,
			Description: req.Description,
			Completed:   req.Completed,
			CreatedAt:   timestamppb.New(time.Now()),
		},
	}, nil
}
