package client

import (
	"context"
	"log"

	"github.com/allegro/bigcache/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

type EthereumClient struct {
	networkClients           map[string][]*ethclient.Client
	networkContractCodeCache *bigcache.BigCache
}

var _ bind.ContractCaller = (*EthereumClient)(nil)

func NewEthereumClient() *EthereumClient {
	contractCodeCache, err := newBigCacheWithDefaultConfig(context.Background())
	if err != nil {
		log.Fatal("failed to new bigcache for contract code", zap.Error(err))
	}

	return &EthereumClient{
		networkContractCodeCache: contractCodeCache,
	}
}
