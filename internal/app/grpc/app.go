package grpcapp

import (
	"fmt"
	"log"
	"net"

	"github.com/p3rfect05/grpc_service/internal/handlers"
	"google.golang.org/grpc"
)

type App struct {
	log        *log.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(
	log *log.Logger,
	port int,
) *App {
	gRPCServer := grpc.NewServer()
	handlers.Register(gRPCServer)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}
func (a *App) Run() error {
	const op = "grpcapp.Run"
	a.log.Printf("%s: starting grpc server...\n", op)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	a.log.Println("grpc server is running on port: ", l.Addr().String())
	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil

}

func (a *App) Stop() {
	const op = "grpcapp.Stop"
	a.log.Println(op + ": " + "server is shutting down...")
	a.gRPCServer.GracefulStop()
}
