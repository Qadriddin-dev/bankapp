package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleTotal() {
	cards := []types.Card{
		types.Card{
			Balance: -10_000_00,
			Active:  false,
		},
		{
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Balance: 15_000_00,
			Active:  false,
		},
	}

	fmt.Println(Total(cards))
	//Output:-1000000

}
