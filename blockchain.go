/* @createdby: Syed Muhammad Ammar
 * @StudentId: 20i2417
 * @Assignment: 01
 */

package MidLevelBlockchain

import (
	"fmt"
	"os"
	"text/tabwriter"
)

// Blockchain represents a blockchain.
type Blockchain struct {
	Blocks []*Block
}

// DisplayBlocks prints all blocks in the blockchain.
// DisplayBlocks prints all blocks in the blockchain.
func (bc *Blockchain) DisplayBlocks() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintln(w, "Block\tTransaction\tNonce\tPrevious Hash\tCurrent Hash")
	for i, block := range bc.Blocks {
		// Limit hash display to 16 characters and append "..." if it exceeds that length
		prevHash := limitHashDisplay(block.PreviousHash, 16)
		currHash := limitHashDisplay(block.CurrentHash, 16)

		fmt.Fprintf(w, "%d\t%s\t%d\t%s\t%s\n", i, block.Transaction, block.Nonce, prevHash, currHash)
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
			bc.Blocks[i].Transaction = newTransaction
		}

		bc.Blocks[i].CurrentHash = bc.Blocks[i].CalculateHash()
	}
}

// VerifyChain verifies the integrity of the blockchain.
func (bc *Blockchain) VerifyChain() bool {
	for i := 0; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]

		// Check if the block's content hash matches its current hash
		if currentBlock.CurrentHash != currentBlock.CalculateHash() {
			return false
		}

		// For all blocks except the first, check if the previous hash matches
		if i > 0 {
			previousBlock := bc.Blocks[i-1]
			if currentBlock.PreviousHash != previousBlock.CurrentHash {
				return false
			}
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
	block.Transaction = newTransaction
	// No hash recalculation here
}
