package hash

// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
// 示例 1：
//
// 输入：nums = [100,4,200,1,3,2]
// 输出：4
// 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
// 示例 2：
//
// 输入：nums = [0,3,7,2,5,8,4,6,0,1]
// 输出：9
func longestConsecutive(nums []int) int {
	set := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		set[num] = struct{}{}
	}
	res := 0
	for _, num := range nums {
		//num-1不存在 从num开始遍历
		if !exists(set, num-1) {
			curNum := num + 1
			curCount := 1
			for exists(set, curNum) {
				curNum++
				curCount++
			}
			if curCount > res {
				res = curCount
			}
		}
	}
	return res
}
func exists(set map[int]struct{}, target int) bool {
	_, ok := set[target]
	return ok
}
