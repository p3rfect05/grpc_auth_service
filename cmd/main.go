package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/p3rfect05/grpc_service/internal/app"
	"github.com/p3rfect05/grpc_service/internal/config"
)

var logger = log.New(os.Stdout, "[LOG]\t", log.Ldate|log.Ltime)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
	dsn := config.MustLoadDatabaseURL()
	application := app.New(logger, cfg.GRPC.Port, cfg.Storage, cfg.TokenTTL, dsn)
	if err := runMigration(dsn); err != nil {
		panic(err)
	}
	application.GRPCServer.MustRun()
}

func runMigration(dsn config.DatabaseURLConfig) error {
	dsnString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dsn.User, dsn.Password, dsn.Host, dsn.Port, dsn.DBName)
	m, err := migrate.New("file:///app/migrations", dsnString)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			logger.Println("everything is up-to-date")
			return nil
		}
		return err
	}
	return nil
}
