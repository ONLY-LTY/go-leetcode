package slidingwindow

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaximumSubarraySum(t *testing.T) {
	nums := []int{1, 5, 4, 2, 9, 9, 9}
	assert.Equal(t, int64(15), maximumSubarraySum(nums, 3))
}
