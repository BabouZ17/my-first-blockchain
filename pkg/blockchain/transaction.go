package blockchain

type Transaction struct {
	Sender   string
	Receiver string
	Amount   float32
	Coinbase bool
}

func NewTransaction(sender, receiver string, amount float32, coinbase bool) *Transaction {
	return &Transaction{Sender: sender, Receiver: receiver, Amount: amount, Coinbase: coinbase}
}
