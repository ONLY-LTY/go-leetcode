package dp

import "go-leetcode/util"

// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
//
// 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的
// 子序列
// 。
//
// 示例 1：
//
// 输入：nums = [10,9,2,5,3,7,101,18]
// 输出：4
// 解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
// 示例 2：
//
// 输入：nums = [0,1,0,3,2,3]
// 输出：4
// 示例 3：
//
// 输入：nums = [7,7,7,7,7,7,7]
// 输出：1
func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	//dp[i] 表示以num[i]结尾的最长递增子序列长度
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	res := 0
	for i := 0; i < len(nums); i++ {
		//寻找前面比num[i]小的数字 因为小的话可以把当前元素加上构成递增子序列
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = util.Max(dp[i], dp[j]+1)
			}
		}
		res = util.Max(res, dp[i])
	}
	return res
}
