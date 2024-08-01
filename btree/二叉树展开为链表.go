package btree

// 给你二叉树的根结点 root ，请你将它展开为一个单链表：
//
// 展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
// 展开后的单链表应该与二叉树 先序遍历 顺序相同
// 递归实现
var prev *TreeNode

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten2(root.Right)
	flatten2(root.Left)
	root.Right = prev
	root.Left = nil
	prev = root
}

// 非递归实现 先序遍历放入到数组 在构建二叉树
func flatten2(root *TreeNode) {
	var list []int
	perOrder(root, &list)
	current := root
	for i := 1; i < len(list); i++ {
		current.Left = nil
		current.Right = &TreeNode{Val: list[i]}
		current = current.Right
	}
}

// 前序遍历
func perOrder(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	*list = append(*list, root.Val)
	perOrder(root.Left, list)
	perOrder(root.Right, list)
}
