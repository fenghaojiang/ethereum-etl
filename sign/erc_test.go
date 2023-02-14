package sign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnErcSignatureHash(t *testing.T) {
	t.Run("test on erc signature", func(t *testing.T) {
		assert.Equal(t, ErcTransferEventSigHash().String(), "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
		assert.Equal(t, ErcApprovalEventSigHash().String(), "0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
	})
}
