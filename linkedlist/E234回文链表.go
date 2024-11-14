package linkedlist

// isPalindrome
// 给你一个单链表的头节点 head ，请你判断该链表是否为
// 回文链表：回文 序列是向前和向后读都相同的序列。
// 如果是，返回 true ；否则，返回 false 。
// 示例 1：
// 输入：head = [1,2,2,1]
// 输出：true
//
// 示例 2：
// 输入：head = [1,2]
// 输出：false
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	var tmp []int
	for head != nil {
		tmp = append(tmp, head.Val)
		head = head.Next
	}
	mid := len(tmp)/2 + 1
	left := 0
	right := len(tmp) - 1
	for left < mid {
		if tmp[left] != tmp[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 从中间反转链表 然后比较两个链表是否相等
func isPalindrome2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow := head
	fast := slow
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	reverse := reverseList(slow)
	for reverse != nil {
		if reverse.Val != head.Val {
			return false
		}
		reverse = reverse.Next
		head = head.Next
	}
	return true
}
