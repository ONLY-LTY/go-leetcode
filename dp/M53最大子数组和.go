package dp

import "go-leetcode/util"

// maxSubArray 53. 最大子数组和
// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
// 子数组
// 是数组中的一个连续部分。
//
// 示例 1：
//
// 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出：6
// 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
// 示例 2：
//
// 输入：nums = [1]
// 输出：1
// 示例 3：
//
// 输入：nums = [5,4,-1,7,8]
// 输出：23
//
// dp只依赖上一个值 根据滚动数组 可以优化去掉数组 见: maxSubArray2
func maxSubArray(nums []int) int {
	//dp[i] 表示以i结尾的连续子数组最大和
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxRes := nums[0]
	for i := 1; i < len(nums); i++ {
		//i-1结尾的加上当前元素 和 当前元素取最大值
		dp[i] = util.Max(dp[i-1]+nums[i], nums[i])
		//全局最大值
		if dp[i] > maxRes {
			maxRes = dp[i]
		}
	}
	return maxRes
}

func maxSubArray2(nums []int) int {
	//用 temp 记录局部最优值，用 result 记录全局最优值。
	//每遍历一个新元素时，判断（已遍历的连续子数组的和）加上（当前元素值），与（当前元素值）对比谁更大。
	maxVal, tempMax := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		tempMax = tempMax + nums[i]
		tempMax = util.Max(tempMax, nums[i])
		maxVal = util.Max(maxVal, tempMax)
	}
	return maxVal
}
