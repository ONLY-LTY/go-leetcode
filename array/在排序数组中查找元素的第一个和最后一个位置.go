package array

func searchRange(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	//先找到位置
	hitIndex := -1
	for left <= right {
		mid := left + ((right - left) / 2)
		if nums[mid] == target {
			hitIndex = mid
			break
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if hitIndex != -1 {
		//在找左右边界
		hitNum := nums[hitIndex]
		l := hitIndex
		for l >= 0 && nums[l] == hitNum {
			l--
		}
		r := hitIndex
		for r < len(nums) && nums[r] == hitNum {
			r++
		}
		return []int{l + 1, r - 1}
	}
	return []int{-1, -1}
}
