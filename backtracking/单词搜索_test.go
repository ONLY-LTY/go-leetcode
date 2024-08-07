package backtracking

import (
	"testing"
)

func Test(t *testing.T) {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	if !exist(board, "SEE") {
		t.Fail()
	}
}
