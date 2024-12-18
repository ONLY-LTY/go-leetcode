package array

import (
	"sort"
	"strconv"
	"strings"
)

// 给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。
//
// 注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。
//
// 示例 1：
//
// 输入：nums = [10,2]
// 输出："210"
// 示例 2：
//
// 输入：nums = [3,30,34,5,9]
// 输出："9534330"
func largestNumber(nums []int) string {
	numStr := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		numStr[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(numStr, func(i, j int) bool {
		a := strings.Join([]string{numStr[i], numStr[j]}, "")
		b := strings.Join([]string{numStr[j], numStr[i]}, "")
		return strings.Compare(a, b) > 0
	})
	return strings.Join(numStr, "")
}
