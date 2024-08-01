package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	assert.Equal(t, 6, maxSubArray(nums))

	nums = []int{5, 4, -1, 7, 8}
	assert.Equal(t, 23, maxSubArray(nums))
}
