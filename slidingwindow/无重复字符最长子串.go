package slidingwindow

// lengthOfLongestSubstring
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串的长度。
//
// 示例 1:
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
// 示例 2:
// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
// 示例 3:
// 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//
// 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	//定义布尔数组 字符是否重复标记
	var bitSet [256]bool
	//使用滑动窗口
	var result, windowStart, windowEnd = 0, 0, 0
	for windowEnd < len(s) {
		if bitSet[s[windowEnd]] {
			//发现有重复的 左边界开始缩小 直到重复的字符移出了左边界
			bitSet[s[windowStart]] = false
			windowStart++
		} else {
			//窗口右边界不断地右移 并保存重复位置
			bitSet[s[windowEnd]] = true
			windowEnd++
		}
		if result < windowEnd-windowStart {
			result = windowEnd - windowStart
		}
	}
	return result
}
