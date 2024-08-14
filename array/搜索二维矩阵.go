package array

// 二维转换一维 二分查找
func searchMatrix(matrix [][]int, target int) bool {
	left := 0
	row := len(matrix)
	cell := len(matrix[0])
	right := row*cell - 1
	for left <= right {
		mid := (left + right) / 2
		num := matrix[mid/cell][mid%cell]
		if num == target {
			return true
		} else if num < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// 从左左下脚开始遍历元素
func searchMatrix2(matrix [][]int, target int) bool {
	row := len(matrix) - 1
	cell := 0
	for row >= 0 && cell <= len(matrix[0])-1 {
		if matrix[row][cell] == target {
			return true
		}
		//做下角元素大于目标元素 往前一行搜索
		if matrix[row][cell] > target {
			row--
		} else {
			//否则像下一列搜索
			cell++
		}
	}
	return false
}
