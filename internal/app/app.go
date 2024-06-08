package app

import (
	"ToDoAppGrpc/internal/app/grpc"
	"ToDoAppGrpc/internal/config"
	"ToDoAppGrpc/internal/lib/logger/sl"
	"ToDoAppGrpc/internal/storage/postgres"
	"log/slog"
)

type App struct {
	GRPCSrv *grpc.App
}

func New(log *slog.Logger, cfg *config.Config) *App {

	db, err := postgres.New(cfg)
	if err != nil {
		log.Error("Failed connect db err: %s", sl.Err(err))
	}

	grpcApp := grpc.New(log, db.Db, cfg.GrpcPort, cfg.GrpcHost)

	return &App{
		GRPCSrv: grpcApp,
	}
}
