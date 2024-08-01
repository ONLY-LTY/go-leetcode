package linkedlist

// sortList 链表排序 归并排序
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	//中间节点截断 分成两个链表
	middleNode := getMiddleNode(head)
	tmp := middleNode.Next
	middleNode.Next = nil

	//排序左边的链表
	left := sortList(head)
	//排序右边的链表
	right := sortList(tmp)
	//合并有序链表
	return mergeTwoLists(left, right)
}

// getMiddleNode 获取链表的中间节点 快慢指针
func getMiddleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
