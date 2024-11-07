package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOutArea(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	res := numIslands(grid)
	assert.Equal(t, 1, res)

	grid = [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	res = numIslands(grid)
	assert.Equal(t, 3, res)
}
