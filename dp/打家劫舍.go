package dp

import (
	"go-leetcode/btree"
	"go-leetcode/util"
)

// rob 198. 打家劫舍I
// 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
//
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
//
// 示例 1：
//
// 输入：[1,2,3,1]
// 输出：4
// 解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
//
//	偷窃到的最高金额 = 1 + 3 = 4 。
//
// 示例 2：
//
// 输入：[2,7,9,3,1]
// 输出：12
// 解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
//
//	偷窃到的最高金额 = 2 + 9 + 1 = 12
func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	//如果偷第i房间，那么dp[i] = dp[i - 2] + nums[i] ，即：第i-1房一定是不考虑的，找出 下标i-2（包括i-2）以内的房屋，最多可以偷窃的金额为dp[i-2] 加上第i房间偷到的钱。
	//
	//如果不偷第i房间，那么dp[i] = dp[i - 1]，即考 虑i-1房，（注意这里是考虑，并不是一定要偷i-1房，这是很多同学容易混淆的点）
	//
	//然后dp[i]取最大值，即dp[i] = max(dp[i - 2] + nums[i], dp[i - 1]);
	dp := make([]int, length)
	dp[0] = nums[0]
	dp[1] = util.Max(nums[0], nums[1])
	for i := 2; i < length; i++ {
		dp[i] = util.Max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[length-1]
}

func robRange(nums []int, start int, end int) int {
	length := end - start + 1
	if length == 1 {
		return nums[start]
	}
	dp := make([]int, length)
	dp[0] = nums[start]
	dp[1] = util.Max(nums[start], nums[start+1])
	for i := 2; i < length; i++ {
		dp[i] = util.Max(dp[i-2]+nums[i+start], dp[i-1])
	}
	return dp[length-1]
}

// rob2 打家劫舍II
// 你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
//
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。
//
// 示例 1：
//
// 输入：nums = [2,3,2]
// 输出：3
// 解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
// 示例 2：
//
// 输入：nums = [1,2,3,1]
// 输出：4
// 解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
//
//	偷窃到的最高金额 = 1 + 3 = 4 。
//
// 示例 3：
//
// 输入：nums = [1,2,3]
// 输出：3
func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	//分成两个数组去解决
	return util.Max(robRange(nums, 0, len(nums)-2),
		robRange(nums, 1, len(nums)-1))
}

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

// dp 返回一个大小为 2 的数组
// arr[0] 表示不抢 node 的话，得到的最大钱数
// arr[1] 表示抢 node 的话，得到的最大钱数
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
