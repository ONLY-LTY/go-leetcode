package linkedlist

// 递归
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}

// 遍历
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head := &ListNode{}
	current := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			current = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			current = l2
			l2 = l2.Next
		}
	}
	if l1 != nil {
		current.Next = l1
	}
	if l2 != nil {
		current.Next = l2
	}
	return head.Next
}
