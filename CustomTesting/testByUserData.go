/* @createdby: Syed Muhammad Ammar
 * @StudentId: 20i2417
 * @Assignment: 01
 */

package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"

	MidLevelBlockchain "github.com/Ammar123890/Mid-Level-Blockchain"
)

func main() {
	blockchain := MidLevelBlockchain.Blockchain{}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nBlockchain Menu:")
		fmt.Println("1. Add Multiple Transactions to a Block")
		fmt.Println("2. Display Blocks")
		fmt.Println("3. Change a Block's Transaction")
		fmt.Println("4. Verify Blockchain")
		fmt.Println("5. Set number of transactions per block")
		fmt.Println("6. Set the max and mix hash")
		fmt.Println("7. Exit")
		fmt.Print("Enter your choice: ")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid choice.")
			continue
		}

		switch choice {
		case 1:
			addBlockWithMultipleTransactions(&blockchain, reader)
		case 2:
			displayBlocks(&blockchain)
		case 3:
			changeBlock(&blockchain, reader)
		case 4:
			verifyBlockchain(&blockchain)
		case 5:
			setNumberOfTransactionsPerBlock(&blockchain, reader)
		case 6:
			setBlockHashRange(&blockchain, reader)
		case 7:
			fmt.Println("Exiting the blockchain application.")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}

func addBlockWithMultipleTransactions(bc *MidLevelBlockchain.Blockchain, reader *bufio.Reader) {
	fmt.Println("Enter multiple transactions (comma-separated): ")
	transactionsStr, _ := reader.ReadString('\n')
	transactionsStr = strings.TrimSpace(transactionsStr)

	if len(transactionsStr) == 0 {
		fmt.Println("Transactions cannot be empty.")
		return
	}

	transactions := strings.Split(transactionsStr, ",")

	for i, transaction := range transactions {
		transactions[i] = strings.TrimSpace(transaction)
	}
	encryptionKey := []byte("12345678901234567890123456789012") // AES-256

	for i, transaction := range transactions {
		encryptedTx, err := MidLevelBlockchain.Encrypt([]byte(transaction), encryptionKey)
		if err != nil {
			fmt.Println("Error encrypting transaction:", err)
			return
		}
		transactions[i] = hex.EncodeToString(encryptedTx)
	}

	previousHash := ""
	if len(bc.Blocks) > 0 {
		previousHash = bc.Blocks[len(bc.Blocks)-1].CurrentHash
	}

	bc.MineBlock(transactions, previousHash)

}

func displayBlocks(bc *MidLevelBlockchain.Blockchain) {
	fmt.Println("\nBlocks in the blockchain:")
	bc.DisplayBlocks()
}

func changeBlock(bc *MidLevelBlockchain.Blockchain, reader *bufio.Reader) {
	if len(bc.Blocks) == 0 {
		fmt.Println("No blocks to change.")
		return
	}

	fmt.Print("Enter the index of the block to change: ")
	indexStr, _ := reader.ReadString('\n')
	indexStr = strings.TrimSpace(indexStr)
	index, err := strconv.Atoi(indexStr)
	if err != nil || index < 0 || index >= len(bc.Blocks) {
		fmt.Println("Invalid block index.")
		return
	}

	fmt.Print("Enter the new transaction: ")
	newTransaction, _ := reader.ReadString('\n')
	newTransaction = strings.TrimSpace(newTransaction)

	if len(newTransaction) == 0 {
		fmt.Println("Transaction cannot be empty.")
		return
	}

	encryptionKey := []byte("12345678901234567890123456789012")

	bc.ChangeBlock(index, newTransaction, encryptionKey)
	fmt.Println("Block updated successfully.")
}

func verifyBlockchain(bc *MidLevelBlockchain.Blockchain) {
	isValid := bc.VerifyChain()
	if isValid {
		fmt.Println("Blockchain is valid.")
	} else {
		fmt.Println("Blockchain is invalid.")
	}
}

func setNumberOfTransactionsPerBlock(bc *MidLevelBlockchain.Blockchain, reader *bufio.Reader) {
	fmt.Print("Enter number of transactions per block: ")
	numStr, _ := reader.ReadString('\n')
	numStr = strings.TrimSpace(numStr)
	num, err := strconv.Atoi(numStr)
	if err != nil || num <= 0 {
		fmt.Println("Invalid number. Please enter a positive integer.")
		return
	}

	bc.SetNumberOfTransactionsPerBlock(num)
	fmt.Printf("Number of transactions per block set to %d.\n", num)
}

func setBlockHashRange(bc *MidLevelBlockchain.Blockchain, reader *bufio.Reader) {
	fmt.Print("Enter minimum acceptable hash value (default: 0000000): ")
	minStr, _ := reader.ReadString('\n')
	minStr = strings.TrimSpace(minStr)

	fmt.Print("Enter maximum acceptable hash value (default: fffffff): ")
	maxStr, _ := reader.ReadString('\n')
	maxStr = strings.TrimSpace(maxStr)

	bc.SetBlockHashRangeForBlockCreation(minStr, maxStr)
	fmt.Printf("Block hash range set from %s to %s.\n", minStr, maxStr)
}
