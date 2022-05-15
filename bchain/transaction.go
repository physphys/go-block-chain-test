package bchain

import "fmt"

type (
	walletAddress string

	transactions []transaction

	transaction struct {
		sender    walletAddress
		recipient walletAddress
		amount    money
	}
)

func NewTransaction(sender string, recipient string, amount money) (transaction, error) {
	if err := amount.validate(); err != nil {
		return transaction{}, err
	}

	t := transaction{
		sender:    walletAddress(sender),
		recipient: walletAddress(recipient),
		amount:    amount,
	}

	return t, nil
}

func (tx transaction) String() string {
	return fmt.Sprintf(
		"{ %s sent %v => %s }",
		tx.sender,
		tx.amount,
		tx.recipient,
	)
}
