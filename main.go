package main

import (
	"fmt"
)

type CardValue int

const (
	// Diamond
	TwoDiamond CardValue = iota + 0x102
	ThreeDiamond
	FourDiamond
	FiveDiamond
	SixDiamond
	SevenDiamond
	EightDiamond
	NineDiamond
	TenDiamond
	JackDiamond
	QueenDiamond
	KingDiamond
	AceDiamond

	// Club
	TwoClub CardValue = iota + 0x202
	ThreeClub
	FourClub
	FiveClub
	SixClub
	SevenClub
	EightClub
	NineClub
	TenClub
	JackClub
	QueenClub
	KingClub
	AceClub

	// Heart
	TwoHeart CardValue = iota + 0x302
	ThreeHeart
	FourHeart
	FiveHeart
	SixHeart
	SevenHeart
	EightHeart
	NineHeart
	TenHeart
	JackHeart
	QueenHeart
	KingHeart
	AceHeart

	// Spade
	TwoSpade CardValue = iota + 0x402
	ThreeSpade
	FourSpade
	FiveSpade
	SixSpade
	SevenSpade
	EightSpade
	NineSpade
	TenSpade
	JackSpade
	QueenSpade
	KingSpade
	AceSpade
)

// print hello world
func one() string {
	return "hello world"
}

//sort the array in ascending order
func two(input []int) string {
	//sort the array in ascending order
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] > input[j] {
				input[i], input[j] = input[j], input[i]
			}
		}
	}
	//return the sorted array as a string
	return fmt.Sprint(input)
}

func main() {
	fmt.Println(one())
	fmt.Println(two([]int{10, 3, 4, 2, 7}))

	fmt.Println(AceClub, TwoDiamond, ThreeDiamond, FourDiamond, FiveDiamond)
	fmt.Println(TwoClub, ThreeClub, FourClub, FiveClub)
	fmt.Println(TwoHeart, ThreeHeart, FourHeart, FiveHeart)
	fmt.Println(TwoSpade, ThreeSpade, FourSpade, FiveSpade)
}
