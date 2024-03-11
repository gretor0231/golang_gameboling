package functions

import (
	"golang_gameboling/enum"
	"sort"
)

//calculate hand type
func three(input enum.Cards) enum.Result {
	// input cards not len 5 and not belong CardValue
	sort.Sort(input) // Cards in assending order
	if len(input) != 5 {
		return enum.NotEnough
	} else {
		return enum.Highcard
	}
}
