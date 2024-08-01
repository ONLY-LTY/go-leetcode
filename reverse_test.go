package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	assert.Equal(t, 321, reverse(123))
	assert.Equal(t, -321, reverse(-123))
	assert.Equal(t, 21, reverse(120))
	assert.Equal(t, 0, reverse(0))

	assert.Equal(t, true, isPalindrome(121))
	assert.Equal(t, false, isPalindrome(-121))
}
