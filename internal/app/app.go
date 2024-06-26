package app

import (
	"log"
	"time"

	grpcapp "github.com/p3rfect05/grpc_service/internal/app/grpc"
	"github.com/p3rfect05/grpc_service/internal/config"
	auth "github.com/p3rfect05/grpc_service/internal/services"
	postgres "github.com/p3rfect05/grpc_service/internal/storage"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *log.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
	DSN config.DatabaseURLConfig,
) *App {
	var DBConnectRetries = 3

	log.Println("connecting to the database...")
	storage, err := postgres.New(DSN)
	for ; DBConnectRetries > 0 && err != nil; DBConnectRetries-- {
		time.Sleep(time.Second * 5)
		storage, err = postgres.New(DSN)
	}
	if DBConnectRetries == 0 {
		panic("could not connect to database: " + err.Error())
	}
	log.Println("connected to the database")
	authService := auth.New(log, storage, tokenTTL)
	grpcServer := grpcapp.New(log, authService, grpcPort)
	return &App{
		GRPCServer: grpcServer,
	}
}
