package backtracking

// 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是
// 回文串 :是向前和向后读都相同的字符串。
// 。返回 s 所有可能的分割方案。
//
// 示例 1：
//
// 输入：s = "aab"
// 输出：[["a","a","b"],["aa","b"]]
// 示例 2：
//
// 输入：s = "a"
// 输出：[["a"]]
//
// 提示：
//
// 1 <= s.length <= 16 s 仅由小写英文字母组成
// 参考: https://programmercarl.com/0131.%E5%88%86%E5%89%B2%E5%9B%9E%E6%96%87%E4%B8%B2.html#%E6%80%9D%E8%B7%AF
func partition(s string) [][]string {
	var res [][]string
	dfsPartition(s, 0, []string{}, &res)
	return res
}

func dfsPartition(str string, start int, path []string, res *[][]string) {
	if start == len(str) {
		temp := make([]string, len(path))
		copy(temp, path)
		*res = append(*res, temp)
	}
	for i := start; i < len(str); i++ {
		if isPalindrome(str, start, i) {
			dfsPartition(str, i+1, append(path, str[start:i+1]), res)
		}
	}
}
func isPalindrome(str string, start int, end int) bool {
	for start < end {
		if str[start] != str[end] {
			return false
		}
		start++
		end--
	}
	return true
}
