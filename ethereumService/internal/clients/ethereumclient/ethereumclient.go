package ethereumclient

import (
	"context"
	"log/slog"
	"math/big"
	"time"

	"github.com/Dorrrke/ethereum-test-task.git/ethereumService/internal/domain/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EtherClient struct {
	ethClient *ethclient.Client
	log       *slog.Logger
	addr      string
}

func NewEtheClient(addr string, logr *slog.Logger) *EtherClient {
	client, err := ethclient.Dial("https://goerli.infura.io/v3/a5d1234cb60d492fa3022831c896a8b0")
	if err != nil {
		logr.Error("Error connect to network", slog.String("error", err.Error()))
		panic(err)
	}
	return &EtherClient{
		ethClient: client,
		log:       logr,
		addr:      addr,
	}
}

func (ec *EtherClient) GetBalance(addr string) (big.Int, error) {
	const op = "ethereumclient.GetBalance"
	log := ec.log.With(slog.String("op", op))
	account := common.HexToAddress(addr)
	balance, err := ec.ethClient.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Error("Error get balance", slog.String("error", err.Error()))
		return *big.NewInt(0), err
	}
	return *balance, nil
}

func (ec *EtherClient) GetLatestBlock() (models.Block, error) {
	const op = "ethereumclient.GetLatestBlock"
	log := ec.log.With(slog.String("op", op))
	header, err := ec.ethClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Error("Error get latest block", slog.String("error", err.Error()))
		return models.Block{Number: "", TransactionsCount: 0, BlockComplicacy: "", Date: ""}, err
	}
	blockNumber := header.Number

	block, err := ec.ethClient.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Error("Error get block info by number", slog.String("error", err.Error()))
		return models.Block{Number: "", TransactionsCount: 0, BlockComplicacy: "", Date: ""}, err
	}

	tm := time.Unix(int64(block.Time()), 0)

	blockInfo := models.Block{
		Number:            block.Number().String(),
		TransactionsCount: block.Transactions().Len(),
		BlockComplicacy:   block.Difficulty().String(),
		Date:              tm.Format("2 January 2006 15:04"),
	}

	return blockInfo, nil

}

func (ec *EtherClient) VerifyAddress(addr string) (bool, error) {
	const op = "ethereumclient.VerifyAddress"
	log := ec.log.With(slog.String("op", op))
	address := common.HexToAddress(addr)
	bytecode, err := ec.ethClient.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Error("Error check bitecode", slog.String("error", err.Error()))
		return false, err
	}

	isContract := len(bytecode) > 0

	return !isContract, nil
}
