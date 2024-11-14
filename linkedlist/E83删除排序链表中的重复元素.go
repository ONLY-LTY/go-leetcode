package linkedlist

// 给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
func deleteDuplicates(head *ListNode) *ListNode {
	//只有一个节点或者没有节点 直接返回
	if head == nil || head.Next == nil {
		return head
	}
	current := head
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}
