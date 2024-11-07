package btree

// 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 小的元素（从 1 开始计数）。
// 输入：root = [3,1,4,null,2], k = 1
// 输出：1
// 输入：root = [5,3,6,2,4,null,null,1], k = 3
// 输出：3
// 中序遍历到K个元素
func kthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
	return 0
}
