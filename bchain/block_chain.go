package bchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"time"
)

type (
	BlockChain struct {
		blocks              []block
		pendingTransactions []transaction
		miningReward        money
	}
)

func NewBlockChain() *BlockChain {
	genesisBlock := NewBlock(
		time.Now(),
		make([]transaction, 10),
		sha256.Sum256([]byte(fmt.Sprintf("00"))),
	)
	genesisBlock.mining()

	return &BlockChain{
		miningReward: money(12.5),
		blocks:       []block{genesisBlock},
	}
}

func (bc BlockChain) String() string {
	return fmt.Sprintf("blocks: %s", bc.blocks)
}

func (bc BlockChain) getLastBlock() block {
	return bc.blocks[len(bc.blocks)-1]
}

func (bc *BlockChain) AddTransaction(tx transaction) {
	bc.pendingTransactions = append(bc.pendingTransactions, tx)
}

func (bc *BlockChain) Mine(rewardAddress address) {
	block := NewBlock(
		time.Now(),
		bc.pendingTransactions,
		bc.getLastBlock().hash,
	)
	block.mining()
	bc.blocks = append(bc.blocks, block)

	bc.pendingTransactions = make([]transaction, 0)

	bc.AddTransaction(
		NewTransaction(
			"",
			string(rewardAddress),
			float64(bc.miningReward),
		),
	)
}

func (bc BlockChain) IsValid() (bool, error) {
	for i := len(bc.blocks) - 1; i >= 0; i-- {
		b := bc.blocks[i]

		if b.hash != b.calcHash() {
			return false, errors.New("invalid current hash")
		}

		if !b.isValidHash() {
			return false, errors.New("invalid nonce")
		}

		if i != 0 && b.previousHash != bc.blocks[i-1].hash {
			return false, errors.New("invalid previous hash")
		}
	}

	return true, nil
}

func (bc BlockChain) GetBalanceOfAddress(a string) money {
	var balance money
	for _, block := range bc.blocks {
		for _, transaction := range block.transactions {
			if transaction.sender == address(a) {
				balance -= transaction.amount
			}
			if transaction.recipient == address(a) {
				balance += transaction.amount
			}
		}
	}

	return balance
}

// private

func (bc BlockChain) LastBlock() block {
	return bc.blocks[len(bc.blocks)-1]
}
