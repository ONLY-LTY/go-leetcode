package array

// 给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
//
// 示例 1:
//
// 输入: nums = [1,2,3,4,5,6,7], k = 3
// 输出: [5,6,7,1,2,3,4]
// 解释:
// 向右轮转 1 步: [7,1,2,3,4,5,6]
// 向右轮转 2 步: [6,7,1,2,3,4,5]
// 向右轮转 3 步: [5,6,7,1,2,3,4]
// 示例 2:
//
// 输入：nums = [-1,-100,3,99], k = 2
// 输出：[3,99,-1,-100]
// 解释:
// 向右轮转 1 步: [99,-1,-100,3]
// 向右轮转 2 步: [3,99,-1,-100]
func rotate(nums []int, k int) {
	//防止k超出数组大小
	k = k % len(nums)
	//反转数组
	reverseInt(nums, 0, len(nums)-1)
	//反转[k,len-1]
	reverseInt(nums, k, len(nums)-1)
	//反转[0,k-1]
	reverseInt(nums, 0, k-1)
}

func reverseInt(nums []int, begin, end int) {
	for begin < end {
		tmp := nums[begin]
		nums[begin] = nums[end]
		nums[end] = tmp

		begin++
		end--
	}
}
