package bchain

import (
	"fmt"
	"time"
)

type (
	chain struct {
		blocks              blocks
		pendingTransactions transactions
		miningReward        money
	}
)

func NewChain() *chain {
	genesisBlock := NewBlock(
		time.Now(),
		make([]transaction, 10),
		[32]byte{},
	)
	genesisBlock.mining()

	return &chain{
		miningReward: money(12.5),
		blocks:       []block{genesisBlock},
	}
}

func (c chain) String() string {
	return fmt.Sprintf("blocks: %s", c.blocks)
}

func (c *chain) AddTransaction(tx transaction) {
	c.pendingTransactions = append(c.pendingTransactions, tx)
}

func (c *chain) Mine(rewardAddress walletAddress) error {
	block := NewBlock(
		time.Now(),
		c.pendingTransactions,
		c.blocks.last().hash,
	)
	block.mining()
	c.blocks = append(c.blocks, block)

	c.pendingTransactions = make([]transaction, 0)

	t, err := NewTransaction("", string(rewardAddress), c.miningReward)
	if err != nil {
		return err
	}
	c.AddTransaction(t)

	return nil
}

func (c chain) Validate() error {
	return c.blocks.Validate()
}

func (c chain) GetBalanceOfAddress(a string) money {
	var balance money
	for _, block := range c.blocks {
		for _, transaction := range block.committedTransactions {
			if transaction.sender == walletAddress(a) {
				balance -= transaction.amount
			}
			if transaction.recipient == walletAddress(a) {
				balance += transaction.amount
			}
		}
	}

	return balance
}
