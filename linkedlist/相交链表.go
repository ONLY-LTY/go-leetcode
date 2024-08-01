package linkedlist

// 给你两个单链表的头节点 headA 和 headB ，
// 请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
// 思路:主要计算出两个链表的长度差 然后是两个指针保持这个长度差
// 第一轮遍历之后pA和pB已经维持长度差
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA := headA
	pB := headB
	//这里pA和pB要么指向相交的节点  要么都指向NULL 不会死循环
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}
