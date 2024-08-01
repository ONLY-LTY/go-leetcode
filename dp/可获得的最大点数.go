package dp

import (
	"go-leetcode/util"
)

// maxScore 1423. 可获得的最大点数
// 几张卡牌 排成一行，每张卡牌都有一个对应的点数。点数由整数数组 cardPoints 给出。
//
// 每次行动，你可以从行的开头或者末尾拿一张卡牌，最终你必须正好拿 k 张卡牌。
//
// 你的点数就是你拿到手中的所有卡牌的点数之和。
//
// 给你一个整数数组 cardPoints 和整数 k，请你返回可以获得的最大点数。
//
// 示例 1：
//
// 输入：cardPoints = [1,2,3,4,5,6,1], k = 3
// 输出：12
// 解释：第一次行动，不管拿哪张牌，你的点数总是 1 。但是，先拿最右边的卡牌将会最大化你的可获得点数。最优策略是拿右边的三张牌，最终点数为 1 + 6 + 5 = 12 。
// 示例 2：
//
// 输入：cardPoints = [2,2,2], k = 2
// 输出：4
// 解释：无论你拿起哪两张卡牌，可获得的点数总是 4 。
// 示例 3：
//
// 输入：cardPoints = [9,7,7,9,7,7,9], k = 7
// 输出：55
// 解释：你必须拿起所有卡牌，可以获得的点数为所有卡牌的点数之和。
// 示例 4：
//
// 输入：cardPoints = [1,1000,1], k = 1
// 输出：1
// 解释：你无法拿到中间那张卡牌，所以可以获得的最大点数为 1 。
// 示例 5：
//
// 输入：cardPoints = [1,79,80,1,1,1,200,1], k = 3
// 输出：202
func maxScore(cardPoints []int, k int) int {
	//从两边选卡片，选 k 张，卡片总数量为 n 张，即有 n - k 张不被选择。
	//
	//所有卡片总和 totalSum 固定，要使选择的 k 张的总和最大，反过来就是要让不被选择的 n - k 张总和最小。
	//
	//原问题等价为：从 cardPoints 中找长度为 n - k 的连续段，使其总和最小。
	//
	//具体的，用变量 totalSum 代指 cardPoints 总和，windowSum 代表长度固定为 n - k 的当前窗口总和，minv 代表所有长度为 n - k 的窗口中总和最小的值。
	//
	//起始先将滑动窗口压满，取得第一个滑动窗口的目标值 windowSum（同时更新为 minv），随后往后继续处理 cardPoints，每往前滑动一位，需要删除一个和添加一个元素，并不断更新 minv，最终 totalSum - minv 即是答案
	n := len(cardPoints)
	windowLength := n - k
	totalSum := 0
	for i := 0; i < n; i++ {
		totalSum += cardPoints[i]
	}

	//初始化窗口
	windowSum := 0
	for i := 0; i < windowLength; i++ {
		windowSum += cardPoints[i]
	}

	minValue := windowSum
	for i := windowLength; i < n; i++ {
		//不断移动窗口 计算窗口内的和
		windowSum = windowSum + cardPoints[i] - cardPoints[i-windowLength]
		//取窗口最小值
		minValue = util.Min(minValue, windowSum)
	}
	return totalSum - minValue
}
