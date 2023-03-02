package client

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/fenghaojiang/ethereum-etl/common"
	"github.com/fenghaojiang/ethereum-etl/model"
	"go.uber.org/zap"
)

type EthereumClient struct {
	networkClients           []*rpc.Client
	networkContractCodeCache *bigcache.BigCache
	blacklistContracts       sync.Map
	maxRetries               uint64
	timeout                  time.Duration
	queue                    *common.Queue[model.BlockBrief]
}

var _ bind.ContractCaller = (*EthereumClient)(nil)

func NewEthereumClient() *EthereumClient {
	contractCodeCache, err := newBigCacheWithDefaultConfig(context.Background())
	if err != nil {
		log.Fatal("failed to new bigcache for contract code", zap.Error(err))
	}
	return &EthereumClient{
		networkContractCodeCache: contractCodeCache,
		queue:                    common.NewQueue[model.BlockBrief](),
	}
}

func (e *EthereumClient) Client() *rpc.Client {
	if len(e.networkClients) == 0 {
		return nil
	}
	rand.Seed(time.Now().Unix())
	return e.networkClients[rand.Intn(len(e.networkClients))]
}
