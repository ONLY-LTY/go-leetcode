package btree

// 给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
//
// 叶子节点 是指没有子节点的节点。
//
// 示例 1：
//
// 输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
// 输出：[[5,4,11,2],[5,8,4,5]]
// 示例 2：
//
// 输入：root = [1,2,3], targetSum = 5
// 输出：[]
// 示例 3：
//
// 输入：root = [1,2], targetSum = 0
// 输出：[]
//
// 提示：
func pathSumII(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	pathSumDfs(root, targetSum, []int(nil), &res)
	return res
}

func pathSumDfs(root *TreeNode, targetSum int, path []int, res *[][]int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		*res = append(*res, append([]int{}, path...))
	}
	pathSumDfs(root.Left, targetSum-root.Val, path, res)
	pathSumDfs(root.Right, targetSum-root.Val, path, res)
}
