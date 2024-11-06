package slidingwindow

// 给定两个字符串 s 和 p，找到 s 中所有 p 的 变位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
//
// 变位词 指字母相同，但排列不同的字符串。
//
// 示例 1：
//
// 输入: s = "cbaebabacd", p = "abc"
// 输出: [0,6]
// 解释:
// 起始索引等于 0 的子串是 "cba", 它是 "abc" 的变位词。
// 起始索引等于 6 的子串是 "bac", 它是 "abc" 的变位词。
// 示例 2：
//
// 输入: s = "abab", p = "ab"
// 输出: [0,1,2]
// 解释:
// 起始索引等于 0 的子串是 "ab", 它是 "ab" 的变位词。
// 起始索引等于 1 的子串是 "ba", 它是 "ab" 的变位词。
// 起始索引等于 2 的子串是 "ab", 它是 "ab" 的变位词。
func findAnagrams(s string, p string) []int {
	var result []int
	if len(s) < len(p) {
		return result
	}
	var pStrCount [26]int
	//统计P字符串中每个字符的个数
	for i := 0; i < len(p); i++ {
		pStrCount[p[i]-'a']++
	}
	//初始化窗口 窗口大小为P字符串的长度
	windowStart := 0
	windowEnd := 0
	reset := len(p)
	for windowEnd < len(p) {
		pStrCount[s[windowEnd]-'a']--
		//如果是非P字符串中的 这里肯定小于0
		if pStrCount[s[windowEnd]-'a'] >= 0 {
			reset--
		}
		windowEnd++
	}
	//P字符串中的字符个数已经没有了 符合条件
	if reset == 0 {
		result = append(result, windowStart)
	}
	for windowEnd < len(s) {
		//左边元素出窗口
		if pStrCount[s[windowStart]-'a'] >= 0 {
			reset++
		}
		pStrCount[s[windowStart]-'a']++
		windowStart++

		//右边元素进窗口
		pStrCount[s[windowEnd]-'a']--
		if pStrCount[s[windowEnd]-'a'] >= 0 {
			reset--
		}
		windowEnd++

		//符合条件 记录
		if reset == 0 {
			result = append(result, windowStart)
		}
	}
	return result
}
