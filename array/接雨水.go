package array

import (
	"go-leetcode/util"
)

// trap 42.接雨水
// 参考：https://zhuanlan.zhihu.com/p/79811305
func trap(height []int) int {
	length := len(height)
	//maxRight[i] 代表第 i 列右边最高的墙的高度,不包括i
	maxRight := make([]int, length)
	for i := length - 2; i > 0; i-- {
		maxRight[i] = util.Max(maxRight[i+1], height[i+1])
	}
	maxLeft := 0
	totalSum := 0
	for i := 1; i < length-1; i++ {
		// i列左边最高墙的高度 不包括i
		maxLeft = util.Max(maxLeft, height[i-1])
		// 比较 i列左边和右边最高墙
		minimum := util.Min(maxLeft, maxRight[i])
		// 计算i列能接水的数量
		if minimum > height[i] {
			totalSum = totalSum + (minimum - height[i])
		}
	}
	return totalSum
}
