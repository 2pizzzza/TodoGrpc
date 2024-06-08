package todo

import (
	"ToDoAppGrpc/internal/todo_v1"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func (s *ServerAPI) CompleteTodo(
	ctx context.Context, req *todo_v1.IdRequest,
) (res *todo_v1.TodoResponse, err error) {

	return &todo_v1.TodoResponse{
		Todo: &todo_v1.Todo{
			Title:       "dg",
			Description: "req.Description",
			Completed:   true,
			CreatedAt:   timestamppb.New(time.Now()),
		},
	}, nil
}
