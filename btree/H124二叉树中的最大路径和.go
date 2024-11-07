package btree

import (
	"go-leetcode/util"
	"math"
)

// 二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。
// 同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
//
// 路径和 是路径中各节点值的总和。
//
// 给你一个二叉树的根节点 root ，返回其 最大路径和 。
// 输入：root = [1,2,3]
// 输出：6
// 解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6
//
// 输入：root = [-10,9,20,null,null,15,7]
// 输出：42
// 解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42
func maxPathSum(root *TreeNode) int {
	var maxSum = math.MinInt
	maxPathSumRecursion(root, &maxSum)
	return maxSum
}

// maxPathSumRecursion 从下往上遍历 求得每个节点的最大路径合 并在遍历的过程中 保存全局最大和
func maxPathSumRecursion(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}
	//左子树
	maxL := maxPathSumRecursion(root.Left, maxSum)
	//右子树
	maxR := maxPathSumRecursion(root.Right, maxSum)
	//当前节点和左右子树组合的路径和 并在遍历过程中保存最大的路径和
	*maxSum = util.Max(maxR+maxL+root.Val, *maxSum)
	//当前节点的最大路径和
	return util.Max(0, util.Max(maxL, maxR)+root.Val)
}
