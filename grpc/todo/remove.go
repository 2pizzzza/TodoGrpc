package todo

import (
	"ToDoAppGrpc/internal/todo_v1"
	"context"
)

func (s *ServerAPI) RemoveTodo(
	ctx context.Context, req *todo_v1.IdRequest,
) (res *todo_v1.InfoResponse, err error) {

	return &todo_v1.InfoResponse{
		Message: "Success deleted todo",
	}, nil
}
