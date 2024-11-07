package btree

// 给定两个整数数组 preorder 和 inorder ，
// 其中 preorder 是二叉树的先序遍历，
// inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点
// 输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
// 输出: [3,9,20,null,null,15,7]
func buildTreePre(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	//前序遍历第一个节点为根节点
	root := &TreeNode{Val: preorder[0]}
	for pos, node := range inorder {
		//根据中序遍历找到 左子树和右子树
		if node == root.Val {
			//这里中序遍历的左右子树比较好理解 [0,pos] [pos+1,end]
			//前序遍历的左子树的个数刚好等于pos 所以前序遍历的左右子树分别为 [1,pos+1] [pos+1,end]
			root.Left = buildTreePre(preorder[1:pos+1], inorder[:pos])
			root.Right = buildTreePre(preorder[pos+1:], inorder[pos+1:])
		}
	}
	return root
}
