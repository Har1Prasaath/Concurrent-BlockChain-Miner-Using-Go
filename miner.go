package main

import (
	"fmt"
	"sync"
	"time"
)

// Miner mines a new block
func Miner(blockchain *Blockchain, transactions []Transaction, difficulty int, wg *sync.WaitGroup, resultChan chan Block) {
	defer wg.Done()

	lastBlock := blockchain.Chain[len(blockchain.Chain)-1]
	newBlock := NewBlock(lastBlock.Index+1, transactions, lastBlock.Hash)

	for {
		newBlock.Nonce++
		newBlock.Hash = newBlock.CalculateHash()
		if newBlock.IsValidHash(difficulty) {
			fmt.Printf("Miner found a valid block with nonce: %d\n", newBlock.Nonce)
			resultChan <- newBlock
			return
		}
	}
}

// StartMining starts multiple miners concurrently
func StartMining(blockchain *Blockchain, transactions []Transaction, difficulty int, numMiners int) Block {
	var wg sync.WaitGroup
	resultChan := make(chan Block, numMiners)
	timeout := time.After(10 * time.Second) // 10-second timeout

	for i := 0; i < numMiners; i++ {
		wg.Add(1)
		go Miner(blockchain, transactions, difficulty, &wg, resultChan)
	}

	// Wait for the first miner to find a valid block or timeout
	select {
	case validBlock := <-resultChan:
		wg.Wait()
		return validBlock
	case <-timeout:
		fmt.Println("Mining timed out!")
		wg.Wait()
		return Block{} // Return an empty block to indicate timeout
	}
}
