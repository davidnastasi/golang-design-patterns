package factorymethod

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float32) string
}

type PaymentType int

const (
	Cash     PaymentType = iota + 1
	DebitCard
	CreditCard
)

type CashPM struct{}
type DebitCardPM struct{}
type CreditCardPM struct {}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using debit card\n", amount)
}

func (d *CreditCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using new credit card implementation\n", amount)
}

func GetPaymentMethod(paymentType PaymentType) (PaymentMethod, error) {
	switch paymentType {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	case CreditCard:
		return new(CreditCardPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recognized\n", paymentType))
	}
}
