package grpcapp

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/Dorrrke/ethereum-test-task.git/ethereumService/internal/clients/ethereumclient"
	ethereumgrpc "github.com/Dorrrke/ethereum-test-task.git/ethereumService/internal/grpc"
	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(
	log *slog.Logger,
	port int,
	etheClient *ethereumclient.EtherClient,
) *App {
	gRPCServer := grpc.NewServer()
	ethereumgrpc.Register(gRPCServer, log, etheClient)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		a.log.Info("Error start grpc", slog.String("error", err.Error()))
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"
	log := a.log.With(slog.String("op", op))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("Starting gRPC server", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil

}

func (a *App) Stop() {
	const op = "grpcapp.Stop"
	log := a.log.With(slog.String("op", op))

	log.Info("Stopping server ", slog.Int("port", a.port))
	a.gRPCServer.GracefulStop()

}
