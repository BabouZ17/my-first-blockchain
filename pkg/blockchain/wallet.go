package blockchain

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
)

type Wallet struct {
	PrivateKey   *rsa.PrivateKey
	PublicKey    *rsa.PublicKey
	Transactions []*Transaction
}

func NewWallet() (*Wallet, error) {
	privateKey, publicKey, err := GenerateRSAKeys()

	if err != nil {
		return nil, err
	}

	return &Wallet{
		PrivateKey:   privateKey,
		PublicKey:    publicKey,
		Transactions: []*Transaction{},
	}, err
}

func GenerateRSAKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

func (w *Wallet) AddTransaction(receiver *Wallet, amount float32) *Transaction {
	// Create transaction
	transaction := NewTransaction(w.PublicKey.N.String(), receiver.PublicKey.N.String(), amount, false)

	// Add new transaction to receiver
	receiver.Transactions = append(receiver.Transactions, transaction)

	// Add new transaction to sender
	w.Transactions = append(w.Transactions, transaction)

	return w.Transactions[len(w.Transactions)-1]
}

func (w *Wallet) SignTransaction(ts *Transaction) (string, error) {
	dataString := fmt.Sprintf("%s%s%f%t", ts.Sender, ts.Receiver, ts.Amount, ts.Coinbase)

	hashedData := sha256.Sum256([]byte(dataString))

	signedHash, err := rsa.SignPKCS1v15(
		rand.Reader,
		w.PrivateKey,
		crypto.SHA256,
		hashedData[:],
	)

	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signedHash), nil
}

func VerifyTransaction(ts *Transaction, publicKey *rsa.PublicKey, signature string) error {
	dataString := fmt.Sprintf("%s%s%f%t", ts.Sender, ts.Receiver, ts.Amount, ts.Coinbase)

	hashedData := sha256.Sum256([]byte(dataString))

	signatureBytes, err := base64.StdEncoding.DecodeString(signature)

	if err != nil {
		return err
	}

	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashedData[:], signatureBytes)

	if err != nil {
		return errors.New("the transaction signature not valid")
	}
	return nil
}
