package btree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumNumbers(t *testing.T) {
	res := sumNumbers(CreateBinaryTree([]int{1, 2, 3}))
	assert.Equal(t, 25, res)

	res = sumNumbers(CreateBinaryTree([]int{4, 9, 0, 5, 1}))
	assert.Equal(t, 1026, res)
}
