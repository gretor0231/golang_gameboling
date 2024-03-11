package enum

type Cards []Card

func ConvertToCards(inputVal []int) Cards {
	var cards []Card
	for _, val := range inputVal {
		cards = append(cards, Card{
			Suit:   Suit(val / 0x100),
			Number: val % 0x100,
		})
	}
	for k, v := range cards {
		if v.Number == 14 {
			cards[k].Number = 1
		}
	}

	return cards
}

func (c Cards) Len() int {
	return len(c)
}

// Swap swaps the elements with the given indices
func (c Cards) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// Less compares two card number
func (c Cards) Less(i, j int) bool {
	return c[i].Number < c[j].Number
}
