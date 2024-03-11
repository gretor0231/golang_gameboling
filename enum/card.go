package enum

type Suit int

// Card represents a playing card with a suit and number
type Card struct {
	Suit   Suit
	Number int
}

func ConvertInputToCard(inputVal int) Card {
	card := Card{
		Suit:   Suit(inputVal / 0x100),
		Number: inputVal % 0x100,
	}
	if card.Number == 14 {
		card.Number = 1
	}
	return card

}

const (
	// Diamond
	TwoDiamond   int = 0x102
	ThreeDiamond     = 0x103
	FourDiamond      = 0x104
	FiveDiamond      = 0x105
	SixDiamond       = 0x106
	SevenDiamond     = 0x107
	EightDiamond     = 0x108
	NineDiamond      = 0x109
	TenDiamond       = 0x10a
	JackDiamond      = 0x10b
	QueenDiamond     = 0x10c
	KingDiamond      = 0x10d
	AceDiamond       = 0x10e

	// Club
	TwoClub   int = 0x202
	ThreeClub     = 0x203
	FourClub      = 0x204
	FiveClub      = 0x205
	SixClub       = 0x206
	SevenClub     = 0x207
	EightClub     = 0x208
	NineClub      = 0x209
	TenClub       = 0x20a
	JackClub      = 0x20b
	QueenClub     = 0x20c
	KingClub      = 0x20d
	AceClub       = 0x20e

	// Heart
	TwoHeart   int = 0x302
	ThreeHeart     = 0x303
	FourHeart      = 0x304
	FiveHeart      = 0x305
	SixHeart       = 0x306
	SevenHeart     = 0x307
	EightHeart     = 0x308
	NineHeart      = 0x309
	TenHeart       = 0x30a
	JackHeart      = 0x30b
	QueenHeart     = 0x30c
	KingHeart      = 0x30d
	AceHeart       = 0x30e

	// Spade
	TwoSpade   int = 0x402
	ThreeSpade     = 0x403
	FourSpade      = 0x404
	FiveSpade      = 0x405
	SixSpade       = 0x406
	SevenSpade     = 0x407
	EightSpade     = 0x408
	NineSpade      = 0x409
	TenSpade       = 0x40a
	JackSpade      = 0x40b
	QueenSpade     = 0x40c
	KingSpade      = 0x40d
	AceSpade       = 0x40e
)
