package main

import (
	"fmt"
	"golang_gameboling/enum"
)

func main() {
	//fmt.Println(functions.One())
	//fmt.Println(functions.Two([]int{10, 3, 4, 2, 7}))
	//fmt.Println(three(Cards{AceClub, TwoDiamond, ThreeDiamond, FourDiamond, FiveDiamond}))
	fmt.Println(enum.ConvertInputToCard(enum.AceClub))
}
