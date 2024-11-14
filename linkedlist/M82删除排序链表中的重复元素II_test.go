package linkedlist

import "testing"

func TestDeleteDuplicates(t *testing.T) {
	head := ArrayToListNode([]int{1, 2, 3, 3, 4, 4, 5})
	deleteDuplicatesII(head).Print()
}
