package array

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
func maxSubArray(nums []int) int {
	/**
	用 temp 记录局部最优值，用 result 记录全局最优值。
	每遍历一个新元素时，判断（已遍历的连续子数组的和）加上（当前元素值），与（当前元素值）对比谁更大。
	（1）如果已遍历的连续子数组的和 + 当前元素值 >= 当前元素值
		说明（已遍历的连续子数组的和）是大于等于0的，令 temp = 已遍历的连续子数组的和 + 当前元素值。
	（2）如果已遍历的连续子数组的和 + 当前元素值 < 当前元素值
		说明（已遍历的连续子数组的和）是小于0的，加上这部分只会拖累当前元素，故应该直接抛弃掉这部分，令 temp = 当前元素值。
	（3）对比 temp 和 result，如果 temp 更大，则更新到 result 中
	*/
	maxVal, tempMax := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		tempMax = tempMax + nums[i]
		tempMax = util.Max(tempMax, nums[i])
		maxVal = util.Max(maxVal, tempMax)
	}
	return maxVal
}
