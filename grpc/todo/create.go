package todo

import (
	"ToDoAppGrpc/internal/domain/models"
	"ToDoAppGrpc/internal/todo_v1"
	"context"
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

func (s *ServerAPI) CreateTodo(
	ctx context.Context, req *todo_v1.CreateTodoRequest) (*todo_v1.TodoResponse, error) {
	v, err := protovalidate.New()
	if err != nil {
		log.Fatalln("error protovalidate", err)
	}

	if err := v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	tempTodo := models.Model{
		Title:       req.Title,
		Description: req.Description,
		Completed:   req.Completed,
	}
	if s == nil {
		log.Fatalf("ServerAPI is nil")
	}
	if s.todo == nil {
		log.Fatalf("Todo interface is nil")
	}

	log.Printf("Creating todo with title: %s, description: %s", tempTodo.Title, tempTodo.Description)
	todo, err := s.todo.Create(ctx, tempTodo.Title, tempTodo.Description, tempTodo.Completed)
	log.Printf("ya sdez")
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &todo_v1.TodoResponse{
		Todo: &todo_v1.Todo{
			Id:          int64(todo.ID),
			Title:       todo.Title,
			Description: todo.Description,
			Completed:   todo.Completed,
			CreatedAt:   timestamppb.New(todo.CreatedAt),
		},
	}, nil
}
