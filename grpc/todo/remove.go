package todo

import (
	"ToDoAppGrpc/internal/todo_v1"
	"context"
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *ServerAPI) RemoveTodo(
	ctx context.Context, req *todo_v1.IdRequest,
) (res *todo_v1.InfoResponse, err error) {
	v, err := protovalidate.New()
	if err != nil {
		log.Fatalln("error protovalidate", err)
	}

	if err := v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	message, err := s.todo.RemoveTodo(ctx, int(req.Id))

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &todo_v1.InfoResponse{
		Message: message,
	}, nil
}
