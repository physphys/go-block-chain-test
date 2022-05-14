package bchain

import "fmt"

type (
	address string
	money   float64

	transaction struct {
		sender    address
		recipient address
		amount    money
	}
)

func NewTransaction(sender string, recipient string, amount float64) transaction {
	t := transaction{
		sender:    address(sender),
		recipient: address(recipient),
		amount:    money(amount),
	}

	return t
}

func (tx transaction) String() string {
	return fmt.Sprintf(
		"{ %s sent %v => %s }",
		tx.sender,
		tx.amount,
		tx.recipient,
	)
}
