package btree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// CreateBinaryTree 按层构建二叉树 空节点 用-1
func CreateBinaryTree(arr []int) *TreeNode {
	// 如果数组为空，则返回nil
	if len(arr) == 0 {
		return nil
	}

	// 创建一个节点的队列，用于按层构建二叉树
	var queue []*TreeNode

	// 创建根节点，并将其加入队列
	root := &TreeNode{Val: arr[0]}
	queue = append(queue, root)

	// 初始化一个索引i，它将在数组中移动
	i := 1

	// 迭代直到数组所有元素都被处理
	for i < len(arr) && len(queue) > 0 {
		// 出队列
		node := queue[0]
		queue = queue[1:]

		// 如果左孩子存在，则创建它并加入队列
		if i < len(arr) && arr[i] != -1 { // 假设-1表示nil节点
			node.Left = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Left)
		}
		i++

		// 如果右孩子存在，则创建它并加入队列
		if i < len(arr) && arr[i] != -1 { // 假设-1表示nil节点
			node.Right = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
