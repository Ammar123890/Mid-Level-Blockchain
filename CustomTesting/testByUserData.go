/* @createdby: Syed Muhammad Ammar
 * @StudentId: 20i2417
 * @Assignment: 01
 */

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
		fmt.Println("6. Set number of transactions per block")
		fmt.Println("7. Set the max and mix hash")
		fmt.Println("8. Exit")
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
			addBlockWithMultipleTransactions(&blockchain, reader)
		case 2:
			displayBlocks(&blockchain)
		case 3:
			changeBlock(&blockchain, reader)
		case 4:
			verifyBlockchain(&blockchain)
		case 5:
			tamperBlock(&blockchain, reader)
		case 6:
			setNumberOfTransactionsPerBlock(&blockchain, reader)
		case 7:
			setBlockHashRange(&blockchain, reader)
		case 8:
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

	bc.MineBlock(transactions, previousHash)
	fmt.Println("Block added successfully.")
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

	bc.ChangeBlock(index, newTransaction)
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

func testWithData(bc *MidLevelBlockchain.Blockchain) {
	block1 := MidLevelBlockchain.NewBlock([]string{"Alice to Bob"}, 123, "")
	block2 := MidLevelBlockchain.NewBlock([]string{"Bob to Charlie"}, 456, block1.CurrentHash)

	bc.Blocks = append(bc.Blocks, block1)
	bc.Blocks = append(bc.Blocks, block2)

	fmt.Println("Sample Data Added to Blockchain:")
	bc.DisplayBlocks()

	index := len(bc.Blocks) - 1
	bc.ChangeBlock(index, "Mallory to Eve")

	fmt.Println("\nBlock 2 Transaction Changed:")
	bc.DisplayBlocks()

	isValid := bc.VerifyChain()
	if isValid {
		fmt.Println("\nBlockchain is valid.")
	} else {
		fmt.Println("\nBlockchain is invalid.")
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
