package dp

import "go-leetcode/util"

// 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
// 说明：每次只能向下或者向右移动一步。
//
// 示例 1：
//
// 输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
// 输出：7
// 解释：因为路径 1→3→1→1→1 的总和最小。
// 示例 2：
//
// 输入：grid = [[1,2,3],[4,5,6]]
// 输出：12
func minPathSum(grid [][]int) int {
	//此题可以原地修改网格数据
	m := len(grid)
	n := len(grid[0])
	//第一列 赋值
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	//第一行 赋值
	for i := 1; i < n; i++ {
		grid[0][i] += grid[0][i-1]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] = grid[i][j] + util.Min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[m-1][n-1]
}
