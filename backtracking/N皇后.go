package backtracking

import "strings"

func solveNQueens(n int) [][]string {
	var res [][]string
	chessBoard := initChessBoard(n)
	dfcSolveNQueens(0, chessBoard, &res)
	return res
}

func dfcSolveNQueens(row int, chessBoard [][]string, res *[][]string) {
	if row == len(chessBoard) {
		//如果已经是最后一行了 说明找到了
		tmp := make([]string, len(chessBoard))
		for i := 0; i < len(chessBoard); i++ {
			tmp[i] = strings.Join(chessBoard[i], "")
		}
		*res = append(*res, tmp)
		return
	}

	for cell := 0; cell < len(chessBoard); cell++ {
		//每一行从左向右 放一个棋 检查当前位置是否能放
		if isValidPos(chessBoard, row, cell) {
			//如果放了之后 当前行就不能放了
			chessBoard[row][cell] = "Q"
			//然后去下一行尝试
			dfcSolveNQueens(row+1, chessBoard, res)
			//下一行尝试完了之后 恢复
			chessBoard[row][cell] = "."
		}
	}
}

func isValidPos(chessBoard [][]string, r int, c int) bool {
	//每次同行选择一个 所以行不需要检查

	//检查列 这里只用到r行就行了
	for i := 0; i < r; i++ {
		if chessBoard[i][c] == "Q" {
			return false
		}
	}

	//左上对角线
	for i, j := r-1, c-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessBoard[i][j] == "Q" {
			return false
		}
	}
	//右上对角线
	for i, j := r-1, c+1; i >= 0 && j < len(chessBoard); i, j = i-1, j+1 {
		if chessBoard[i][j] == "Q" {
			return false
		}
	}

	return true
}

func initChessBoard(n int) [][]string {
	chessBoard := make([][]string, n)
	for i := 0; i < n; i++ {
		chessBoard[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			chessBoard[i][j] = "."
		}
	}
	return chessBoard
}
