package client

import "sync"

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
