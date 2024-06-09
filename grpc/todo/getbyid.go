package todo

import (
	"ToDoAppGrpc/internal/todo_v1"
	"context"
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

func (s *ServerAPI) GetTodoById(
	ctx context.Context, req *todo_v1.IdRequest,
) (res *todo_v1.TodoResponse, err error) {
	v, err := protovalidate.New()
	if err != nil {
		log.Fatalln("error protovalidate", err)
	}

	if err := v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	tempTodo, err := s.todo.GetTodoById(ctx, int(req.Id))

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &todo_v1.TodoResponse{
		Todo: &todo_v1.Todo{
			Id:          int64(tempTodo.ID),
			Title:       tempTodo.Title,
			Description: tempTodo.Description,
			Completed:   tempTodo.Completed,
			CreatedAt:   timestamppb.New(tempTodo.CreatedAt),
		},
	}, nil
}
