package dp

import (
	"go-leetcode/util"
	"math"
)

// maxProfit 买卖股票的最佳时机
// 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
//
// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
//
// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
//
// 示例 1：
//
// 输入：[7,1,5,3,6,4]
// 输出：5
// 解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//
//	注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
//
// 示例 2：
//
// 输入：prices = [7,6,4,3,1]
// 输出：0
// 解释：在这种情况下, 没有交易完成, 所以最大利润为 0。
func maxProfit(prices []int) int {
	//当天之前股价最小值
	previousMin := math.MaxInt
	maxProfit := 0
	for _, price := range prices {
		if price < previousMin {
			previousMin = price
		}
		maxProfit = util.Max(maxProfit, price-previousMin)
	}
	return maxProfit
}

func newIntArray(row int, cell int) [][]int {
	array := make([][]int, row)
	for i := range array {
		array[i] = make([]int, cell)
	}
	return array
}

// maxProfit2 股票买卖II 无限次交易次数
// 题目：
// 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
// 示例 1:
// 输入: [7,1,5,3,6,4]
// 输出: 7
// 解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4。随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
// 示例 2:
// 输入: [1,2,3,4,5]
// 输出: 4
// 解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
// 示例 3:
// 输入: [7,6,4,3,1]
// 输出: 0
// 解释: 在这种情况下, 没有交易完成, 所以最大利润为 0
func maxProfit2(prices []int) int {
	length := len(prices)
	dp := newIntArray(length, 2)
	//第一天不持有股票 利润
	dp[0][0] = 0
	//第一天持有股票 利润
	dp[0][1] = -prices[0]

	for i := 1; i < length; i++ {
		//今天不持有股票 最大利润
		//1. 昨天本来就不持有股票
		//2. 昨天持有股票，今天买了
		dp[i][0] = util.Max(dp[i-1][0], dp[i-1][1]+prices[i])
		//今天持有股票 最大利润
		//1. 昨天本来就持有股票 今天不操作
		//2. 昨天不持有股票，今天买入
		dp[i][1] = util.Max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return util.Max(dp[length-1][0], dp[length-1][1])
}
