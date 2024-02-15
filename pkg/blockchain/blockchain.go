package blockchain

type BlockChain struct {
	Blocks []*Block
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func (chain *BlockChain) AddBlock(Data string, coinbaseRcpt string, transactions []*Transaction) {
	// Add a new coinbase transaction to reward the miner
	coinbaseTransaction := &Transaction{"coinbase", coinbaseRcpt, 10.0, true}
	newTransactions := []*Transaction{coinbaseTransaction}
	newTransactions = append(newTransactions, transactions...)

	previousBlock := chain.Blocks[len(chain.Blocks)-1]

	newBlock := CreateBlock(Data, previousBlock.Hash, newTransactions)
	chain.Blocks = append(chain.Blocks, newBlock)
}
