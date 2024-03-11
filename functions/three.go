package functions

import (
	"golang_gameboling/enum"
	"sort"
)

//calculate hand type

//Complete the function f that computes the ranking of a poker hand. The input is consisted of five integers that represent different poker cards.
//Belows are the correspoding relationship between each interger(for better reading, the integer is shown in hexadecimal) and poker card:
//
//0x102,0x103,0x104,0x105,0x106,0x107,0x108,0x109,0x10a,0x10b,0x10c,0x10d,0x10e correspond to 2,3,4,5,6,7,8,9,10,J,Q,K,A of diamond.
//0x202,0x203,0x204,0x205,0x206,0x207,0x208,0x209,0x20a,0x20b,0x20c,0x20d,0x20e correspond to 2,3,4,5,6,7,8,9,10,J,Q,K,A of club.
//0x302,0x303,0x304,0x305,0x306,0x307,0x308,0x309,0x30a,0x30b,0x30c,0x30d,0x30e correspond to 2,3,4,5,6,7,8,9,10,J,Q,K,A of heart.
//0x402,0x403,0x404,0x405,0x406,0x407,0x408,0x409,0x40a,0x40b,0x40c,0x40d,0x40e correspond to 2,3,4,5,6,7,8,9,10,J,Q,K,A of spade.
//
//And here are the definition of return value:
//
//1. Royal flush: A poker hand with the A, K, Q, K and 10 all in the same suit.
//2. Straight flush: A poker hand with consecutive cards in the same suit. For example: [0x109,0x10a,0x10b,0x10c,0x10d],[0x10e,0x102,0x103,0x104,0x105]
//3. Four of a kind: A poker hand with 4 cards with the same rank plus 1 arbitrary card.
//4. Full house: A poker hand with 3 of a kind and a pair.
//5. Flush: A poker hand with all 5 cards in the same suit.
//6. Straight: A poker hand with 5 consecutive cards (regardless of suit).
//7. Three of a kind: A poker hand with 3 cards with the same rank plus 2 cards in different rank.
//8. Two pair: A poker hand with two pairs of similar-ranking cards plus 1 arbitrary card.
//9. One pair: A poker hand with 2 cards in same rank plus 3 cards in different rank.
//10. High card: A poker hand that do not make any of the poker hands given above.
//
func three(input []int) enum.Result {
	// input cards not len 5 and not belong Card

	cards := enum.ConvertToCards(input)
	sort.Sort(cards) // sort cards' number by ascending order
	if len(input) != 5 {
		return enum.NotEnough
	} else if (input[0] == enum.TenDiamond && input[1] == enum.JackDiamond && input[2] == enum.QueenDiamond && input[3] == enum.KingDiamond && input[4] == enum.AceDiamond) ||
		(input[0] == enum.TenClub && input[1] == enum.JackClub && input[2] == enum.QueenClub && input[3] == enum.KingClub && input[4] == enum.AceClub) ||
		(input[0] == enum.TenHeart && input[1] == enum.JackHeart && input[2] == enum.QueenHeart && input[3] == enum.KingHeart && input[4] == enum.AceHeart) ||
		(input[0] == enum.TenSpade && input[1] == enum.JackSpade && input[2] == enum.QueenSpade && input[3] == enum.KingSpade && input[4] == enum.AceSpade) {
		return enum.Royalflush
	} else if (cards[0].Suit == cards[1].Suit && cards[1].Suit == cards[2].Suit && cards[2].Suit == cards[3].Suit && cards[3].Suit == cards[4].Suit) &&
		cards[0].Number == cards[1].Number-1 && cards[1].Number == cards[2].Number-1 && cards[2].Number == cards[3].Number-1 && cards[3].Number == cards[4].Number-1 {
		return enum.Straightflush
	} else if (cards[0].Number == cards[1].Number && cards[1].Number == cards[2].Number && cards[2].Number == cards[3].Number) ||
		(cards[1].Number == cards[2].Number && cards[2].Number == cards[3].Number && cards[3].Number == cards[4].Number) {
		return enum.Fourofakind
	} else if (cards[0].Number == cards[1].Number && cards[1].Number == cards[2].Number && cards[3].Number == cards[4].Number) ||
		(cards[0].Number == cards[1].Number && cards[2].Number == cards[3].Number && cards[3].Number == cards[4].Number) {
		return enum.Fullhouse
	} else if cards[0].Suit == cards[1].Suit && cards[1].Suit == cards[2].Suit && cards[2].Suit == cards[3].Suit && cards[3].Suit == cards[4].Suit {
		return enum.Flush
	} else if (cards[0].Number == cards[1].Number-1 && cards[1].Number == cards[2].Number-1 && cards[2].Number == cards[3].Number-1 && cards[3].Number == cards[4].Number-1) ||
		(cards[0].Number == 10 && cards[1].Number == 11 && cards[2].Number == 12 && cards[3].Number == 13 && cards[4].Number == 1) {
		return enum.Straight
	} else if (cards[0].Number == cards[1].Number && cards[1].Number == cards[2].Number) ||
		(cards[1].Number == cards[2].Number && cards[2].Number == cards[3].Number) ||
		(cards[2].Number == cards[3].Number && cards[3].Number == cards[4].Number) {
		return enum.Threeofakind
	} else if (cards[0].Number == cards[1].Number && cards[2].Number == cards[3].Number) || (cards[1].Number == cards[2].Number && cards[3].Number == cards[4].Number) {
		return enum.Twopair
	} else if (cards[0].Number == cards[1].Number) || (cards[1].Number == cards[2].Number) || (cards[2].Number == cards[3].Number) || (cards[3].Number == cards[4].Number) {
		return enum.Onepair
	} else {
		return enum.Highcard
	}
}
