package btree

import "math"

// isValidBST
// 因为中序遍历搜索二叉树的结果是有序的 所以在遍历的过程中比较前一个节点值即可
func isValidBST(root *TreeNode) bool {
	//数组模拟栈
	var stack []*TreeNode
	//前一个节点的值
	inorder := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		//获取最后一个元素
		root = stack[len(stack)-1]
		//模拟出栈
		stack = stack[:len(stack)-1]
		//当前节点如果比前一个节点小 则返回false
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}
