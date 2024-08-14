package stack

import "go-leetcode/util"

// 单调栈
func largestRectangleArea(heights []int) int {
	res := 0
	for i := 0; i < len(heights); i++ {
		if i == len(heights)-1 || heights[i] > heights[i+1] {
			minHeights := heights[i]
			for j := i; j >= 0; j-- {
				minHeights = util.Min(minHeights, heights[j])
				res = util.Max(res, minHeights*(i-j+1))
			}
		}
	}
	return res
}
