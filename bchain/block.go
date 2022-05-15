package bchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"time"
)

type (
	blocks    []block
	blockHash [32]byte

	block struct {
		committedAt           time.Time
		committedTransactions transactions
		previousHash          blockHash
		hash                  blockHash
		nonce                 int
	}
)

func NewBlock(timestamp time.Time, transactions []transaction, previousHash [32]byte) block {
	b := block{
		committedAt:           timestamp,
		committedTransactions: transactions,
		previousHash:          previousHash,
		nonce:                 0,
	}
	b.hash = b.calcHash()

	return b
}

func (b block) String() string {
	return fmt.Sprintf("%s", b.committedTransactions)
}

// private

func (b *block) mining() {
	for {
		b.nonce++
		b.hash = b.calcHash()

		if err := b.validateHash(); err == nil {
			fmt.Println("Succeeded to mining")
			return
		}
	}
}

func (b block) calcHash() [32]byte {
	s := fmt.Sprintf(
		"%s:%s:%s:%s",
		b.previousHash,
		b.committedAt,
		b.committedTransactions,
		string(rune(b.nonce)),
	)
	byteSha256 := sha256.Sum256([]byte(s))

	return byteSha256
}

func (b block) validateHash() error {
	if b.hash != b.calcHash() {
		return errors.New("current hash does not match with calculated")
	}

	strSha256 := fmt.Sprintf("%x", b.hash)
	firstTwoCharacter := string([]rune(strSha256)[:2])

	if firstTwoCharacter != "00" {
		return errors.New("invalid hash")
	}

	return nil
}

func (bs blocks) last() block {
	return bs[len(bs)-1]
}

func (bs blocks) Validate() error {
	for i, b := range bs {
		if err := b.validateHash(); err != nil {
			return err
		}

		if i == 0 {
			continue
		}

		if b.previousHash != bs[i-1].hash {
			fmt.Println(len(bs))
			fmt.Println(fmt.Sprintf("%x", bs[0].hash))
			fmt.Println(fmt.Sprintf("%x", bs[0].previousHash))
			fmt.Println(fmt.Sprintf("%x", bs[1].hash))
			fmt.Println(fmt.Sprintf("%x", bs[1].previousHash))

			return errors.New("invalid previous hash")
		}
	}

	return nil
}
