package backtracking

// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//
// 示例 1：
//
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// 示例 2：
//
// 输入：nums = [0,1]
// 输出：[[0,1],[1,0]]
// 示例 3：
//
// 输入：nums = [1]
// 输出：[[1]]
func permute(nums []int) [][]int {
	var paths [][]int
	dfs(nums, []int{}, &paths)
	return paths
}

func dfs(nums []int, path []int, paths *[][]int) {
	if len(nums) == len(path) {
		//这里需要新数组记录 排序 因为path在整体遍历中会不停的修改
		tmp := make([]int, len(path))
		copy(tmp, path)
		*paths = append(*paths, tmp)
	}
	//循环处理每个数字
	for i := 0; i < len(nums); i++ {
		//数字已经在path中 就跳过
		if contains(path, nums[i]) {
			continue
		}
		// 添加到path中
		//path = append(path, nums[i])
		// 继续遍历
		dfs(nums, append(path, nums[i]), paths)
		// 移除最后一个数字
		//path = path[:len(path)-1]
	}
}

func contains(path []int, num int) bool {
	for i := 0; i < len(path); i++ {
		if path[i] == num {
			return true
		}
	}
	return false
}
