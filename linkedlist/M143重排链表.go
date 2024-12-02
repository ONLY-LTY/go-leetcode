package linkedlist

// 给定一个单链表 L 的头节点 head ，单链表 L 表示为：
//
// L0 → L1 → … → Ln - 1 → Ln
// 请将其重新排列后变为：
//
// L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
// 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
// 示例 1：
// 输入：head = [1,2,3,4]
// 输出：[1,4,2,3]

// 示例 2：
// 输入：head = [1,2,3,4,5]
// 输出：[1,5,2,4,3]
func reorderList(head *ListNode) {
	//数组保存
	var array []*ListNode
	current := head
	for current != nil {
		array = append(array, current)
		current = current.Next
	}
	//按照规律排序
	start := 0
	end := len(array) - 1
	for start < end {
		array[start].Next = array[end]
		start++
		if start == end {
			break
		}
		array[end].Next = array[start]
		end--
	}
	array[start].Next = nil
}

func reorderList2(head *ListNode) {
	//中间节点
	middleNode := getMiddleNode(head)
	middleNext := middleNode.Next
	middleNode.Next = nil
	//反转后面
	newHead := reverseList(middleNext)
	//直接合并
	for head != nil && newHead != nil {
		headNext := head.Next
		newHeadNext := newHead.Next

		head.Next = newHead
		head = headNext

		newHead.Next = head
		newHead = newHeadNext
	}
}
