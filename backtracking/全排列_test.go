package backtracking

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	nums := []int{5, 4, 6, 2}
	res := permute(nums)
	fmt.Println(res)
	nums = []int{0, 1}
	res = permute(nums)
	fmt.Println(res)
}
