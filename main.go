package main

import (
	"Assignment01/assignment01bca"
	"fmt"
)

func main() {

	BlocksStore := []*assignment01bca.Block{}
	block := assignment01bca.NewBlock("First Block", 1, "0")
	BlocksStore = append(BlocksStore, block)

	fmt.Println("Blockchain Application")

	for {
		fmt.Println("\nOptions:")
		fmt.Println("1. Create a new block")
		fmt.Println("2. Display blockchain")
		fmt.Println("3. Edit block transaction")
		fmt.Println("4. Verify blockchain validity")
		fmt.Println("5. Add a few blocks") // Added option
		fmt.Println("6. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("\nCreating a new block...")
			fmt.Print("Enter transaction: ")
			var trans string
			fmt.Scan(&trans)

			fmt.Print("\nEnter nonce: ")
			var nonce int
			fmt.Scan(&nonce)

			fmt.Print("Enter previous hash: ")
			var prevHash string
			fmt.Scan(&prevHash)

			block := assignment01bca.NewBlock(trans, nonce, prevHash)
			BlocksStore = append(BlocksStore, block)
			fmt.Println("Block created successfully!")

		case 2:
			fmt.Println("\nDisplaying the blockchain...")

			assignment01bca.DisplayBlocks(BlocksStore)

		case 3:
			fmt.Println("\nEditing a block transaction...")
			if len(BlocksStore) == 0 {
				fmt.Println("No blocks to edit. Create a block first.")
				continue
			}

			fmt.Print("Enter the block index to edit (0 to ", len(BlocksStore)-1, "): ")
			var index int
			fmt.Scan(&index)

			if index < 0 || index >= len(BlocksStore) {
				fmt.Println("Invalid block index.")
				continue
			}

			fmt.Print("Enter the new transaction: ")
			var newTrans string
			fmt.Scan(&newTrans)

			assignment01bca.ChangeBlock(BlocksStore[index], newTrans)
			fmt.Println("Block transaction edited successfully!")

		case 4:
			fmt.Println("\nVerifying blockchain validity...")
			isValid := assignment01bca.VerifyChain(BlocksStore)
			if isValid {
				fmt.Println("Blockchain is valid.")
			} else {
				fmt.Println("Blockchain is NOT valid. Changes detected!")
			}

		case 5:
			// Add a few blocks (in this example, we add three blocks)
			fmt.Println("\nAdding a few blocks...")
			for i := 0; i < 3; i++ {

				fmt.Printf("\nCreating Block %d...\n", i+1)
				fmt.Print("Enter transaction: ")
				var trans string
				fmt.Scan(&trans)

				// You can generate nonce and previous hash as needed.

				block := assignment01bca.NewBlock(trans, i+1, BlocksStore[len(BlocksStore)-1].Hash)
				BlocksStore = append(BlocksStore, block)
				fmt.Printf("Block %d created successfully!\n", i+1)
			}

		case 6:
			fmt.Println("\nExiting the program.")
			return

		default:
			fmt.Println("\nInvalid choice. Please select a valid option.")
		}
	}
}
