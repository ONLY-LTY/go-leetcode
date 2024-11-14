package linkedlist

// 递归
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//记录第二个节点
	secondNode := head.Next
	//第一个节点指向 交换好的链表
	head.Next = swapPairs(head.Next.Next)
	//第二节点指向头结点
	secondNode.Next = head
	return secondNode
}

func swapPairs2(head *ListNode) *ListNode {
	return reverseKGroup(head, 2)
}
