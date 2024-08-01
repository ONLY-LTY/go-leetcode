package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() {
	current := l
	fmt.Printf("[")
	for current != nil {
		if current.Next == nil {
			fmt.Printf("%d", current.Val)
		} else {
			fmt.Printf("%d ", current.Val)
		}
		current = current.Next
	}
	fmt.Println("]")
}

func (l *ListNode) Equals(other *ListNode) bool {
	if l == nil && other == nil {
		return true
	}
	if l == nil && other != nil {
		return false
	}
	for l != nil && other == nil {
		return false
	}
	currentL1 := l
	currentL2 := other
	for currentL1 != nil && currentL2 != nil {
		if currentL1.Val != currentL2.Val {
			return false
		}
		currentL1 = currentL1.Next
		currentL2 = currentL2.Next
	}
	return currentL1 == nil && currentL2 == nil
}

func ArrayToListNode(array []int) *ListNode {
	head := &ListNode{}
	current := head
	for _, v := range array {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	return head.Next
}
