package btree

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
