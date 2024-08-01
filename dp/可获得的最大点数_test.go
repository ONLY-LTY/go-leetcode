package dp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxScore(t *testing.T) {
	cardPoints := []int{1, 2, 3, 4, 5, 6, 1}
	assert.Equal(t, 12, maxScore(cardPoints, 3))

	cardPoints = []int{2, 2, 2}
	assert.Equal(t, 4, maxScore(cardPoints, 2))

	cardPoints = []int{9, 7, 7, 9, 7, 7, 9}
	assert.Equal(t, 55, maxScore(cardPoints, 7))

	cardPoints = []int{1, 1000, 1}
	assert.Equal(t, 1, maxScore(cardPoints, 1))

	cardPoints = []int{1, 79, 80, 1, 1, 1, 200, 1}
	assert.Equal(t, 202, maxScore(cardPoints, 3))
}
