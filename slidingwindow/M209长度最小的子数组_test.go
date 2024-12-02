package slidingwindow

import (
	"fmt"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	res := minSubArrayLen(4, []int{1, 4, 4})
	fmt.Println(res)
}
