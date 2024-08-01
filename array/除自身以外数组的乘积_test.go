package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	assert.Equal(t, []int{24, 12, 8, 6}, productExceptSelf(nums))
	assert.Equal(t, []int{24, 12, 8, 6}, productExceptSelf2(nums))
	assert.Equal(t, []int{24, 12, 8, 6}, productExceptSelf3(nums))

	nums = []int{-1, 1, 0, -3, 3}
	assert.Equal(t, []int{0, 0, 9, 0, 0}, productExceptSelf(nums))
	assert.Equal(t, []int{0, 0, 9, 0, 0}, productExceptSelf2(nums))
	assert.Equal(t, []int{0, 0, 9, 0, 0}, productExceptSelf3(nums))

}
