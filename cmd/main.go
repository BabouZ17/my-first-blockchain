package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/BabouZ17/my-first-blockchain/pkg/blockchain"
)

func main() {
	// Create a new blockchain
	blockChain := blockchain.InitBlockChain()

	// Create users's wallets
	walletOfSponge, err := blockchain.NewWallet()
	if err != nil {
		log.Fatal("Could not create Sponge's wallet")
	}
	walletOfBob, err := blockchain.NewWallet()
	if err != nil {
		log.Fatal("Could not create Bob's wallet")
	}

	// Create a transaction from Sponge to Bob
	transaction := walletOfSponge.AddTransaction(walletOfBob, 10.0)

	// Sign the transaction
	signedSignature, err := walletOfSponge.SignTransaction(transaction)
	if err != nil {
		log.Fatal("Could not sign the transaction")
	}

	// Verify signature
	err = blockchain.VerifyTransaction(transaction, walletOfSponge.PublicKey, signedSignature)
	if err != nil {
		log.Fatal("Signature is not valid")
	}

	// Add block to the chain
	blockChain.AddBlock("Sponge sent 10$ to Bob", "Sponge", []*blockchain.Transaction{transaction})

	// Add new transaction
	transactionBis := walletOfBob.AddTransaction(walletOfSponge, 25.0)

	// Sign the transaction
	signedSignature, err = walletOfBob.SignTransaction(transactionBis)
	if err != nil {
		log.Fatal("Could not sign the transactionBis")
	}

	// Verify signature
	err = blockchain.VerifyTransaction(transactionBis, walletOfBob.PublicKey, signedSignature)
	if err != nil {
		log.Fatal("Signature is not valid")
	}

	// Add block to the chain
	blockChain.AddBlock("Bob sent 25$ to Sponge", "Bob", []*blockchain.Transaction{transactionBis})

	for _, block := range blockChain.Blocks {
		fmt.Printf("Prevhash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("Is valid pow: %s", strconv.FormatBool(pow.Validate()))

		fmt.Println()

		fmt.Println("Transactions: ")

		for _, ts := range block.Transactions {
			fmt.Printf("Sender: %s\n", ts.Sender)
			fmt.Printf("Receiver: %s\n", ts.Receiver)
			fmt.Printf("Amount: %f\n", ts.Amount)
			fmt.Printf("Coinbase: %t\n", ts.Coinbase)
			fmt.Println()
		}
	}

	fmt.Println("Transactions in wallet of Bob:")
	for _, ts := range walletOfBob.Transactions {
		fmt.Printf("Sender: %s\n", ts.Sender)
		fmt.Printf("Receiver: %s\n", ts.Receiver)
		fmt.Printf("Amount: %f\n", ts.Amount)
		fmt.Println()
	}

	fmt.Println("Transactions in wallet of Sponge:")
	for _, ts := range walletOfSponge.Transactions {
		fmt.Printf("Sender: %s\n", ts.Sender)
		fmt.Printf("Receiver: %s\n", ts.Receiver)
		fmt.Printf("Amount: %f\n", ts.Amount)
		fmt.Println()
	}
}
