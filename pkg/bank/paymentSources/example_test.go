package paymentSources

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSources() {
	cards := []types.Card{
		{
			PAN: "5058 xxxx xxxx 8888",
			Active: true,
			Balance: 10_000,
		},
		{
			PAN: "5058 xxxx xxxx 4444",
			Active: true,
			Balance: 15_000,
		},
		{
			PAN: "5058 xxxx xxxx 3333",
			Active: false,
			Balance: 15_000,
		},
		{
			PAN: "5058 xxxx xxxx 2222",
			Active: true,
			Balance: 0,
		},
	}

	paymentSources := PaymentSources(cards)

	for _, paymentSource := range paymentSources {
		fmt.Println(paymentSource.Number)
	}

	// Output:
	// 5058 xxxx xxxx 8888
	// 5058 xxxx xxxx 4444
}
