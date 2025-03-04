package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

// Block represents a block in the blockchain
type Block struct {
	Index        int           `json:"index"`
	Timestamp    string        `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	PrevHash     string        `json:"prevHash"`
	Nonce        int           `json:"nonce"`
	Hash         string        `json:"hash"`
}

// Blockchain represents the blockchain
type Blockchain struct {
	Chain []Block `json:"chain"`
}

// NewBlock creates a new block
func NewBlock(index int, transactions []Transaction, prevHash string) Block {
	block := Block{
		Index:        index,
		Timestamp:    time.Now().String(),
		Transactions: transactions,
		PrevHash:     prevHash,
		Nonce:        0,
	}
	block.Hash = block.CalculateHash()
	return block
}

// CalculateHash computes the hash of a block
func (b Block) CalculateHash() string {
	blockData, _ := json.Marshal(b)
	hash := sha256.Sum256(blockData)
	return hex.EncodeToString(hash[:])
}

// IsValidHash checks if the hash meets the difficulty requirement
func (b Block) IsValidHash(difficulty int) bool {
	prefix := string(b.Hash[:difficulty])
	for i := 0; i < difficulty; i++ {
		if prefix[i] != '0' {
			return false
		}
	}
	return true
}

// NewBlockchain creates a new blockchain with a genesis block
func NewBlockchain() Blockchain {
	genesisBlock := NewBlock(0, []Transaction{}, "0")
	return Blockchain{Chain: []Block{genesisBlock}}
}

// DisplayBlockchain prints the blockchain
func (bc Blockchain) DisplayBlockchain() {
	for _, block := range bc.Chain {
		fmt.Printf("Index: %d, Hash: %s, PrevHash: %s, Nonce: %d\n", block.Index, block.Hash, block.PrevHash, block.Nonce)
	}
}
