package dp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	assert.Equal(t, 5, maxProfit(prices))
	assert.Equal(t, 7, maxProfit2(prices))
	prices = []int{1, 2, 3, 4, 5}
	assert.Equal(t, 4, maxProfit(prices))
	assert.Equal(t, 4, maxProfit2(prices))
	prices = []int{7, 6, 4, 3, 1}
	assert.Equal(t, 0, maxProfit(prices))
	assert.Equal(t, 0, maxProfit2(prices))
}
