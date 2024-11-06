package btree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZigzagLevelOrder(t *testing.T) {
	res := zigzagLevelOrder(CreateBinaryTree([]int{3, 9, 20, -1, -1, 15, 7}))
	assert.Equal(t, [][]int{{3}, {20, 9}, {15, 7}}, res)
}
