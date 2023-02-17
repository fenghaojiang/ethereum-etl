package client

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fenghaojiang/ethereum-etl/common/log"
	"go.uber.org/zap"
)

// Errors may show up when calling contract code
// Some error messages reveal that you should not call the contract again
func blackListError(err error) bool {
	if err == nil {
		return false
	}
	if strings.Contains(err.Error(), "execution reverted") ||
		strings.Contains(err.Error(), "abi: attempting to unmarshall an empty string while arguments are expected") ||
		strings.Contains(err.Error(), "no contract code at given address") ||
		strings.Contains(err.Error(), "out of gas") || strings.Contains(err.Error(), "gas uint64 overflow") {
		return true
	}
	return false
}

func (e *EthereumClient) wrapContractToBlackList(err error, contract common.Address) error {
	if blackListError(err) {
		log.Warn("black list error found, no need to call contract", zap.String("contract", contract.String()))
		e.blacklistContracts.Store(strings.ToLower(contract.String()), true)
	}
	return err
}
