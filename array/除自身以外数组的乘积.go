package array

// productExceptSelf 238. 除自身以外数组的乘积
// 给你一个整数数组 nums，返回 array answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
//
// 题目数据 保证 array nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
//
// 请 不要使用除法，且在 O(n) 时间复杂度内完成此题。
//
// 示例 1:
//
// 输入: nums = [1,2,3,4]
// 输出: [24,12,8,6]
// 示例 2:
//
// 输入: nums = [-1,1,0,-3,3]
// 输出: [0,0,9,0,0]
func productExceptSelf(nums []int) []int {
	//productRight[i] 表示第i个元素右边的乘积
	productRight := make([]int, len(nums))
	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		productRight[i] = right
		right = right * nums[i]
	}

	answer := make([]int, len(nums))
	productLeft := 1
	for i := 0; i < len(nums); i++ {
		//i个元素左边乘积 * i个元素右边乘积
		answer[i] = productLeft * productRight[i]
		productLeft = productLeft * nums[i]
	}
	return answer
}

// 当前元素的值 等于左边的乘积 * 右边的乘积
func productExceptSelf2(nums []int) []int {
	answer := make([]int, len(nums))
	for i := 0; i < len(answer); i++ {
		answer[i] = 1
	}
	productLeft, productRight := 1, 1
	for i := len(nums) - 1; i >= 0; i-- {
		//i元素先乘以左边乘积
		answer[i] = answer[i] * productLeft
		productLeft = productLeft * nums[i]

		//i元素再乘以右边乘积
		answer[len(nums)-i-1] *= productRight
		productRight = productRight * nums[len(nums)-i-1]
	}
	return answer
}

func productExceptSelf3(nums []int) []int {
	length := len(nums)
	answer := make([]int, length)
	for i := 0; i < length; i++ {
		answer[i] = 1
	}
	productLeft, productRight := 1, 1
	for i := 0; i < length; i++ {
		answer[i] = answer[i] * productLeft
		productLeft = productLeft * nums[i]

		r := length - 1 - i
		answer[r] = answer[r] * productRight
		productRight = productRight * nums[r]
	}
	return answer
}
