package linkedlist

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
// 整体思路是让前面的指针先移动 n 步，之后前后指针共同移动直到前面的指针到尾部为止
// 设预先指针 pre 的下一个节点指向 head，设前指针为 start，后指针为 end，二者都等于 pre
// start 先向前移动n步 之后 start 和 end 共同向前移动，此时二者的距离为 n，当 start 到尾部时，end 的位置恰好为倒数第 n 个节点
// 因为要删除该节点，所以要移动到该节点的前一个才能删除，所以循环结束条件为 start.next != null
// 删除后返回 pre.next，为什么不直接返回 head 呢，因为 head 有可能是被删掉的点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pre := &ListNode{}
	pre.Next = head
	slow := pre
	fast := pre
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return slow
}
