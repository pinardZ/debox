package web3

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pinardZ/debox/bscserver/web3/erc20"
	"log"
)

type Client struct {
	rpcClient *rpc.Client
	EthClient *ethclient.Client
}

func Connect(host string) (*Client, error) {
	rpcClient, err := rpc.Dial(host)
	if err != nil {
		return nil, err
	}
	ethClient := ethclient.NewClient(rpcClient)
	return &Client{rpcClient, ethClient}, nil
}