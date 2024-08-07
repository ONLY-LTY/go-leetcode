package string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAyAtoi(t *testing.T) {
	s := "4193 with words"
	assert.Equal(t, 4193, myAtoi(s))
	s = "   -42"
	assert.Equal(t, -42, myAtoi(s))
	s = "42"
	assert.Equal(t, 42, myAtoi(s))
}
