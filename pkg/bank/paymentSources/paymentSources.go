package paymentSources

import "bank/pkg/bank/types"

func PaymentSources(cards []types.Card) []types.PaymentSource {
	var paymentSource []types.PaymentSource
	for _, card := range cards {
		if card.Balance > 0 && card.Active {
			payment := types.PaymentSource{
				Type:    "card",
				Number:  string(card.PAN),
				Balance: card.Balance,
			}
			paymentSource = append(paymentSource, payment)
		}
	}
	return paymentSource
}