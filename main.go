package main

import (
	"fmt"
	"sort"
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

type Result int

const (
	Royalflush    Result = 1  //A poker hand with the A, K, Q, K and 10 all in the same suit.
	Straightflush        = 2  //A poker hand with consecutive cards in the same suit. For example: [0x109,0x10a,0x10b,0x10c,0x10d],[0x10e,0x102,0x103,0x104,0x105]
	Fourofakind          = 3  // A poker hand with 4 cards with the same rank plus 1 arbitrary card.
	Fullhouse            = 4  // A poker hand with 3 of a kind and a pair.
	Flush                = 5  // A poker hand with all 5 cards in the same suit.
	Straight             = 6  // A poker hand with 5 consecutive cards (regardless of suit).
	Threeofakind         = 7  // poker hand with 3 cards with the same rank plus 2 cards in different rank.
	Twopair              = 8  // A poker hand with two pairs of similar-ranking cards plus 1 arbitrary card.
	Onepair              = 9  // poker hand with 2 cards in same rank plus 3 cards in different rank.
	Highcard             = 10 // A poker hand that do not make any of the poker hands given above.
	NotEnough            = -1 // A poker hand that do not have five cards.
)

type Cards []CardValue

func (c Cards) Len() int {
	return len(c)
}

// Swap swaps the elements with the given indices
func (c Cards) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// Less compares the age of two people
func (c Cards) Less(i, j int) bool {
	return c[i] < c[j]
}

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

//calculate hand type
func three(input Cards) Result {
	// input cards not len 5 and not belong CardValue
	sort.Sort(input) // Cards in assending order
	if len(input) != 5 {
		return NotEnough
	} else {
		return Highcard
	}
}

func main() {
	fmt.Println(one())
	fmt.Println(two([]int{10, 3, 4, 2, 7}))
	fmt.Println(three(Cards{AceClub, TwoDiamond, ThreeDiamond, FourDiamond, FiveDiamond}))

}
