package ethereumgrpc

import (
	"context"
	"log"
	"log/slog"

	"github.com/Dorrrke/ethereum-test-task.git/ethereumService/internal/clients/ethereumclient"
	etherservice1 "github.com/Dorrrke/proto-ethereum/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverApi struct {
	etherservice1.UnimplementedEthereumServiceServer
	client *ethereumclient.EtherClient
}

func Register(gPRC *grpc.Server, log *slog.Logger, etheClient *ethereumclient.EtherClient) {
	etherservice1.RegisterEthereumServiceServer(gPRC, &serverApi{client: etheClient})
}

func (s *serverApi) GetBalance(ctx context.Context, req *etherservice1.GetBalanceRequest) (*etherservice1.GetBalanceResponse, error) {
	// panic("GetBalance not implemented")
	balance, err := s.client.GetBalance(req.GetEthereumAddr())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	log.Println(balance)
	return &etherservice1.GetBalanceResponse{
		EthereumBalance: balance.String(),
	}, nil
}

func (s *serverApi) GetLatestBlock(ctx context.Context, req *etherservice1.GetLatestBlockRequest) (*etherservice1.GetLatestBlockResponse, error) {
	// panic("GetLatestBlock not implemented")
	block, err := s.client.GetLatestBlock()
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &etherservice1.GetLatestBlockResponse{
		BlockNumber:       block.Number,
		TransactionsCount: int64(block.TransactionsCount),
		BlockComplicacy:   block.BlockComplicacy,
		Date:              block.Date,
	}, nil
}

func (s *serverApi) VerifyAddress(ctx context.Context, req *etherservice1.VerifyAddressRequest) (*etherservice1.VerifyAddressResponse, error) {
	// panic("VerifyAddress not implemented")
	//  TODO: Переделать под передачу адреса извне
	res, err := s.client.VerifyAddress(req.GetAddr())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &etherservice1.VerifyAddressResponse{
		NumberValid: res,
	}, nil
}
