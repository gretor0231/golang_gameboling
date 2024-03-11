package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwo(t *testing.T) {

	res := Two([]int{10, 3, 4, 2, 7})
	assert.Equal(t, res, "[2 3 4 7 10]")
}
