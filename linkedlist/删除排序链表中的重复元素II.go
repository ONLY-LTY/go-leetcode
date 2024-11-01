package linkedlist

// 给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。
func deleteDuplicatesII(head *ListNode) *ListNode {
	dummyNode := &ListNode{Next: head, Val: -1}
	current := dummyNode
	for current.Next != nil && current.Next.Next != nil {
		//节点值
		val := current.Next.Val
		//节点和下一个节点相等
		if current.Next.Next.Val == val {
			//向后继续找到相同节点值
			for current.Next != nil && current.Next.Val == val {
				current.Next = current.Next.Next
			}
		} else {
			current = current.Next
		}
	}
	return dummyNode.Next
}
