package client

import (
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthereumClient struct {
	networkClients           map[string][]*ethclient.Client
	networkContractCodeCache sync.Map //map[string][]byte
}

var _ bind.ContractCaller = (*EthereumClient)(nil)

func NewEthereumClient() *EthereumClient {
	return &EthereumClient{}
}
