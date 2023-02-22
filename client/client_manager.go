package client

import (
	"fmt"
	"sync"
)

type ClientManager struct {
	NetworkClients sync.Map
}

type Network string

const (
	EthereumNetwork          Network = "Ethereum"
	ArbitrumNetwork          Network = "Arbitrum"
	OptimismNetwork          Network = "Optimism"
	PolygonNetwork           Network = "Polygon"
	BinanceSmartChainNetwork Network = "BinanceSmartChain"
)

func NewClientManager() *ClientManager {
	return &ClientManager{
		NetworkClients: sync.Map{},
	}
}

func (cm *ClientManager) NetworkClient(network Network) (*EthereumClient, error) {
	c, ok := cm.NetworkClients.Load(network)
	if !ok {
		return nil, fmt.Errorf("network client not exists in client manager")
	}
	return c.(*EthereumClient), nil
}

func (cm *ClientManager) AddNetworkClient(network Network, e *EthereumClient) {
	cm.NetworkClients.Store(network, e)
}
