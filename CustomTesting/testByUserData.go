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
		fmt.Println("0. Test with Sample Data") // New option
		fmt.Println("1. Add a Block")
		fmt.Println("2. Display Blocks")
		fmt.Println("3. Change a Block")
		fmt.Println("4. Verify Blockchain")
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
			testWithData(&blockchain) // Call the new function for testing with sample data
		case 1:
			addBlock(&blockchain, reader)
		case 2:
			displayBlocks(&blockchain)
		case 3:
			changeBlock(&blockchain, reader)
		case 4:
			verifyBlockchain(&blockchain)
		case 6:
			fmt.Println("Exiting the blockchain application.")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}

func addBlock(bc *MidLevelBlockchain.Blockchain, reader *bufio.Reader) {
	fmt.Print("Enter transaction: ")
	transaction, _ := reader.ReadString('\n')
	transaction = strings.TrimSpace(transaction)

	if len(transaction) == 0 {
		fmt.Println("Transaction cannot be empty.")
		return
	}

	fmt.Print("Enter nonce (an integer): ")
	nonceStr, _ := reader.ReadString('\n')
	nonceStr = strings.TrimSpace(nonceStr)
	nonce, err := strconv.Atoi(nonceStr)
	if err != nil {
		fmt.Println("Invalid input for nonce.")
		return
	}

	previousHash := ""
	if len(bc.Blocks) > 0 {
		previousHash = bc.Blocks[len(bc.Blocks)-1].CurrentHash
	}

	block := MidLevelBlockchain.NewBlock(transaction, nonce, previousHash)
	bc.Blocks = append(bc.Blocks, block)
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

	bc.ChangeBlock(index, newTransaction) // Fixed this line
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

func testWithData(bc *MidLevelBlockchain.Blockchain) {
	block1 := MidLevelBlockchain.NewBlock("Alice to Bob", 123, "")
	block2 := MidLevelBlockchain.NewBlock("Bob to Charlie", 456, block1.CurrentHash)

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
