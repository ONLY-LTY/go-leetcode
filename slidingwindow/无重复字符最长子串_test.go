package slidingwindow

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "abcabcbb"
	assert.Equal(t, 3, lengthOfLongestSubstring(s))

	s = "bbbbb"
	assert.Equal(t, 1, lengthOfLongestSubstring(s))

	s = "pwwkew"
	assert.Equal(t, 3, lengthOfLongestSubstring(s))
}
