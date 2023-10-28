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
	var currentHash string

	difficulty := defaultDifficulty
	if len(bc.Blocks)%5 == 0 { // For every 5 blocks, increase difficulty by 1.
		difficulty++
	}

	block := &Block{
		Transactions: transactions,
		Nonce:        nonce,
		PreviousHash: previousHash,
		CurrentHash:  "", // Temporarily set to empty; will be updated below
	}

	// Calculate Merkle Root for transactions
	var txData [][]byte
	for _, tx := range transactions {
		txData = append(txData, []byte(tx))
	}
	tree := NewMerkleTree(txData)
	block.MerkleRoot = hex.EncodeToString(tree.RootNode.Data)

	// Mining process using PoW
	for {
		block.Nonce = nonce
		currentHash = block.CalculateHash()

		if isValidHash(currentHash, difficulty) {
			block.CurrentHash = currentHash
			break
		}

		nonce++
	}

	bc.Blocks = append(bc.Blocks, block)
}

// isValidHash checks if the hash has a required number of leading zeros (as defined by difficulty)
func isValidHash(hash string, difficulty int) bool {
	prefix := ""

	for i := 0; i < difficulty; i++ {
		prefix += "0"
	}

	return hash[:difficulty] == prefix
}

// DisplayBlocks prints all blocks in the blockchain.
func (bc *Blockchain) DisplayBlocks() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintln(w, "Block\tTransaction\tNonce\tPrevious Hash\tCurrent Hash")
	for i, block := range bc.Blocks {
		// Limit hash display to 16 characters and append "..." if it exceeds that length
		prevHash := limitHashDisplay(block.PreviousHash, 16)
		currHash := limitHashDisplay(block.CurrentHash, 16)

		fmt.Fprintf(w, "%d\t%s\t%d\t%s\t%s\n", i, strings.Join(block.Transactions, ", "), block.Nonce, prevHash, currHash)

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

func (bc *Blockchain) ChangeBlock(blockIndex int, newTransaction string) {
	if blockIndex < 0 || blockIndex >= len(bc.Blocks) {
		fmt.Println("Invalid block index")
		return
	}

	for i := blockIndex; i < len(bc.Blocks); i++ {
		if i > 0 {
			bc.Blocks[i].PreviousHash = bc.Blocks[i-1].CurrentHash
		}

		if i == blockIndex {
			bc.Blocks[i].Transactions = append(bc.Blocks[i].Transactions, newTransaction)
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

// TamperBlock tampers with the transaction of a given block but doesn't update its hash.
func (bc *Blockchain) TamperBlock(blockIndex int, newTransaction string) {
	if blockIndex < 0 || blockIndex >= len(bc.Blocks) {
		fmt.Println("Invalid block index")
		return
	}

	block := bc.Blocks[blockIndex]
	block.Transactions = append(block.Transactions, newTransaction)

	// No hash recalculation here
}

// Assuming we add a variable to keep track of the number of transactions per block:
var maxTransactionsPerBlock int = 5 // default

func (bc *Blockchain) setNumberOfTransactionsPerBlock(num int) {
	maxTransactionsPerBlock = num
}

// For the range of acceptable hash values (not implemented in mining, but here's a setter):
var minAcceptableHashValue string = "0000000" // default low end
var maxAcceptableHashValue string = "fffffff" // default high end

func (bc *Blockchain) setBlockHashRangeForBlockCreation(min, max string) {
	minAcceptableHashValue = min
	maxAcceptableHashValue = max
}
