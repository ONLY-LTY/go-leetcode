package graph

// 在给定的 m x n 网格 grid 中，每个单元格可以有以下三个值之一：
//
// 值 0 代表空单元格；
// 值 1 代表新鲜橘子；
// 值 2 代表腐烂的橘子。
// 每分钟，腐烂的橘子 周围 4 个方向上相邻 的新鲜橘子都会腐烂。
//
// 返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1 。
func orangesRotting(grid [][]int) int {
	//新鲜橘子的数量
	freshNum := 0
	//腐烂橘子队列
	var rotQueue []Coordinates
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 1 {
				freshNum++
			}
			if grid[r][c] == 2 {
				rotQueue = append(rotQueue, Coordinates{r: r, c: c})
			}
		}
	}
	minutes := 0
	//bfs 腐烂橘子
	for len(rotQueue) > 0 {
		//新鲜橘子没有了 返回过去了多少分钟
		if freshNum == 0 {
			return minutes
		}
		//过去一分钟
		minutes++
		//对目前所有的腐烂橘子进行感染
		size := len(rotQueue)
		for i := 0; i < size; i++ {
			rot := rotQueue[i]
			//感染上下左右橘子 并把感染的橘子加入到队列
			freshNum -= roting(grid, rot.r-1, rot.c, &rotQueue)
			freshNum -= roting(grid, rot.r+1, rot.c, &rotQueue)
			freshNum -= roting(grid, rot.r, rot.c-1, &rotQueue)
			freshNum -= roting(grid, rot.r, rot.c+1, &rotQueue)
		}
		//新感染的橘子
		rotQueue = rotQueue[:size]
	}
	if freshNum > 0 {
		return -1
	}
	return minutes
}

func roting(grid [][]int, r int, c int, rotQueue *[]Coordinates) int {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] != 1 {
		return 0
	}
	grid[r][c] = 2
	*rotQueue = append(*rotQueue, Coordinates{r: r, c: c})
	return 1
}

type Coordinates struct {
	r int
	c int
}
