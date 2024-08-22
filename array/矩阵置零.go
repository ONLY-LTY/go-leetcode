package array

// 给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
func setZeroes(matrix [][]int) {
	//需要置零的行
	row := make([]bool, len(matrix))
	//需要置零的列
	cell := make([]bool, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				cell[j] = true
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if row[i] || cell[j] {
				matrix[i][j] = 0
			}
		}
	}
}
