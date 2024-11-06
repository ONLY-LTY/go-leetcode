package linkedlist

import (
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	head := ArrayToListNode([]int{1, 2, 3, 4, 5})
	reverseKGroup(head, 2).Print()
}
