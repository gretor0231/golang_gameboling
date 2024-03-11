package functions

import (
	"github.com/stretchr/testify/assert"
	"golang_gameboling/enum"
	"testing"
)

func TestThree(t *testing.T) {
	res := three([]int{enum.AceClub, enum.AceHeart, enum.AceDiamond, enum.AceSpade, enum.TenClub})
	assert.Equal(t, res, enum.Result(3))
}
