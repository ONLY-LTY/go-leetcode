package dp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	res := longestCommonSubsequence("abcde", "ace")
	assert.Equal(t, 3, res)

	res = longestCommonSubsequence("abc", "def")
	assert.Equal(t, 0, res)
}
