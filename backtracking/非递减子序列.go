package backtracking

// 给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的长度至少是2。
//
// 示例:
//
// 输入: [4, 6, 7, 7]
// 输出: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7], [4,7,7]]
// 说明:
//
// 给定数组的长度不会超过15。
// 数组中的整数范围是 [-100,100]。
// 给定数组中可能包含重复数字，相等的数字应该被视为递增的一种情况。
func findSubsequences(nums []int) [][]int {
	var res [][]int
	dfsFindSubsequences(nums, 0, []int{}, &res)
	return res
}

func dfsFindSubsequences(nums []int, start int, path []int, res *[][]int) {
	if len(path) >= 2 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
	}
	//同层保存重复数字 遍历标记
	used := make(map[int]bool)
	for i := start; i < len(nums); i++ {
		//重复数字不处理
		if used[nums[i]] {
			continue
		}
		//当前数字比上一个数字大才处理 要满足递增
		if len(path) == 0 || nums[i] >= path[len(path)-1] {
			//因为是同层 这里回溯完了之后不需要状态回溯
			used[nums[i]] = true
			dfsFindSubsequences(nums, i+1, append(path, nums[i]), res)
		}
	}
}
