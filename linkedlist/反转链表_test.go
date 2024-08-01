package linkedlist

import (
	"testing"
)

func TestReverseList(t *testing.T) {
	head := ArrayToListNode([]int{1, 2, 3, 4, 5})
	actual := reverseList(head)
	expected := ArrayToListNode([]int{5, 4, 3, 2, 1})
	if !actual.Equals(expected) {
		t.Fail()
	}

	head2 := ArrayToListNode([]int{1, 2, 3, 4, 5})
	actual2 := reverseList2(head2)
	if !actual2.Equals(expected) {
		t.Fail()
	}
}
