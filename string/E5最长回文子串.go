package string

// 给你一个字符串 s，找到 s 中最长的回文子串 。
// 示例 1：
//
// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。
// 示例 2：
//
// 输入：s = "cbbd"
// 输出："bb"
func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		//单字符中心 aba
		res1 := lookCenter(s, i, i)
		//双字符中心 aabb
		res2 := lookCenter(s, i, i+1)

		//比较最大
		if len(res1) > len(res) {
			res = res1
		}
		if len(res2) > len(res) {
			res = res2
		}
	}
	return res
}

// 中心向两边扩散 找到回文串
func lookCenter(s string, i int, j int) string {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	return s[i+1 : j]
}
