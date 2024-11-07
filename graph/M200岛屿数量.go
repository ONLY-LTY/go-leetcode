package graph

// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//
// 此外，你可以假设该网格的四条边均被水包围。
func numIslands(grid [][]byte) int {
	res := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			//如果是陆地 则开始遍历
			if grid[r][c] == '1' {
				res++
				dfs(grid, r, c)
			}
		}
	}
	return res
}

// 深度遍历
func dfs(grid [][]byte, r int, c int) {
	if OutArea(grid, r, c) {
		return
	}
	if grid[r][c] != '1' {
		return
	}
	//遍历了 标记一下
	grid[r][c] = '2'
	dfs(grid, r-1, c)
	dfs(grid, r+1, c)
	dfs(grid, r, c-1)
	dfs(grid, r, c+1)
}

// OutArea 是否在数组内
func OutArea(grid [][]byte, r int, c int) bool {
	return r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0])
}
