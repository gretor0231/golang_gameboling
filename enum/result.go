package enum

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
