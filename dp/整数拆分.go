package dp

import "go-leetcode/util"

// 给定一个正整数 n ，将其拆分为 k 个 正整数 的和（ k >= 2 ），并使这些整数的乘积最大化。
//
// 返回 你可以获得的最大乘积 。
//
// 示例 1:
//
// 输入: n = 2
// 输出: 1
// 解释: 2 = 1 + 1, 1 × 1 = 1。
// 示例 2:
//
// 输入: n = 10
// 输出: 36
// 解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
func integerBreak(n int) int {
	//定义 dp[i] 代表 i 拆分至少两个正整数后的最大乘积
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			//1. i*(i-j) 代表的是将 i 拆分为 j 和 i-j，且不再对 i-j 继续进行拆分. 因为最小拆分2个
			//2. j×dp[i−j] 代表的是将 i 拆分成 j 和 i-j，且对 i-j 继续拆分，
			//	 拆分为多个整数，而 dp[i−j] 是在之前的动态规划已经求得的整数 i-j 分割后的最大乘积
			dp[i] = util.Max(dp[i], util.Max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}
