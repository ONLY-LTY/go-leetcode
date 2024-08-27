package array

import (
	"go-leetcode/util"
	"sort"
)

// merge 45.合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
//
// 示例 1：
//
// 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
// 输出：[[1,6],[8,10],[15,18]]
// 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
// 示例 2：
//
// 输入：intervals = [[1,4],[4,5]]
// 输出：[[1,5]]
// 解释：区间 [1,4] 和 [4,5] 可被视为重叠区间
func merge(intervals [][]int) [][]int {
	//原区间排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int
	res = append(res, intervals[0])
	//当前处理区间
	currentIndex := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > res[currentIndex][1] {
			//新增区间
			currentIndex++
			res = append(res, intervals[i])
		} else {
			//合并区间 [[1,4],[2,3]]
			//左区间不变
			//右区间需要取最大值
			res[currentIndex][1] = util.Max(intervals[i][1], res[currentIndex][1])
		}
	}
	return res
}
