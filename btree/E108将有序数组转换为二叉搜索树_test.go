package btree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	nums := []int{-10, -3, 0, 5, 9}
	if !assert.Equal(t, inorderTraversal(sortedArrayToBST(nums)), nums) {
		t.Fail()
	}
}
