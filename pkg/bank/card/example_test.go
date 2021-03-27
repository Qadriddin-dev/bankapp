package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleTotal() {
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
		},
		{
			Balance: 2_000_00,
			Active:  true,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  false,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: -1_000_00,
			Active:  true,
		},
	}))
	//Output:
	//100000
	//300000
	//0
	//0

}

func ExamplePaymentSource() {
	cards := []types.Card{
		{
			PAN:     "5058 xxxx xxxx 8888",
			Balance: 30_000_03,
			Active:  true,
		},
		{
			PAN:     "5058 xxxx xxxx 9999",
			Balance: -50_000_05,
			Active:  true,
		},
		{
			PAN:     "5058 xxxx xxxx 2020",
			Balance: 40_000_04,
			Active:  false,
		},
	}
	payment := PaymentSource(cards)
	for _, payments := range payment {
		fmt.Println(payments.Number)
	}

	//Output:
	//5058 xxxx xxxx 8888

}
