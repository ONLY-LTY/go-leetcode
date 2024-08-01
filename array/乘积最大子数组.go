package array

import (
	"go-leetcode/util"
	"math"
)

// maxProduct 152. 乘积最大子数组
// 给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续
// 子数组
// （该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
//
// 测试用例的答案是一个 32-位 整数。
//
// 示例 1:
//
// 输入: nums = [2,3,-2,4]
// 输出: 6
// 解释: 子数组 [2,3] 有最大乘积 6。
// 示例 2:
//
// 输入: nums = [-2,0,-1]
// 输出: 0
// 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
func maxProduct(nums []int) int {
	//由于存在负数，那么会导致最大的变最小的，最小的变最大的。因此还需要维护当前最小值min。
	tmpMin, tmpMax := 1, 1
	maxVal := math.MinInt
	for i := 0; i < len(nums); i++ {
		tmpMin = tmpMin * nums[i]
		tmpMax = tmpMax * nums[i]

		tmpMin = util.Min(nums[i], util.Min(tmpMin, tmpMax))
		tmpMax = util.Max(nums[i], util.Max(tmpMin, tmpMax))

		maxVal = util.Max(maxVal, tmpMax)
	}
	return maxVal
}
