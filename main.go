package main

import (
	"fmt"
)

func main() {
	// Initialize blockchain
	blockchain := NewBlockchain()

	// Define difficulty (number of leading zeros required in the hash)
	difficulty := 4

	// Create wallets for users (Re-added the wallet creation)
	aliceWallet := NewWallet()
	bobWallet := NewWallet()

	// Create some transactions
	transactions := []Transaction{
		NewTransaction("Alice", "Bob", 10),
		NewTransaction("Bob", "Charlie", 5),
	}

	// Example usage of wallets to sign transactions
	signature1 := aliceWallet.SignTransaction(transactions[0])
	signature2 := bobWallet.SignTransaction(transactions[1])

	fmt.Println("Signature 1:", signature1)
	fmt.Println("Signature 2:", signature2)

	// Start mining with 4 concurrent miners
	fmt.Println("Starting mining...")
	validBlock := StartMining(&blockchain, transactions, difficulty, 4)

	// Add the mined block to the blockchain
	blockchain.Chain = append(blockchain.Chain, validBlock)
	fmt.Println("Block mined and added to the blockchain!")

	// Display the blockchain
	fmt.Println("Blockchain:")
	blockchain.DisplayBlockchain()

	// Simulate broadcasting the block to the network
	network := NewNetwork()
	network.BroadcastBlock(validBlock)
}
