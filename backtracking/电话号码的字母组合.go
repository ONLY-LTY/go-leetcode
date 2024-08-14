package backtracking

var dict = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}

// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
// 示例 1：
//
// 输入：digits = "23"
// 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
// 示例 2：
//
// 输入：digits = ""
// 输出：[]
// 示例 3：
//
// 输入：digits = "2"
// 输出：["a","b","c"]
func letterCombinations(digits string) []string {
	var res []string
	if len(digits) == 0 {
		return res
	}
	dfsLetterCombinations(digits, 0, "", &res)
	return res
}

func dfsLetterCombinations(digits string, level int, path string, res *[]string) {
	if level == len(digits) {
		*res = append(*res, path)
		return
	}
	number := string(digits[level])
	letter := dict[number]
	for i := 0; i < len(letter); i++ {
		//遍历
		dfsLetterCombinations(digits, level+1, path+letter[i], res)
	}
}
