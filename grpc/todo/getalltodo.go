package todo

import (
	"ToDoAppGrpc/internal/todo_v1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ServerAPI) GetAllTodo(ctx context.Context, in *emptypb.Empty) (*todo_v1.TodosResponse, error) {

	todos := []*todo_v1.Todo{}

	tempTodo, err := s.todo.GetAllTodo(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	for _, m := range tempTodo {
		todos = append(todos, &todo_v1.Todo{
			Id:          int64(m.ID),
			Title:       m.Title,
			Description: m.Description,
			Completed:   m.Completed,
			CreatedAt:   timestamppb.New(m.CreatedAt),
		})
	}

	return &todo_v1.TodosResponse{Todo: todos}, nil
}
