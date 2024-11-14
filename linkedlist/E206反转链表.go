package linkedlist

// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
// 递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//反转下一个节点之后的整个链表
	node := reverseList(head.Next)
	//head下一个节点指向 反转之后最后一个节点
	//这里令反转之后的最后一个节点指向head
	head.Next.Next = head
	//head的下一个节点指向nil
	head.Next = nil
	return node
}

// 遍历
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode
	current := head
	for current != nil {
		//1. 保存当前节点的下一个节点
		next := current.Next
		//2. 令当前节点的下一个节点指向前面反转好的节点
		current.Next = pre
		//3. 令前面反转好的 指向当前节点
		pre = current
		//4. 令下一个节点为当前节点 继续循环
		current = next
	}
	return pre
}
