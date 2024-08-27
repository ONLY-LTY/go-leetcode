package array

import (
	"fmt"
	"testing"
)

func TestLargestNumber(t *testing.T) {
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
	fmt.Println(largestNumber([]int{10, 2}))
}
