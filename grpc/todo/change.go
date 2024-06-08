package todo

import (
	"ToDoAppGrpc/internal/todo_v1"
	"context"
)

func (s *ServerAPI) ChangeTodo(
	ctx context.Context, req *todo_v1.ChangeTodoRequest) (res *todo_v1.InfoResponse, err error) {

	return &todo_v1.InfoResponse{
		Message: "Success change todo",
	}, nil
}
