package blockchain

import (
	"math/rand"
)

type Block struct {
	Hash         string
	Data         string
	PrevHash     string
	Nonce        int
	Transactions []*Transaction
}

func CreateBlock(data, prevHash string, transactions []*Transaction) *Block {
	r := rand.New(rand.NewSource(99))
	initialNonce := r.Int()

	block := &Block{"", data, prevHash, initialNonce, transactions}
	pow := NewProofOfWork(block)

	var hash []byte
	nonce, hash := pow.MineBlock()

	block.Nonce = nonce
	block.Hash = string(hash[:])

	return block
}

func Genesis() *Block {
	coinbaseTransaction := &Transaction{"Coinbase", "Genesis", 0.0, true}
	return CreateBlock("Genesis", "", []*Transaction{coinbaseTransaction})
}
