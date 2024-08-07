package backtracking

// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
// 示例 1：
//
// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]
// 示例 2：
//
// 输入：n = 1
// 输出：["()"]
func generateParenthesis(n int) []string {
	var res []string
	dfsParentheses(n, n, "", &res)
	return res
}

func dfsParentheses(left int, right int, str string, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, str)
	}
	if left > 0 {
		dfsParentheses(left-1, right, str+"(", res)
	}
	if right > 0 && right > left {
		dfsParentheses(left, right-1, str+")", res)
	}
}
