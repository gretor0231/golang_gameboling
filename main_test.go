package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOne(t *testing.T) {
	res := one()
	assert.Equal(t, res, "hello world")
}

func TestTwo(t *testing.T) {

	res := two([]int{10, 3, 4, 2, 7})
	assert.Equal(t, res, "[2 3 4 7 10]")
}
