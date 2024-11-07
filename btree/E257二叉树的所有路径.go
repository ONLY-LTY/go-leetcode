package btree

import "strconv"

// Given a binary tree, return all root-to-leaf paths.
//
// Note: A leaf is a node with no children.
// Input:
//
//	1
//
// /   \
// 2     3
//
//	\
//	 5
//
// Output: ["1->2->5", "1->3"]
//
// Explanation: All root-to-leaf paths are: 1->2->5, 1->3
func binaryTreePaths(root *TreeNode) []string {
	var res []string
	dfsPath(root, "", &res)
	return res
}

// 前序遍历
func dfsPath(root *TreeNode, path string, res *[]string) {
	if root == nil {
		return
	}
	if path == "" {
		path = strconv.Itoa(root.Val)
	} else {
		path = path + "->" + strconv.Itoa(root.Val)
	}
	if root.Left == nil && root.Right == nil {
		*res = append(*res, path)
	}
	dfsPath(root.Left, path, res)
	dfsPath(root.Right, path, res)
}
