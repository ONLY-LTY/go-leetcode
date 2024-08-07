package backtracking

import "sort"

// 给定一个可包含重复数字的整数集合 nums ，按任意顺序 返回它所有不重复的全排列。
//
// 示例 1：
//
// 输入：nums = [1,1,2]
// 输出：
// [[1,1,2],
// [1,2,1],
// [2,1,1]]
// 示例 2：
//
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
func permuteUnique(nums []int) [][]int {
	var res [][]int
	used := make([]bool, len(nums))
	sort.Ints(nums)
	dfsPermuteUnique(nums, used, []int{}, &res)
	return res
}

func dfsPermuteUnique(nums []int, used []bool, path []int, res *[][]int) {
	if len(path) == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		used[i] = true
		dfsPermuteUnique(nums, used, append(path, nums[i]), res)
		used[i] = false
	}
}
