package linkedlist

// 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。
// 请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyNode := &ListNode{}
	dummyNode.Next = head
	//left的前一个节点
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	//right节点
	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}
	//left节点
	leftNode := pre.Next
	//right节点的下一个节点
	nextR := rightNode.Next
	//断链
	pre.Next = nil
	rightNode.Next = nil
	//反转left到right
	newNode := reverseList(leftNode)
	//接链
	pre.Next = newNode
	leftNode.Next = nextR

	return dummyNode.Next
}
