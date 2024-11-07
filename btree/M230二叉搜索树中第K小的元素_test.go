package btree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKthSmallest(t *testing.T) {
	res := kthSmallest(CreateBinaryTree([]int{3, 1, 4, -1, 2}), 1)
	assert.Equal(t, 1, res)

	res = kthSmallest(CreateBinaryTree([]int{5, 3, 6, 2, 4, -1, -1, 1}), 3)
	assert.Equal(t, 3, res)
}
