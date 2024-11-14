package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	head := ArrayToListNode([]int{1, 2, 2, 1})
	assert.Equal(t, isPalindrome(head), true)

	head2 := ArrayToListNode([]int{1, 2, 2, 1})
	assert.Equal(t, isPalindrome(head2), true)
}
