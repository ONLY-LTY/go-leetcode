package btree

// 给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
// 每条从根节点到叶节点的路径都代表一个数字：
//
// 例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
// 计算从根节点到叶节点生成的 所有数字之和 。
//
// 叶节点 是指没有子节点的节点。
// Input: [1,2,3]
//
//	  1
//	 / \
//	2   3
//
// Output: 25
// 解释：
// 从根到叶子节点路径 1->2 代表数字 12
// 从根到叶子节点路径 1->3 代表数字 13
// 因此，数字总和 = 12 + 13 = 25
//
// Input: [4,9,0,5,1]
//
//	  4
//	 / \
//	9   0
//
// / \
// 5   1
// Output: 1026
// 解释：
// 从根到叶子节点路径 4->9->5 代表数字 495
// 从根到叶子节点路径 4->9->1 代表数字 491
// 从根到叶子节点路径 4->0 代表数字 40
// 因此，数字总和 = 495 + 491 + 40 = 1026
func sumNumbers(root *TreeNode) int {
	var res int
	dfs(root, 0, &res)
	return res
}

// 前序遍历
func dfs(root *TreeNode, sum int, res *int) {
	if root == nil {
		return
	}
	sum = sum*10 + root.Val
	//叶子节点
	if root.Left == nil && root.Right == nil {
		*res += sum
		return
	}
	dfs(root.Left, sum, res)
	dfs(root.Right, sum, res)
}
