package backtracking

import "go-leetcode/graph"

// 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
//
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
// 同一个单元格内的字母不允许被重复使用。

// 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
// 输出：true
func exist(board [][]byte, word string) bool {
	r := len(board)
	c := len(board[0])
	//初始化访问标记二维数组
	visited := make([][]bool, r)
	for i := 0; i < r; i++ {
		visited[i] = make([]bool, c)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			//找到开始位置 进行回溯
			if board[i][j] == word[0] &&
				dfsSearch(board, visited, i, j, 0, word) {
				return true
			}
		}
	}
	return false
}

var direction = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func dfsSearch(board [][]byte, visited [][]bool, r int, c int, index int, word string) bool {
	//当前节点已经访问过了 或者当前节点超出矩阵范围 或者当前节点和当前字符不相等
	if graph.OutArea(board, r, c) || visited[r][c] || board[r][c] != word[index] {
		return false
	}
	//遍历到了最后一个字符
	if index == len(word)-1 {
		return true
	}
	//当前节点标记访问
	visited[r][c] = true
	//上下左右开始搜索下一个单词字符
	for i := 0; i < len(direction); i++ {
		if dfsSearch(board, visited, r+direction[i][0], c+direction[i][1], index+1, word) {
			return true
		}
	}
	//搜索完成之后还原标记位
	visited[r][c] = false
	return false
}
