package string

func reverseWords(s string) string {
	b := []byte(s)
	slow := removeBlank(b)
	b = b[0:slow]

	reverseByte(b, 0, len(b)-1)
	wordStart := 0
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			reverseByte(b, wordStart, i-1)
			wordStart = i + 1
		}
	}
	return string(b)
}

// removeBlank 删除单词之间多余的空格
func removeBlank(s []byte) int {
	slow := 0
	for fast := 0; fast < len(s); fast++ {
		//找到符合条件的单词起始位置
		if s[fast] != ' ' {
			//单词前面是否需要加一个空格 首单词前面不需要加空格 其他情况都需要加
			if slow != 0 {
				s[slow] = ' '
				slow++
			}
			//从单词起始位置开始复制单词
			for fast < len(s) && s[fast] != ' ' {
				s[slow] = s[fast]
				fast++
				slow++
			}
		}
	}
	return slow
}

// reverseByte 反转字节数组 支持指定区间
func reverseByte(s []byte, begin int, end int) {
	for begin < end {
		s[begin], s[end] = s[end], s[begin]
		begin++
		end--
	}
}
