package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	actual := []int{5, 6, 7, 1, 2, 3, 4}
	rotate(nums, 3)
	assert.Equal(t, actual, nums)
}
