package btree

// 给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
//
// 路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
func pathSumIII(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := pathSumContainsRoot(root, targetSum)
	res += pathSumIII(root.Left, targetSum)
	res += pathSumIII(root.Right, targetSum)
	return res

}

// 二叉树路径合 包含跟节点
func pathSumContainsRoot(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := 0
	if targetSum == root.Val {
		res++
	}
	res += pathSumContainsRoot(root.Left, targetSum-root.Val)
	res += pathSumContainsRoot(root.Right, targetSum-root.Val)
	return res
}
