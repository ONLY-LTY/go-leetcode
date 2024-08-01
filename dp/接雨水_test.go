package dp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrap(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	assert.Equal(t, 6, trap(height))

	height = []int{4, 2, 0, 3, 2, 5}
	assert.Equal(t, 9, trap(height))
}
