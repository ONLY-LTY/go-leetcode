package btree

// 给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[3],[9,20],[15,7]]
func levelOrder(root *TreeNode) [][]int {
	var list [][]int
	if root == nil {
		return [][]int{}
	}
	//将根节点放入队列
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var subList []int
		//记录队列大小
		queueSize := len(queue)
		for i := 0; i < queueSize; i++ {
			//每个节点的子节点入队列
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			subList = append(subList, queue[i].Val)
		}
		list = append(list, subList)
		//父节点处队列
		queue = queue[queueSize:]
	}
	return list
}
