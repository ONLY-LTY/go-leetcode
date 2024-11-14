package dp

import "go-leetcode/util"

// 给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
//
// 示例 1：
//
// 输入：nums = [1,5,11,5]
// 输出：true
// 解释：数组可以分割成 [1, 5, 5] 和 [11] 。
// 示例 2：
//
// 输入：nums = [1,2,3,5]
// 输出：false
// 解释：数组不能分割成两个元素和相等的子集。
func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
	}
	// 和为奇数的话 肯定不可能
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([]int, target+1)
	//遍历物品 这里可以看做物品的重量和价值相等
	for i := 0; i < len(nums); i++ {
		//倒序遍历背包
		for j := target; j >= nums[i]; j-- {
			dp[j] = util.Max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	return target == dp[target]
}
