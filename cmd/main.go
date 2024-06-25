package main

import (
	"fmt"
	"log"
	"os"

	"github.com/p3rfect05/grpc_service/internal/app"
	"github.com/p3rfect05/grpc_service/internal/config"
)

var logger = log.New(os.Stdout, "[LOG]\t", log.Ldate|log.Ltime)

func main() {
	os.Setenv("CONFIG_PATH", "./config/config.yaml")
	cfg := config.MustLoad()
	fmt.Println(cfg)

	application := app.New(logger, cfg.GRPC.Port, cfg.Storage, cfg.TokenTTL)
	application.GRPCServer.MustRun()
}
