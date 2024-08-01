package array

import (
	"sort"
)

// towSum
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
// 你可以按任意顺序返回答案。
//
// 示例 1：
// 输入：nums = [2,7,11,15], target = 9
// 输出：[0,1]
// 解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
//
// 示例 2：
// 输入：nums = [3,2,4], target = 6
// 输出：[1,2]
//
// 示例 3：
// 输入：nums = [3,3], target = 6
// 输出：[0,1]
func towSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}

// threeSum 三数之和
// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
//
// 你返回所有和为 0 且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
// 思路: 核心要数组先排序 然后遍历每个元素 处理元素的时候使用双指针
func threeSum(nums []int) [][]int {
	var result [][]int
	//用于去重
	seen := make(map[[3]int]bool)
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		//因为num已经排序 如果当前值大于0的话 肯定不可能sum为0
		if nums[i] > 0 {
			break
		}
		L := i + 1
		R := len(nums) - 1
		for L < R {
			sum := nums[i] + nums[L] + nums[R]
			if sum == 0 {
				tri := []int{nums[i], nums[L], nums[R]}
				triKey := [3]int{nums[i], nums[L], nums[R]}
				if !seen[triKey] {
					result = append(result, tri)
					seen[triKey] = true
				}
				L++
				R--
			} else if sum < 0 {
				L++
			} else {
				R--
			}
		}
	}
	return result
}

func threeSum2(nums []int) [][]int {
	var result [][]int
	if len(nums) < 3 {
		return result
	}
	//数组排序
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		//因为num已经排序 如果当前值大于0的话 肯定不可能sum为0
		if nums[i] > 0 {
			break
		}
		//如果前一个值和当前值一样 则跳过 不然会有重复
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		L := i + 1
		R := len(nums) - 1
		for L < R {
			sum := nums[i] + nums[L] + nums[R]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[L], nums[R]})
				//因为相同元素要去重
				for L < R && nums[L] == nums[L+1] {
					//L和L+1重复 跳过
					L++
				}
				for L < R && nums[R] == nums[R-1] {
					//R和R-1重复 跳过
					R--
				}
				L++
				R--
			} else if sum < 0 {
				L++
			} else {
				R--
			}
		}
	}
	return result
}
