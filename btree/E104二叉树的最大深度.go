package btree

import "go-leetcode/util"

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return util.Max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
