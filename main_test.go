package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOne(t *testing.T) {
	res := one()
	assert.Equal(t, res, "hello world")
}
