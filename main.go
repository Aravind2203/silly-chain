package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Transaction struct {
	from   string
	to     string
	amount int64
}
type Block struct {
	prevHash  string
	hash      string
	timeStamp string
	data      Transaction
}

type BlockChain struct {
	blocks []Block
}

func generateHash(keyString string) string {
	h := sha256.New()
	h.Write([]byte(keyString))
	return fmt.Sprintf("%x", h.Sum(nil))
}
func createBlock(isGenosis bool, prevBlock Block, trans Transaction) Block {
	if isGenosis {
		keyString := trans.from + trans.to + fmt.Sprint(trans.amount)
		hash := generateHash(keyString)
		return Block{
			prevHash:  "0",
			hash:      hash,
			timeStamp: time.Now().String(),
			data:      trans,
		}
	} else {
		keyString := trans.from + trans.to + fmt.Sprint(trans.amount)
		hash := generateHash(keyString)
		return Block{
			prevHash:  prevBlock.hash,
			hash:      hash,
			timeStamp: time.Now().String(),
			data:      trans,
		}
	}
}
func (b *BlockChain) addBlock(trans Transaction) {
	if len(b.blocks) == 0 {
		block := createBlock(true, Block{}, trans)
		b.blocks = append(b.blocks, block)
	} else {
		block := createBlock(false, b.blocks[len(b.blocks)-1], trans)
		b.blocks = append(b.blocks, block)
	}
}
func (b *BlockChain) displayBlocks() {
	if len(b.blocks) == 0 {
		fmt.Println("There are no blocks init...")
		return
	}
	for i := 0; i < len(b.blocks); i++ {
		fmt.Println("----------")
		fmt.Printf("Prev Hash: %s\n", b.blocks[i].prevHash)
		fmt.Printf("Hash: %s\n", b.blocks[i].hash)
		fmt.Printf("TimeStamp:%s\n", b.blocks[i].timeStamp)
		fmt.Printf("The Transactin: %v\n", b.blocks[i].data)

	}
}
func addTransaction(from string, to string, amount int64) {
	blockChain.addBlock(Transaction{
		from:   from,
		to:     to,
		amount: amount,
	})
}

var blockChain BlockChain

func main() {
	input := bufio.NewScanner(os.Stdin)
	blockChain = BlockChain{}
	for {
		fmt.Print("Welcome To simulation Blockchain:\n1)Add Transaction\n2)DisplayBlocks\n3)Exit\nEnter your choice:")
		input.Scan()
		choice, err := strconv.ParseInt(input.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		if choice == 1 {
			fmt.Println("Enter From:")
			input.Scan()
			from := input.Text()
			fmt.Println("Enter To:")
			input.Scan()
			to := input.Text()
			fmt.Println("Enter amount")
			input.Scan()
			amount, _ := strconv.ParseInt(input.Text(), 10, 64)
			addTransaction(from, to, amount)
		} else if choice == 2 {
			blockChain.displayBlocks()
		} else if choice == 3 {
			break
		} else {
			fmt.Println("Invalid Choice")
		}

	}
}
