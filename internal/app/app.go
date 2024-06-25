package app

import (
	"log"
	"time"

	grpcapp "github.com/p3rfect05/grpc_service/internal/app/grpc"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *log.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	// TODO: init storage

	grpcServer := grpcapp.New(log, grpcPort)
	return &App{
		GRPCServer: grpcServer,
	}
}
