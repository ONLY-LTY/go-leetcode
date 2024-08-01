package string

func findAnagrams(s string, p string) []int {
	var result []int
	if len(s) < len(p) {
		return result
	}
	var ch [26]int
	//统计P字符串的个数
	for i := 0; i < len(p); i++ {
		ch[p[i]-'a']++
	}
	//初始化窗口
	windowStart := 0
	windowEnd := 0
	reset := len(p)
	for windowEnd < len(p) {
		ch[s[windowEnd]-'a']--
		//如果是非P字符串中的 这里肯定小于0
		if ch[s[windowEnd]-'a'] >= 0 {
			reset--
		}
		windowEnd++
	}
	if reset == 0 {
		result = append(result, windowStart)
	}
	for windowEnd < len(s) {
		//左边元素出窗口
		if ch[s[windowStart]-'a'] >= 0 {
			reset++
		}
		ch[s[windowStart]-'a']++
		windowStart++

		//右边元素进窗口
		ch[s[windowEnd]-'a']--
		if ch[s[windowEnd]-'a'] >= 0 {
			reset--
		}
		windowEnd++

		//符合条件 记录
		if reset == 0 {
			result = append(result, windowStart)
		}
	}
	return result
}
