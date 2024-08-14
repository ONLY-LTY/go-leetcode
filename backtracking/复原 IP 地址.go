package backtracking

import "strings"

// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
//
// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
// 给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。
// 你 不能 重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。
//
// 示例 1：
//
// 输入：s = "25525511135"
// 输出：["255.255.11.135","255.255.111.35"]
// 示例 2：
//
// 输入：s = "0000"
// 输出：["0.0.0.0"]
// 示例 3：
//
// 输入：s = "101023"
// 输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
func restoreIpAddresses(s string) []string {
	var res []string
	dfsRestoreIpAddresses(s, 0, []string{}, &res)
	return res
}

func dfsRestoreIpAddresses(s string, start int, path []string, res *[]string) {
	//IP的4段
	if len(path) == 4 {
		// 需要判断 是否遍历到字符窜的末尾
		if start == len(s) {
			*res = append(*res, strings.Join(path, "."))
		}
		return
	}

	for i := start; i < len(s); i++ {
		if isValid(s, start, i) {
			dfsRestoreIpAddresses(s, i+1, append(path, s[start:i+1]), res)
		}
	}
}

func isValid(s string, start int, end int) bool {
	//理论上不会出现
	if start > end {
		return false
	}
	//0开头非法
	if s[start] == '0' && start != end {
		return false
	}
	num := 0
	for i := start; i <= end; i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
		num = num*10 + int(s[i]-'0')
		if num > 255 {
			return false
		}
	}
	return true
}
