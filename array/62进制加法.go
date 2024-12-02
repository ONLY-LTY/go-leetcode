package array

import (
	"strings"
)

const Base62Char = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func add62Number(s1 string, s2 string) string {
	s1Index := len(s1) - 1
	s2Index := len(s2) - 1
	carry := 0
	var result []byte
	for s1Index >= 0 || s2Index >= 0 || carry > 0 {
		a := 0
		if s1Index >= 0 {
			a = strings.Index(Base62Char, string(s1[s1Index]))
		}
		b := 0
		if s2Index >= 0 {
			b = strings.Index(Base62Char, string(s2[s2Index]))
		}
		sum := a + b + carry
		carry = sum / 62
		result = append(result, Base62Char[sum%62])
		s1Index--
		s2Index--
	}
	reverse(result)
	return string(result)
}

func reverse(array []byte) {
	start := 0
	end := len(array) - 1
	for start < end {
		array[start], array[end] = array[end], array[start]
		start++
		end--
	}
}
