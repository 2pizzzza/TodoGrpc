package todo

import "ToDoAppGrpc/internal/todo_v1"

type ServerAPI struct {
	todo_v1.UnimplementedTodoV1Server
}
