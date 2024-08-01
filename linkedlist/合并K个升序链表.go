package linkedlist

// 给你一个链表数组，每个链表都已经按升序排列。
//
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	left := 0
	right := len(lists) - 1
	for right > 0 {
		for left < right {
			//两两合并
			lists[left] = mergeTwoLists(lists[left], lists[right])
			left++
			right--
		}
		//重新开始
		left = 0
	}
	return lists[0]
}
