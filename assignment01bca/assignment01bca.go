// Repo-Link : "https://github.com/UM00/assignment01bca.git"
package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	Trans    string
	Nonce    int
	PrevHash string
	Hash     string
}

func HashGenerator(B Block) string {
	input := fmt.Sprintf("%s%d%s", B.Trans, B.Nonce, B.PrevHash)
	Hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(Hash[:])
}

func CalculateHash(strngTohash string) string {
	Hash := sha256.Sum256([]byte(strngTohash))
	return hex.EncodeToString(Hash[:])
}

func NewBlock(trans string, nonce int, prevHash string) *Block {
	block := Block{
		Trans:    trans,
		Nonce:    nonce,
		PrevHash: prevHash,
	}
	block.Hash = HashGenerator(block)
	return &block
}

func ChangeBlock(block *Block, newTrans string) {
	block.Trans = newTrans
	block.Hash = HashGenerator(*block)
}

func VerifyChain(blocks []*Block) bool {
	for i := 1; i < len(blocks); i++ {
		currentBlock := blocks[i]
		previousBlock := blocks[i-1]

		if currentBlock.PrevHash != HashGenerator(*previousBlock) {
			return false
		}

		if currentBlock.Hash != HashGenerator(*currentBlock) {
			return false
		}
	}

	return true
}

func DisplayBlocks(blocks []*Block) {
	fmt.Println()
	fmt.Println("Blockchain:")
	fmt.Println("--B--L--O--C--K--S--")
	for _, block := range blocks {
		fmt.Println("╔════════════════════════════╗")
		fmt.Printf("║ Transaction: %s\n", block.Trans)
		fmt.Printf("║ Nonce: %d\n", block.Nonce)
		fmt.Printf("║ Previous Hash: %s\n", block.PrevHash)
		fmt.Printf("║ Block Hash: %s\n", block.Hash)
		fmt.Println("╚════════════════════════════╝")
	}
}
