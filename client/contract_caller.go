package client

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/fenghaojiang/ethereum-etl/common/log"
	"go.uber.org/zap"
)

func (e *EthereumClient) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	// contract code would not change by block number in normal ETL jobs
	// we can add contract code in cache to skip `eth_getCode` to get acceleration for contract call
	if value, ok := e.blacklistContracts.Load(strings.ToLower(contract.String())); ok && value.(bool) {
		log.Info("contract code in black list, no need to call again")
		return nil, nil
	}
	cacheCode, _ := e.cacheCode(contract)
	if len(cacheCode) != 0 {
		return cacheCode, nil
	}
	var contractCode hexutil.Bytes
	backOffFunc := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), e.maxRetries)
	err := backoff.Retry(func() error {
		timeCtx, cancel := context.WithTimeout(context.Background(), e.timeout)
		defer cancel()
		c := e.Client()
		if c == nil {
			return fmt.Errorf("no clients available")
		}
		err := c.CallContext(timeCtx, &contractCode, "eth_getCode", contract, blockNumber)
		if e.wrapContractToBlackList(err, contract) != nil {
			log.Error("failed to fetch code from node", zap.String("contract", contract.String()), zap.Error(err))
			return err
		}
		return nil
	}, backOffFunc)
	if err != nil {
		log.Error("fetch code from node reach max retries", zap.String("contract", contract.String()), zap.Error(err))
		return nil, err
	}
	err = e.saveContract(contract, contractCode)
	if err != nil {
		log.Error("failed to save contract code to big cache", zap.Error(err))
	}
	// ignore the save contract error
	return contractCode, nil
}

func (e *EthereumClient) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return nil, nil
}
