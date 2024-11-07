package btree

// 输入：root = [1,null,2,3]
// 输出：[1,3,2]
// 示例 2：
//
// 输入：root = []
// 输出：[]
// 示例 3：
//
// 输入：root = [1]
// 输出：[1]
func inorderTraversal(root *TreeNode) (res []int) {
	//中序遍历 左中右
	//前序遍历 中左右
	//后序遍历 左右中
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}
