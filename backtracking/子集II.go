package backtracking

import (
	"sort"
)

// 给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的
// 子集
// （幂集）。
//
// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
//
// 示例 1：
//
// 输入：nums = [1,2,2]
// 输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
// 示例 2：
//
// 输入：nums = [0]
// 输出：[[],[0]]
func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	dfsSubsetsWithDup(nums, 0, []int{}, &res)
	return res
}

func dfsSubsetsWithDup(nums []int, start int, path []int, res *[][]int) {
	temp := make([]int, len(path))
	copy(temp, path)
	*res = append(*res, temp)

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		dfsSubsetsWithDup(nums, i+1, append(path, nums[i]), res)
	}

}
