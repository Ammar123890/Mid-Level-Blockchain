package main

import (
	"bufio"
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
		fmt.Println("0. Test with Sample Data")
		fmt.Println("1. Add Multiple Transactions to a Block")
		fmt.Println("2. Display Blocks")
		fmt.Println("3. Change a Block's Transaction")
		fmt.Println("4. Verify Blockchain")
		fmt.Println("5. Tamper Block")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid choice.")
			continue
		}

		switch choice {
		case 0:
			testWithData(&blockchain)
		case 1:
			addBlockWithMultipleTransactions(&blockchain, reader) // Updated function for handling multiple transactions
		case 2:
			displayBlocks(&blockchain)
		case 3:
			changeBlock(&blockchain, reader)
		case 4:
			verifyBlockchain(&blockchain)
		case 5:
			tamperBlock(&blockchain, reader)
		case 6:
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

	previousHash := ""
	if len(bc.Blocks) > 0 {
		previousHash = bc.Blocks[len(bc.Blocks)-1].CurrentHash
	}

	// Removing nonce as the new mechanism uses PoW to determine nonce
	bc.MineBlock(transactions, previousHash) // Using the new MineBlock function
	fmt.Println("Block added successfully.")
}

// ... [Rest of the functions remain mostly the same]

func tamperBlock(bc *MidLevelBlockchain.Blockchain, reader *bufio.Reader) {
	if len(bc.Blocks) == 0 {
		fmt.Println("No blocks to tamper.")
		return
	}

	fmt.Print("Enter the index of the block to tamper: ")
	indexStr, _ := reader.ReadString('\n')
	indexStr = strings.TrimSpace(indexStr)
	index, err := strconv.Atoi(indexStr)
	if err != nil || index < 0 || index >= len(bc.Blocks) {
		fmt.Println("Invalid block index.")
		return
	}

	fmt.Print("Enter the new transaction to tamper: ")
	newTransaction, _ := reader.ReadString('\n')
	newTransaction = strings.TrimSpace(newTransaction)

	if len(newTransaction) == 0 {
		fmt.Println("Transaction cannot be empty.")
		return
	}

	bc.TamperBlock(index, newTransaction)
	fmt.Println("Block tampered successfully (hash not recalculated).")
}

// [testWithData function remains the same]
