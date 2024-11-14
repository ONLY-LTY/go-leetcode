package dp

import (
	"github.com/stretchr/testify/assert"
	"go-leetcode/btree"
	"testing"
)

func TestRob(t *testing.T) {
	nums := []int{2, 7, 9, 3, 1}
	assert.Equal(t, 12, rob(nums))

	nums = []int{2, 3, 2}
	assert.Equal(t, 3, rob2(nums))

	nums = []int{1, 2, 3, 1}
	assert.Equal(t, 4, rob2(nums))

	nums = []int{1, 2, 3}
	assert.Equal(t, 3, rob2(nums))

	nums = []int{3, 2, 3, -1, 3, -1, 1}
	assert.Equal(t, 7, rob3(btree.CreateBinaryTree(nums)))
}
