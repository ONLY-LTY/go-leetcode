package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProduct(t *testing.T) {
	nums := []int{2, 3, -2, 4}
	assert.Equal(t, 6, maxProduct(nums))

	nums = []int{-2, 0, -1}
	//assert.Equal(t, 0, maxProduct(nums))

	nums = []int{-2, 3, -4}
	//fmt.Println(maxProduct(nums))
}
