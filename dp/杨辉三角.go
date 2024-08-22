package dp

// 给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
//
// 在「杨辉三角」中，每个数是它左上方和右上方的数的和。
// 示例 1:
//
// 输入: numRows = 5
// 输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
// 示例 2:
//
// 输入: numRows = 1
// 输出: [[1]]
func generate(numRows int) [][]int {
	var result [][]int
	for i := 0; i < numRows; i++ {
		var row []int
		for j := 0; j < i+1; j++ {
			//每一行的开头和结尾都是1
			if j == 0 || j == i {
				row = append(row, 1)
			} else if i > 1 {
				//左上方和右上方的数的和
				row = append(row, result[i-1][j-1]+result[i-1][j])
			}
		}
		result = append(result, row)
	}
	return result
}
