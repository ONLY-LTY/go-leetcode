package array

//给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
//
//你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

// 主要是总结看图规律
func rotateIgraph(matrix [][]int) {
	n := len(matrix)
	tmp := make([][]int, n)
	for i := 0; i < n; i++ {
		tmp[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			//原来的第i行 变成倒数第i列
			//原来的第j列 变成第j行
			tmp[j][n-i-1] = matrix[i][j]
		}
	}
	copy(matrix, tmp)
}
