package dp

// maxArea
// 11. 盛最多水的容器
// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
//
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
// 返回容器可以储存的最大水量
// 示例 1：

// 输入：[1,8,6,2,5,4,8,3,7]
// 输出：49
// 解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
// 示例 2：
//
// 输入：height = [1,1]
// 输出：1
func maxArea(height []int) int {
	maxA := 0
	start := 0
	end := len(height) - 1
	for start < end {
		//容器底部宽度
		width := end - start
		//容器高度
		high := 0
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}
		//容器面积
		temp := width * high
		if temp > maxA {
			maxA = temp
		}
	}
	return maxA
}
