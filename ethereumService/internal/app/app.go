package app

import (
	"log/slog"

	grpcapp "github.com/Dorrrke/ethereum-test-task.git/ethereumService/internal/app/grpc"
	"github.com/Dorrrke/ethereum-test-task.git/ethereumService/internal/clients/ethereumclient"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
) *App {
	// Инит хранилище

	// инит сервис
	etheClient := ethereumclient.NewEtheClient("https://goerli.infura.io/v3/a5d1234cb60d492fa3022831c896a8b0", log)

	grpcApp := grpcapp.New(log, grpcPort, etheClient)

	return &App{
		GRPCSrv: grpcApp,
	}
}
