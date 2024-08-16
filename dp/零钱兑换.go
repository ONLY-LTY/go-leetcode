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

// 给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。
//
// 请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。
//
// 假设每一种面额的硬币有无限个。
//
// 题目数据保证结果符合 32 位带符号整数。
//
// 示例 1：
//
// 输入：amount = 5, coins = [1, 2, 5]
// 输出：4
// 解释：有四种方式可以凑成总金额：
// 5=5
// 5=2+2+1
// 5=2+1+1+1
// 5=1+1+1+1+1
// 示例 2：
//
// 输入：amount = 3, coins = [2]
// 输出：0
// 解释：只用面额 2 的硬币不能凑成总金额 3 。
// 示例 3：
//
// 输入：amount = 10, coins = [10]
// 输出：1
// **排列问题 先遍历背包 在遍历物品**
// **组合问题 先遍历物品 在遍历背包**
func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	//遍历物品
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] = dp[j] + dp[j-coins[i]]
		}
	}
	return dp[amount]
}
