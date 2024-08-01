package linkedlist

// 给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
//
// k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
func reverseKGroup(head *ListNode, k int) *ListNode {
	//如果没有节点 或者只有一个节点 直接返回
	if head == nil || head.Next == nil {
		return head
	}
	//找到第k个节点
	node := head
	for i := 1; i < k; i++ {
		node = node.Next
		//遍历期间如果不够k个 直接返回
		if node == nil {
			return head
		}
	}
	//第k个节点的下一个节点 保存下来
	tmp := node.Next
	//第K个节点断链
	node.Next = nil
	//反转K个节点
	newHead := reverseList(head)
	//之前的头节点 现在变成尾结点 这里尾节点指向 剩余反转好的
	head.Next = reverseKGroup(tmp, k)
	return newHead
}
