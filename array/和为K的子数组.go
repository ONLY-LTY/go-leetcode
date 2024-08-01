package array

// subarraySum 前缀和思想
func subarraySum(nums []int, k int) int {
	//记录[0,right]的和
	sumMap := make(map[int]int)
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
