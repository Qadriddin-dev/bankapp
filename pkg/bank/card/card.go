package card

import (
	"bank/pkg/bank/types"
)

func Total(cards []types.Card) types.Money {
	var sum types.Money = 0
	for _, card := range cards {
		sum += card.Balance

		if card.Active == false {
			break
		}
		if sum < 0 {
			break
		}

	}

	return sum
}
