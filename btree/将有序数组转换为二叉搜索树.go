package btree

// 给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵
// *平衡* 二叉搜索树。 此题主要是平衡
func sortedArrayToBST(nums []int) *TreeNode {
	var helper func(nums []int, left int, right int) *TreeNode
	helper = func(nums []int, left int, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := (left + right) / 2
		//选择任意一个中间数字作为根节点
		root := &TreeNode{Val: nums[mid]}
		//左树
		root.Left = helper(nums, left, mid-1)
		//右树
		root.Right = helper(nums, mid+1, right)
		return root
	}
	return helper(nums, 0, len(nums)-1)
}
