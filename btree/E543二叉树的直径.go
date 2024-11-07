package btree

import "go-leetcode/util"

// diameterOfBinaryTree
// 给你一棵二叉树的根节点，返回该树的 直径 。
//
// 二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。
//
// 两节点之间路径的 长度 由它们之间边数表示。
func diameterOfBinaryTree(root *TreeNode) int {
	//可以将二叉树的直径转换为：二叉树的 每个节点的左右子树的高度和 的最大值。
	var res int
	var deepF func(root *TreeNode) int
	deepF = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := deepF(root.Left)
		right := deepF(root.Right)
		res = util.Max(res, left+right)
		return util.Max(left, right) + 1
	}
	deepF(root)
	return res
}
