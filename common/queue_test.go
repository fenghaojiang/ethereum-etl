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

		q.PushBack(TestStruct{BlockNumber: 1, BlockHash: "0x1"})
		assert.Equal(t, TestStruct{BlockNumber: 1, BlockHash: "0x1"}, q.Front())
		q.PushBack(TestStruct{BlockNumber: 2, BlockHash: "0x2"})
		assert.Equal(t, TestStruct{BlockNumber: 1, BlockHash: "0x1"}, q.PopFront())
		assert.Equal(t, TestStruct{BlockNumber: 2, BlockHash: "0x2"}, q.Front())
	})
}
