/* @createdby: Syed Muhammad Ammar
 * @StudentId: 20i2417
 * @Assignment: 01
 */

package MidLevelBlockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

// Block represents a single block in the blockchain.
type Block struct {
	Transactions []string
	Nonce        int
	PreviousHash string
	CurrentHash  string
	MerkleRoot   string // New field for Merkle root
}

// NewBlock creates a new block.
func NewBlock(transactions []string, nonce int, previousHash string) *Block {
	block := &Block{
		Nonce:        nonce,
		PreviousHash: previousHash,
		Transactions: transactions, // This line was missing
	}

	// Convert transactions to [][]byte
	var txData [][]byte
	for _, tx := range transactions {
		txData = append(txData, []byte(tx))
	}

	// Create a new Merkle tree and set the Merkle root
	tree := NewMerkleTree(txData)
	block.MerkleRoot = hex.EncodeToString(tree.RootNode.Data)

	// Calculate the hash for the new block
	block.CurrentHash = block.CalculateHash()

	return block
}

// CalculateHash calculates the hash of a block.
func (b *Block) CalculateHash() string {
	transactionData := strings.Join(b.Transactions, "")
	data := fmt.Sprintf("%s%d%s%s", transactionData, b.Nonce, b.PreviousHash, b.MerkleRoot)
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}
