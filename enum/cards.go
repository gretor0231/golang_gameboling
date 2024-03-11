package enum

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
