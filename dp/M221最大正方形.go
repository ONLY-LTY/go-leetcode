package dp

import (
	"go-leetcode/util"
	"math"
)

// 在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
//
// 示例 1：
//
// 输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
// 输出：4
// 示例 2：
//
// 输入：matrix = [["0","1"],["1","0"]]
// 输出：1
// 示例 3：
//
// 输入：matrix = [["0"]]
// 输出：0
func maximalSquare(matrix [][]byte) int {
	//dp[i][j] 表示以i,j结尾能组成最大正方形的边长
	dp := newIntArray(len(matrix)+1, len(matrix[0])+1)
	res := math.MinInt
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = util.Min(dp[i-1][j-1], util.Min(dp[i-1][j], dp[i][j-1])) + 1
				}
				if dp[i][j] > res {
					res = dp[i][j]
				}
			}
		}
	}
	return res * res
}
