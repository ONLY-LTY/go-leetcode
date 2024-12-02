package array

// Given an m x n matrix, return all elements of the matrix in spiral order.
//
// Example 1:
// Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [1,2,3,6,9,8,7,4,5]
//
// Example 2:
// Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
// Output: [1,2,3,4,8,12,11,10,9,5,6,7]
//
// 解题思路
// 从左到右，顶部一层遍历完往下移一位，top++；
// 从上到下，遍历完右侧往左移一位，right--；
// 从右到左，判断top <= bottom，即是否上下都走完了。遍历完底部上移，bottom--；
// 从下到上，判断left <= right，遍历完左侧右移，left++；
func spiralOrder(matrix [][]int) []int {
	var ans []int
	left := 0
	right := len(matrix[0]) - 1
	top := 0
	bottom := len(matrix) - 1
	for left <= right && top <= bottom {
		//从左到右
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[top][i])
		}
		top++
		//从上到下
		for i := top; i <= bottom; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--
		//从右到左
		if top <= bottom {
			for i := right; i >= left; i-- {
				ans = append(ans, matrix[bottom][i])
			}
		}
		bottom--
		//从下到上
		if left <= right {
			for i := bottom; i >= top; i-- {
				ans = append(ans, matrix[i][left])
			}
		}
		left++
	}
	return ans
}
