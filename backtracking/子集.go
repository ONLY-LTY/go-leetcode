package backtracking

// 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的
// 子集
// （幂集）。
//
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
//
// 示例 1：
//
// 输入：nums = [1,2,3]
// 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
// 示例 2：
//
// 输入：nums = [0]
// 输出：[[],[0]]
func subsets(nums []int) [][]int {
	var res [][]int
	dfsSubsets(nums, 0, []int{}, &res)
	return res
}
func dfsSubsets(nums []int, level int, path []int, res *[][]int) {
	temp := make([]int, len(path))
	copy(temp, path)
	*res = append(*res, temp)

	for i := level; i < len(nums); i++ {
		//添加元素
		//path = append(path, nums[i])
		//遍历
		dfsSubsets(nums, i+1, append(path, nums[i]), res)
		//回溯
		//path = path[:len(path)-1]
	}
}
