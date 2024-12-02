package linkedlist

import "testing"

func TestReorderList(t *testing.T) {
	head := ArrayToListNode([]int{1, 2, 3, 4})
	reorderList2(head)
	head.Print()

	head = ArrayToListNode([]int{1, 2, 3, 4, 5})
	reorderList2(head)
	head.Print()
}
