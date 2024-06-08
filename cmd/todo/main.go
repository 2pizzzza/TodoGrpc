package main

import (
	"ToDoAppGrpc/internal/config"
	"ToDoAppGrpc/internal/lib/logger/sl"
	"ToDoAppGrpc/internal/storage/postgres"
	"ToDoAppGrpc/internal/todo_v1"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
	"os"
	"strconv"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

type server struct {
	todo_v1.UnimplementedUserV1Server
}

func (s *server) GetUserInfo(ctx context.Context, req *todo_v1.GetRequest) (res *todo_v1.GetResponse, err error) {
	name := "John"
	if req.GetId() == 7 {
		name = "Amit"
	}
	return &todo_v1.GetResponse{
			Info: &todo_v1.UserInfo{
				Id:      req.GetId(),
				Name:    name,
				IsHuman: true,
			},
		},
		nil
}

func main() {

	env, err := config.NewConfig()

	log := setupLogger(env.Env)
	log.Debug("debug logger are enable")

	if err != nil {
		log.Error("Failed load env", sl.Err(err))
	}

	db, err := postgres.New(env)
	if err != nil {
		log.Error("Failed connect db err: %s", sl.Err(err))
	}
	_ = db
	log.Info("port, host", fmt.Sprintf(":%d", env.GrpcPort))
	lis, err := net.Listen("tcp",
		fmt.Sprintf("%s:%d", env.GrpcHost, env.GrpcPort))

	if err != nil {
		log.Error("failed to listen", sl.Err(err))
	}
	_ = lis
	log.Info("Server is life", slog.String("env", env.Env))

	s := grpc.NewServer()
	reflection.Register(s)
	todo_v1.RegisterUserV1Server(s, &server{})

	log.Info(
		"Server config",
		slog.String("host", env.GrpcHost),
		slog.String("port", strconv.Itoa(env.GrpcPort)))

	if err = s.Serve(lis); err != nil {
		log.Error("Problem with start grpc server:", sl.Err(err))
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(

			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
