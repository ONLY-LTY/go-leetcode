package dp

import (
	"go-leetcode/util"
	"math"
	"sort"
)

// 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
//
// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
//
// 你可以认为每种硬币的数量是无限的。
//
// 示例 1：
//
// 输入：coins = [1, 2, 5], amount = 11
// 输出：3
// 解释：11 = 5 + 5 + 1
// 示例 2：
//
// 输入：coins = [2], amount = 3
// 输出：-1
// 示例 3：
//
// 输入：coins = [1], amount = 0
// 输出：0
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	res := math.MaxInt
	sort.Ints(coins)
	dfsCoinChange(coins, amount, 0, []int{}, &res)
	if res == math.MaxInt {
		return -1
	} else {
		return res
	}
}

func dfsCoinChange(candidates []int, target int, index int, path []int, count *int) {
	if target == 0 {
		*count = util.Min(*count, len(path))
		return
	}
	if target < 0 {
		return
	}
	for i := index; i < len(candidates); i++ {
		if target < candidates[i] {
			break
		}
		dfsCoinChange(candidates, target-candidates[i], i, append(path, candidates[i]), count)
	}
}

// 动态规划
func coinChangeDp(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt
	}
	//遍历背包
	for i := 1; i <= amount; i++ {
		//遍历物品
		for j := 0; j < len(coins); j++ {
			//如果背包容量i大于物品j 表示物品j可以放到背包
			//1.不放入背包的话 维持不变 dp[i]
			//2.放入背包的话 数量加一 dp[i-j] + 1
			//而且取最小值
			if i >= coins[j] && dp[i-coins[j]] != math.MaxInt {
				dp[i] = util.Min(dp[i-coins[j]]+1, dp[i])
			}
		}
	}
	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}
