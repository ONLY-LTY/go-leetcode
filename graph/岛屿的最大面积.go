package graph

import "go-leetcode/util"

// 给你一个大小为 m x n 的二进制矩阵 grid 。
//
// 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。
// 你可以假设 grid 的四个边缘都被 0（代表水）包围着。
//
// 岛屿的面积是岛上值为 1 的单元格的数目。
//
// 计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。
func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			//如果是陆地 则开始遍历
			if grid[r][c] == 1 {
				maxArea = util.Max(maxArea, dfsArea(grid, r, c))
			}
		}
	}
	return maxArea
}

// 深度遍历
func dfsArea(grid [][]int, r int, c int) int {
	if outArea2(grid, r, c) {
		return 0
	}
	if grid[r][c] != 1 {
		return 0
	}
	//遍历了 标记一下
	grid[r][c] = 2
	area := 1
	area += dfsArea(grid, r-1, c)
	area += dfsArea(grid, r+1, c)
	area += dfsArea(grid, r, c-1)
	area += dfsArea(grid, r, c+1)
	return area
}

// 是否在数组内
func outArea2(grid [][]int, r int, c int) bool {
	return r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0])
}
