package backtracking

import (
	"fmt"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	res := generateParenthesis(3)
	fmt.Println(res)
}
