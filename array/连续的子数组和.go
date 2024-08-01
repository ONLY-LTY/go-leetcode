package array

// checkSubarraySum
// 用一个变量 sum 记录累加和。 子数组的元素和可以用前缀和相减得到，
// 例如 [i,j] 区间内的元素和，可以由 prefixSum[j] - prefixSum[i] 得到。当 prefixSums[j]−prefixSums[i] 为 k 的倍数时，
// prefixSums[i] 和 prefixSums[j] 除以 k 的余数相同。因此只需要计算每个下标对应的前缀和除以 k 的余数即可，
// 使用 map 存储每个余数第一次出现的下标即可。在 map 中如果存在相同余数的 key，代表当前下标和 map 中这个 key 记录的下标可以满足总和为 k 的倍数这一条件。
// 再判断一下能否满足大小至少为 2 的条件即可。用 2 个下标相减，长度大于等于 2 即满足条件，可以输出 true
func checkSubarraySum(nums []int, k int) bool {
	m := make(map[int]int)
	m[0] = -1
	sum := 0
	for i, n := range nums {
		sum = sum + n
		if r, ok := m[sum%k]; ok {
			if i-2 >= r {
				return true
			}
		} else {
			m[sum%k] = i
		}
	}
	return false
}
