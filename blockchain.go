/* @createdby: Syed Muhammad Ammar
 * @StudentId: 20i2417
 * @Assignment: 01
 */

package MidLevelBlockchain

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

// Blockchain represents a blockchain.
type Blockchain struct {
	Blocks []*Block
}

//Mine Block

const defaultDifficulty = 2

// MineBlock mines a new block for the given transaction and previous hash.
// The PoW mechanism will ensure that the hash of the block starts with a certain number of zeros.
func (bc *Blockchain) MineBlock(transactions []string, previousHash string) {
	nonce := 0

	// Check if enough transactions are available to mine a block
	if len(transactions) < minTransactionsPerBlock {
		fmt.Println("Not enough transactions to mine a new block.")
		return
	}

	difficulty := defaultDifficulty
	if len(bc.Blocks)%5 == 0 {
		difficulty++
	}

	block := NewBlock(transactions, nonce, previousHash)

	// Adjusted Mining process using PoW
	for {
		currentHash := block.CalculateHash()

		if strings.HasPrefix(currentHash, strings.Repeat("0", difficulty)) && isValidHash(currentHash) {
			block.CurrentHash = currentHash
			break
		}

		nonce++
		block.Nonce = nonce
	}

	bc.Blocks = append(bc.Blocks, block)
	fmt.Println("Block added successfully.")
}

// Modify isValidHash to check hash against set range
func isValidHash(hash string) bool {
	return hash >= minAcceptableHashValue && hash <= maxAcceptableHashValue
}

// DisplayBlocks prints all blocks in the blockchain.
// blockchain.go

func (bc *Blockchain) DisplayBlocks() {
	encryptionKey := []byte("12345678901234567890123456789012") // Same key as used for encryption
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintln(w, "Block\tTransaction\tNonce\tPrevious Hash\tCurrent Hash")
	for i, block := range bc.Blocks {
		decryptedTransactions := []string{}

		for _, encryptedTx := range block.Transactions {
			txBytes, err := hex.DecodeString(encryptedTx)
			if err != nil {
				fmt.Println("Error decoding transaction hex:", err)
				return
			}

			decryptedTx, err := Decrypt(txBytes, encryptionKey)
			if err != nil {
				fmt.Println("Error decrypting transaction:", err)
				return
			}

			decryptedTransactions = append(decryptedTransactions, string(decryptedTx))
		}

		// Use decryptedTransactions instead of block.Transactions for display
		prevHash := limitHashDisplay(block.PreviousHash, 16)
		currHash := limitHashDisplay(block.CurrentHash, 16)
		fmt.Fprintf(w, "%d\t%s\t%d\t%s\t%s\n", i, strings.Join(decryptedTransactions, ", "), block.Nonce, prevHash, currHash)
	}

	w.Flush()
}

// limitHashDisplay limits the hash to a specified length and appends "..." if it exceeds that length.
func limitHashDisplay(hash string, maxLength int) string {
	if len(hash) > maxLength {
		return hash[:maxLength-3] + "..."
	}
	return hash
}

// ChangeBlock changes the transaction of a given block.

func (bc *Blockchain) ChangeBlock(blockIndex int, newTransaction string, encryptionKey []byte) {
	if blockIndex < 0 || blockIndex >= len(bc.Blocks) {
		fmt.Println("Invalid block index")
		return
	}

	encryptedTx, err := Encrypt([]byte(newTransaction), encryptionKey)
	if err != nil {
		fmt.Println("Error encrypting transaction:", err)
		return
	}
	hexEncodedTx := hex.EncodeToString(encryptedTx)

	for i := blockIndex; i < len(bc.Blocks); i++ {
		if i > 0 {
			bc.Blocks[i].PreviousHash = bc.Blocks[i-1].CurrentHash
		}

		if i == blockIndex {
			bc.Blocks[i].Transactions = []string{hexEncodedTx} // Replace the transactions, not append
		}

		bc.Blocks[i].CurrentHash = bc.Blocks[i].CalculateHash()
	}
}

// VerifyChain verifies the integrity of the blockchain.
func (bc *Blockchain) VerifyChain() bool {
	for i := 0; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]

		// Check block's content hash matches its current hash
		if currentBlock.CurrentHash != currentBlock.CalculateHash() {
			return false
		}

		// For all blocks except the first, check if previous hash matches
		if i > 0 {
			previousBlock := bc.Blocks[i-1]
			if currentBlock.PreviousHash != previousBlock.CurrentHash {
				return false
			}
		}

		// Check Merkle root integrity
		var txData [][]byte
		for _, tx := range currentBlock.Transactions {
			txData = append(txData, []byte(tx))
		}

		tree := NewMerkleTree(txData)
		if currentBlock.MerkleRoot != hex.EncodeToString(tree.RootNode.Data) {
			return false
		}
	}
	return true
}

// Assuming we add a variable to keep track of the number of transactions per block:
var minTransactionsPerBlock int = 2 // default

func (bc *Blockchain) SetNumberOfTransactionsPerBlock(num int) {
	minTransactionsPerBlock = num
}

// For the range of acceptable hash values (not implemented in mining, but here's a setter):
var minAcceptableHashValue string = "0000000" // default low end
var maxAcceptableHashValue string = "fffffff" // default high end

func (bc *Blockchain) SetBlockHashRangeForBlockCreation(min, max string) {
	minAcceptableHashValue = min
	maxAcceptableHashValue = max
}
