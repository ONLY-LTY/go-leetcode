package btree

// 给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。
// （即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
// 示例 1：
//
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[3],[20,9],[15,7]]
// 示例 2：
//
// 输入：root = [1]
// 输出：[[1]]
// 示例 3：
//
// 输入：root = []
// 输出：[]
func zigzagLevelOrder(root *TreeNode) [][]int {
	var resList [][]int
	if root == nil {
		return resList
	}
	var queue []*TreeNode
	queue = append(queue, root)
	order := true
	for len(queue) > 0 {
		var subList []int
		queueSize := len(queue)
		for i := 0; i < queueSize; i++ {
			node := queue[i]
			subList = append(subList, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		//逆序
		if !order {
			start, end := 0, len(subList)-1
			for start <= end {
				subList[start], subList[end] = subList[end], subList[start]
				start++
				end--
			}
		}
		queue = queue[queueSize:]
		resList = append(resList, subList)
		order = !order
	}
	return resList
}
