package btree

func buildTreePost(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			root.Left = buildTreePre(inorder[:i], postorder[:i])
			root.Right = buildTreePre(inorder[i+1:], postorder[i:len(postorder)-1])
		}
	}
	return root
}
