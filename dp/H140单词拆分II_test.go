package dp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWordBreakII(t *testing.T) {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	res := wordBreakII(s, wordDict)
	assert.Equal(t, []string{"cat sand dog", "cats and dog"}, res)
}
