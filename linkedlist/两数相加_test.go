package linkedlist

import (
	"testing"
)

func TestArrayToListNode(t *testing.T) {
	ArrayToListNode([]int{2, 4, 3})
}

func TestAddTwoNumbers(t *testing.T) {
	res := addTwoNumbers(ArrayToListNode([]int{2, 4, 3}), ArrayToListNode([]int{5, 6, 4}))
	res.Print()
	if !res.Equals(ArrayToListNode([]int{7, 0, 8})) {
		t.Fail()
	}

	res = addTwoNumbers(ArrayToListNode([]int{0}), ArrayToListNode([]int{0}))
	res.Print()
	if !res.Equals(ArrayToListNode([]int{0})) {
		t.Fail()
	}

	res = addTwoNumbers(ArrayToListNode([]int{9, 9, 9, 9, 9, 9, 9}), ArrayToListNode([]int{9, 9, 9, 9}))
	res.Print()
	if !res.Equals(ArrayToListNode([]int{8, 9, 9, 9, 0, 0, 0, 1})) {
		t.Fail()
	}
}
