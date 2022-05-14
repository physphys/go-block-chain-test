package bchain

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type (
	block struct {
		timestamp    time.Time
		transactions []*transaction
		previousHash [32]byte
		hash         [32]byte
		nonce        int
	}
)

func NewBlock(timestamp time.Time, transactions []*transaction, previousHash [32]byte) *block {
	b := block{
		timestamp:    timestamp,
		transactions: transactions,
		previousHash: previousHash,
		nonce:        0,
	}
	b.hash = b.calcHash()

	return &b
}

func (b block) String() string {
	return fmt.Sprintf("%s", b.transactions)
}

// private

func (b *block) mining() {
	for i := 0; !b.isValidHash(); i++ {
		b.nonce++
		b.hash = b.calcHash()
	}
	fmt.Println("Succeeded to mining")
}

func (b block) calcHash() [32]byte {
	s := fmt.Sprintf(
		"%s:%s:%s:%s",
		b.previousHash,
		b.timestamp,
		b.transactions,
		string(rune(b.nonce)),
	)
	byteSha256 := sha256.Sum256([]byte(s))

	return byteSha256
}

func (b block) isValidHash() bool {
	strSha256 := fmt.Sprintf("%x", b.hash)
	return string([]rune(strSha256)[:2]) == "00"
}
