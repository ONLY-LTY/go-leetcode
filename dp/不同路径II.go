package dp

// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
//
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish”）。
//
// 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
//
// 网格中的障碍物和空位置分别用 1 和 0 来表示。
//
// 示例 1：
//
// 输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
// 输出：2
// 解释：3x3 网格的正中间有一个障碍物。
// 从左上角到右下角一共有 2 条不同的路径：
// 1. 向右 -> 向右 -> 向下 -> 向下
// 2. 向下 -> 向下 -> 向右 -> 向右
// 示例 2：
//
// 输入：obstacleGrid = [[0,1],[0,0]]
// 输出：1
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	//如果在起点或终点出现了障碍，直接返回0
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// 初始化第一列, 如果是障碍物, 后面的就都是0, 不用循环了
	for i := 0; i < m && obstacleGrid[i][0] != 1; i++ {
		dp[i][0] = 1
	}
	// 初始化第一行, 如果是障碍物, 后面的就都是0, 不用循环了
	for i := 0; i < n && obstacleGrid[0][i] != 1; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 如果obstacleGrid[i][j]这个点是障碍物, 那么dp[i][j]保持为0
			if obstacleGrid[i][j] != 1 {
				// 否则我们需要计算当前点可以到达的路径数
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
