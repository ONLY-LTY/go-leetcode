package btree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	res := levelOrder(CreateBinaryTree([]int{3, 9, 20, -1, -1, 15, 7}))
	assert.Equal(t, [][]int{{3}, {9, 20}, {15, 7}}, res)
}
