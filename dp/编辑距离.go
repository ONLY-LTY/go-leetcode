package dp

import "go-leetcode/util"

// 给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。
//
// 你可以对一个单词进行如下三种操作：
//
// 插入一个字符
// 删除一个字符
// 替换一个字符
//
// 示例 1：
//
// 输入：word1 = "horse", word2 = "ros"
// 输出：3
// 解释：
// horse -> rorse (将 'h' 替换为 'r')
// rorse -> rose (删除 'r')
// rose -> ros (删除 'e')
// 示例 2：
//
// 输入：word1 = "intention", word2 = "execution"
// 输出：5
// 解释：
// intention -> inention (删除 't')
// inention -> enention (将 'i' 替换为 'e')
// enention -> exention (将 'n' 替换为 'x')
// exention -> exection (将 'n' 替换为 'c')
// exection -> execution (插入 'u')
//
// 提示：
//
// 0 <= word1.length, word2.length <= 500
// word1 和 word2 由小写英文字母组成
func minDistance(word1 string, word2 string) int {
	//字符位置从1开始 所以长度加一
	w1Len := len(word1) + 1
	w2Len := len(word2) + 1
	//dp[i][j] 表示word1的前i个字符 变换成word2的前j个字符最小步骤
	dp := make([][]int, w1Len)
	for i := 0; i < w1Len; i++ {
		dp[i] = make([]int, w2Len)
	}
	//初始化第一列数据 含义就是word1的前i个字符变换成空字符串的最小步骤 每次都是删除操作
	for row := 0; row < w1Len; row++ {
		dp[row][0] = row
	}
	//初始化第一行数据 含义就是空字符串变换成word2的前j个字符的最小步骤 每次都是插入操作
	for cell := 0; cell < w2Len; cell++ {
		dp[0][cell] = cell
	}
	for i := 1; i < w1Len; i++ {
		for j := 1; j < w2Len; j++ {
			//如果word1的第i个字符 和 Word2的j个字符相等
			//dp[i][j]=word1的前i-1个字符变换成word2前j-1个字符
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				//1.插入
				//如果最后一步是插入，那么需要将 string1 的前 i 个字符变换为 string2 的前 j - 1 个字符，
				//最后一步插入 string2的第j个字符即可，这样至少需要 d[i][j - 1] + 1 步，
				//因为将 string1 的前 i 个字符变换为 string2 的前 j - 1 个字符至少需要 d[i][j - 1] 步，
				//加上最后一步的插入操作，所以总共至少需要 d[i][j - 1] + 1 步
				insert := dp[i][j-1] + 1

				//2.删除
				//如果最后一步是删除，那么需要将 string1 的前 i - 1 个字符变换为 string2 的前 j 个字符，
				//最后一步删除 string1的第i个字符，至少需要 d[i - 1][j] + 1 步
				del := dp[i-1][j] + 1

				//修改
				//如果最后一步是修改，那么需要将 string1 的前 i - 1 个字符变换为 string2 的前 j - 1 字符，
				//最后一步将 string1第i个字符 修改为 string2的第j字符，至少需要 d[i - 1][j - 1] + 1 步
				update := dp[i-1][j-1] + 1

				//取变换中最小的
				dp[i][j] = util.Min(util.Min(insert, del), update)
			}
		}
	}
	return dp[w1Len-1][w2Len-1]
}
