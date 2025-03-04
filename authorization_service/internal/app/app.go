package app

import (
	grpcapp "authorization_service/internal/app/grpc"
	"authorization_service/internal/services/auth"
	postgresql "authorization_service/internal/storage/postgreSQL"
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type App struct {
	GRPCServer *grpcapp.App
	Storage    postgresql.Repository
}

func New(
	ctx context.Context,
	log *logrus.Logger,
	grpcPort int,
	dsn string,
	tokenTTL time.Duration,
) *App {
	storage, err := postgresql.NewStorage(ctx, dsn, log)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, tokenTTL)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
		Storage:    storage,
	}
}
