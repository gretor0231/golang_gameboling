package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOne(t *testing.T) {
	res := One()
	assert.Equal(t, res, "hello world")
}
