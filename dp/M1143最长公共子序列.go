package dp

import "go-leetcode/util"

// 给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
//
// 一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
//
// 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
// 两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。
//
// 示例 1：
//
// 输入：text1 = "abcde", text2 = "ace"
// 输出：3
// 解释：最长公共子序列是 "ace" ，它的长度为 3 。
// 示例 2：
//
// 输入：text1 = "abc", text2 = "abc"
// 输出：3
// 解释：最长公共子序列是 "abc" ，它的长度为 3 。
// 示例 3：
//
// 输入：text1 = "abc", text2 = "def"
// 输出：0
// 解释：两个字符串没有公共子序列，返回 0 。
//
// 提示：
//
// 1 <= text1.length, text2.length <= 1000
// text1 和 text2 仅由小写英文字符组成。
func longestCommonSubsequence(text1 string, text2 string) int {
	//dp[i][j]：长度为i的字符串text1与长度为j的字符串text2的最长公共子序列为dp[i][j]
	//i、j从1开始
	dp := make([][]int, len(text1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				//如果text1的第i个字符 与 text2的第j个字符相同，那么找到了一个公共元素，
				//所以dp[i][j] = dp[i - 1][j - 1] + 1;
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				//如果text1第i个字符 与 text2第j个字符 不相同
				//那就看看text1前 i-1 个字符 与text2前 j 个字符的最长公共子序列 和
				//text1前 i 个字符与text2前 j-1 个字符的最长公共子序列，取最大的。
				dp[i][j] = util.Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}
