package linkedlist

import (
	"testing"
)

func TestMergeSortList(t *testing.T) {
	res := mergeTwoLists(ArrayToListNode([]int{2, 3, 4}), ArrayToListNode([]int{4, 5, 6}))
	res.Print()
	res2 := mergeTwoLists2(ArrayToListNode([]int{2, 3, 4}), ArrayToListNode([]int{4, 5, 6}))
	res2.Print()
}

func TestGetMiddleNode(t *testing.T) {
	res := getMiddleNode(ArrayToListNode([]int{2, 3, 4, 5}))
	res.Print()
}

func TestSortList(t *testing.T) {
	res := sortList(ArrayToListNode([]int{3, 5, 4, 2}))
	if !res.Equals(ArrayToListNode([]int{2, 3, 4, 5})) {
		t.Fail()
	}
}
