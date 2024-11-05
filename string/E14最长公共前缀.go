package string

// 编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。
//
// 示例 1：
//
// 输入：strs = ["flower","flow","flight"]
// 输出："fl"
// 示例 2：
//
// 输入：strs = ["dog","racecar","car"]
// 输出：""
// 解释：输入不存在公共前缀。
//
// 提示：
//
// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] 仅由小写英文字母组成
func longestCommonPrefix(str []string) string {
	if len(str) == 0 {
		return ""
	}
	//用第一个字符串
	for i := 0; i < len(str[0]); i++ {
		//和后面的每个字符串做比较
		for j := 1; j < len(str); j++ {
			//如果前i个字符 和 第j个字符串一样长 直接可以返回了
			if i == len(str[j]) || str[j][i] != str[0][i] {
				return str[0][:i]
			}
		}
	}
	return str[0]
}
