package main

import (
	"ToDoAppGrpc/internal/app"
	"ToDoAppGrpc/internal/config"
	"ToDoAppGrpc/internal/lib/logger/sl"
	"log/slog"
	"os"
	"strconv"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {

	env, err := config.NewConfig()

	log := setupLogger(env.Env)
	log.Debug("debug logger are enable")

	if err != nil {
		log.Error("Failed load env", sl.Err(err))
	}

	log.Info(
		"Server config",
		slog.String("host", env.GrpcHost),
		slog.String("port", strconv.Itoa(env.GrpcPort)))

	application := app.New(log, env)

	go application.GRPCSrv.MustRun()

	stop := make(chan os.Signal, 1)

	sgnl := <-stop

	log.Info("stopping application", slog.String("signal", sgnl.String()))

	application.GRPCSrv.Stop()

	log.Info("application stopped")
}

//func (s *server) GetUserInfo(ctx context.Context, req *todo_v1.GetRequest) (res *todo_v1.GetResponse, err error) {
//	name := "John"
//	if req.GetId() == 7 {
//		name = "Amit"
//	}
//	return &todo_v1.GetResponse{
//			Info: &todo_v1.UserInfo{
//				Id:      req.GetId(),
//				Name:    name,
//				IsHuman: true,
//			},
//		},
//		nil
//}

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
