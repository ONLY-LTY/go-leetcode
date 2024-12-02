package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd62Number(t *testing.T) {
	res := add62Number("Z", "1")
	assert.Equal(t, "10", res)
	res = add62Number("1a", "2")
	assert.Equal(t, "1c", res)
}
