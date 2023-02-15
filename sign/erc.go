package sign

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// prevent variables changed by other usage
var _ercTransferEventSigHash = crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))
var _ercApprovalEventSigHash = crypto.Keccak256Hash([]byte("Approval(address,address,uint256)"))

func ErcTransferEventSigHash() common.Hash {
	return _ercTransferEventSigHash
}

func ErcApprovalEventSigHash() common.Hash {
	return _ercApprovalEventSigHash
}
