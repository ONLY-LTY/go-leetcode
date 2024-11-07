package btree

// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
//
// 百度百科中最近公共祖先的定义为：
// 对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）
// 输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
// 输出：3
// 解释：节点 5 和节点 1 的最近公共祖先是节点 3 。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	//q p的最近公共祖先
	//1. 都在左子树里面
	//2. 都在右子树里面
	//3. 左右子树一边一个节点
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	pathP := make([]*TreeNode, 0)
	findPath(root, p, &pathP)
	pathQ := make([]*TreeNode, 0)
	findPath(root, q, &pathQ)

	var ancestor *TreeNode
	for i := 0; i < len(pathP) && i < len(pathQ); i++ {
		if pathP[i] == pathQ[i] {
			ancestor = pathP[i]
		} else {
			break
		}
	}
	return ancestor
}

// findPath 查找root节点到target节点的路径 保存在path中
func findPath(root, target *TreeNode, path *[]*TreeNode) bool {
	if root == nil {
		return false
	}
	*path = append(*path, root)
	if root == target {
		return true
	}
	if (root.Left != nil && findPath(root.Left, target, path)) ||
		(root.Right != nil && findPath(root.Right, target, path)) {
		return true
	}
	*path = (*path)[:len(*path)-1]
	return false
}
