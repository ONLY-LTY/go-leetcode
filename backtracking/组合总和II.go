package backtracking

import "sort"

// candidates 中的每个数字在每个组合中只能使用一次，解集不能包含重复的组合。
//
// 示例 1:
//
// 输入: candidates = [10,1,2,7,6,1,5], target = 8,
// 输出:
// [
// [1,1,6],
// [1,2,5],
// [1,7],
// [2,6]
// ]
// 示例 2:
//
// 输入: candidates = [2,5,2,1,2], target = 5,
// 输出:
// [
// [1,2,2],
// [5]
// ]
//
// 提示:
//
// 1 <= candidates.length <= 100
// 1 <= candidates[i] <= 50
// 1 <= target <= 30
func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	sort.Ints(candidates)
	dfsCombinationSum2(candidates, target, 0, []int{}, &res)
	return res
}

func dfsCombinationSum2(candidates []int, target int, start int, path []int, res *[][]int) {
	if target == 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}
	for i := start; i < len(candidates); i++ {
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		if target < candidates[i] {
			return
		}
		dfsCombinationSum2(candidates, target-candidates[i], i+1, append(path, candidates[i]), res)
	}
}
