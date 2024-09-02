package dp

import "go-leetcode/util"

// 给定一个未排序的整数数组 nums ， 返回最长递增子序列的个数 。
//
// 注意 这个数列必须是 严格 递增的。
//
// 示例 1:
//
// 输入: [1,3,5,4,7]
// 输出: 2
// 解释: 有两个最长递增子序列，分别是 [1, 3, 4, 7] 和[1, 3, 5, 7]。
// 示例 2:
//
// 输入: [2,2,2,2,2]
// 输出: 5
// 解释: 最长递增子序列的长度是1，并且存在5个子序列的长度为1，因此输出5。
func findNumberOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	//dp[i] 表示以num[i]结尾的最长递增子序列长度
	dp := make([]int, len(nums))
	//c[i] 表示以nums[i]结尾的最长递增子序列的组合数
	c := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
		c[i] = 1
	}
	maxDp := 0
	for i := 0; i < len(nums); i++ {
		//寻找前面比num[i]小的数字 因为小的话可以把当前元素加上构成递增子序列
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[i] < dp[j]+1 {
					//当 nums[j] < nums[i]，而且 dp[j] + 1 > dp[i] 时，
					//说明第一次找到 dp[j] + 1长度且以nums[i]结尾的最长递增子序列，
					//则以 nums[i] 结尾的最长递增子序列的组合数就等于以 nums[j] 结尾的组合数，
					//即 c[i] = c[j]。
					dp[i] = dp[j] + 1
					c[i] = c[j]
				} else if dp[i] == dp[j]+1 {
					//当 nums[j] < nums[i]，而且 dp[j] + 1 == dp[i] 时，
					//说明以 nums[i] 结尾且长度为 dp[j] + 1 的递增序列已找到过一次了，
					//则以 nums[i] 结尾的最长递增子序列的组合数要加上以 nums[j] 结尾的组合数，
					//即 c[i] += c[j]。
					c[i] += c[j]
				}
			}
		}
		//保存最大子序列长度
		maxDp = util.Max(maxDp, dp[i])
	}
	res := 0
	//最大子序列长度 个数求和
	for i := 0; i < len(dp); i++ {
		if dp[i] == maxDp {
			res = res + c[i]
		}
	}
	return res
}
