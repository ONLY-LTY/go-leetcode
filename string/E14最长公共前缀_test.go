package string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonPrefix(test *testing.T) {
	prefix := longestCommonPrefix([]string{"flower", "fl", "flight"})
	assert.Equal(test, "fl", prefix)
}
