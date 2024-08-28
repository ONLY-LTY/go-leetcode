package dp

import (
	"go-leetcode/btree"
	"go-leetcode/util"
)

// rob3 打家劫舍III
// 小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 root 。
//
// 除了 root 之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果 两个直接相连的房子在同一天晚上被打劫 ，房屋将自动报警。
//
// 给定二叉树的 root 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额 。
//
// 示例 1:
//
// 输入: root = [3,2,3,null,3,null,1]
// 输出: 7
// 解释: 小偷一晚能够盗取的最高金额 3 + 3 + 1 = 7
// 示例 2:
//
// 输入: root = [3,4,5,1,3,null,1]
// 输出: 9
// 解释: 小偷一晚能够盗取的最高金额 4 + 5 = 9
func rob3(root *btree.TreeNode) int {
	res := dp(root)
	return util.Max(res[0], res[1])
}

// 后续遍历
// dp 返回一个大小为 2 的数组
// dp[0] 表示不抢 node 的话，得到的最大钱数
// dp[1] 表示抢 node 的话，得到的最大钱数
func dp(node *btree.TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}
	left := dp(node.Left)
	right := dp(node.Right)

	//抢当前node 则下家就不能抢了
	doRob := node.Val + left[0] + right[0]

	//不抢当前node 下家可以抢也可以不抢
	doNotRob := util.Max(left[0], left[1]) + util.Max(right[0], right[1])

	return []int{doNotRob, doRob}
}
