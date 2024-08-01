package leetcode

import "strings"

// convert
// 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
//
// 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
//
// P   A   H   N
// A P L S I I G
// Y   I   R
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
//
// 示例 1：
// 输入：s = "PAYPALISHIRING", numRows = 3
// 输出："PAHNAPLSIIGYIR"

// 示例 2：
// 输入：s = "PAYPALISHIRING", numRows = 4
// 输出："PINALSIGYAHRPI"

// 示例 3：
// 输入：s = "A", numRows = 1
// 输出："A"
//
// 直接总结数学规律 输出结果
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([]string, numRows)
	length := len(s)
	k := 2*numRows - 2
	for i := 0; i < length; i++ {
		rows[min(k-i%k, i%k)] += string(s[i])
	}
	var builder strings.Builder
	for _, row := range rows {
		builder.WriteString(row)
	}
	return builder.String()
}

// convert2
// 按照规则填充矩阵 然后输出结果
func convert2(s string, numRows int) string {
	matrix, down, up := make([][]byte, numRows), 0, numRows-2
	for i := 0; i != len(s); {
		if down != numRows {
			matrix[down] = append(matrix[down], s[i])
			down++
			i++
		} else if up > 0 {
			matrix[up] = append(matrix[up], s[i])
			up--
			i++
		} else {
			up = numRows - 2
			down = 0
		}
	}
	solution := make([]byte, 0, len(s))
	for _, row := range matrix {
		for _, item := range row {
			solution = append(solution, item)
		}
	}
	return string(solution)
}
