package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerge(t *testing.T) {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	actual := merge(intervals)
	expected := [][]int{{1, 6}, {8, 10}, {15, 18}}
	assert.Equal(t, expected, actual)

	intervals = [][]int{{1, 4}, {2, 3}}
	actual = merge(intervals)
	expected = [][]int{{1, 4}}
	assert.Equal(t, expected, actual)
}
