package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnQueue(t *testing.T) {

	type TestStruct struct {
		BlockNumber uint64
		BlockHash   string
	}
	t.Run("test on queue", func(t *testing.T) {

		q := NewQueue[TestStruct]()
		f := q.Front()
		assert.Equal(t, TestStruct{}, f)
	})
}
