package array

// 给定一个整数数组和一个整数 k ，请找到该数组中和为 k 的连续子数组的个数。
//
// 示例 1：
//
// 输入:nums = [1,1,1], k = 2
// 输出: 2
// 解释: 此题 [1,1] 与 [1,1] 为两种不同的情况
// 示例 2：
//
// 输入:nums = [1,2,3], k = 3
// 输出: 2
//
// subarraySum 前缀和思想
func subarraySum(nums []int, k int) int {
	//数组区间[0,right]的和
	sumMap := make(map[int]int)
	//注意这里要默认添加0
	sumMap[0] = 1
	sum := 0
	res := 0
	for i := 0; i < len(nums); i++ {
		//[0,right]的和
		sum = sum + nums[i]
		//如果存在前面[0,right]的和 + K = 当前位置[0,right]的和
		//那么说明中间有一段子数组的和为K
		if count, ok := sumMap[sum-k]; ok {
			res = res + count
		}
		//因为元素有可能是负数 所以相同和需要记录次数
		sumMap[sum]++
	}
	return res
}
