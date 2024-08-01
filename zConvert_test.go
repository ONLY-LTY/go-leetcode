package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvert(t *testing.T) {
	assert.Equal(t, "PAHNAPLSIIGYIR", convert("PAYPALISHIRING", 3))
	assert.Equal(t, "PAHNAPLSIIGYIR", convert2("PAYPALISHIRING", 3))

	assert.Equal(t, "PINALSIGYAHRPI", convert("PAYPALISHIRING", 4))
	assert.Equal(t, "PINALSIGYAHRPI", convert2("PAYPALISHIRING", 4))
}
