package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	data         map[string]interface{}
	hash         string //starts with minimum n zeros (n=BlockChain.difficulty)
	previousHash string
	timeStamp    time.Time
	pow          int
}

type BlockChain struct {
	genesisBlock Block
	chain        []Block
	difficulty   int
}

func (b Block) CalculateHash() string {
	data, _ := json.Marshal(b.data)
	blockData := b.previousHash + string(data) + b.timeStamp.String() + strconv.Itoa(b.pow)
	blockHash := sha256.Sum256([]byte(blockData))
	blockHashSlice := blockHash[:]
	return hex.EncodeToString(blockHashSlice)
}

func (b *Block) Mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.CalculateHash()
	}
}

func CreateBlockChain(difficulty int) BlockChain {
	genesisBlock := Block{
		hash:      "0",
		timeStamp: time.Now(),
	}

	return BlockChain{
		genesisBlock,
		[]Block{genesisBlock},
		difficulty,
	}
}

func (b *BlockChain) addBlock(from, to string, amount float64) {
	blockData := map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": amount,
	}

	lastBlock := b.chain[len(b.chain)-1]
	newBlock := Block{
		data:         blockData,
		previousHash: lastBlock.hash,
		timeStamp:    time.Now(),
	}

	newBlock.Mine(b.difficulty)
	b.chain = append(b.chain, newBlock)
}

func (b BlockChain) isValid() bool {
	previousBlock := Block{}
	currentBlock := Block{}

	for i := range b.chain[1:] {
		previousBlock = b.chain[i]
		currentBlock = b.chain[i+1]
		if currentBlock.hash != currentBlock.CalculateHash() || currentBlock.previousHash != previousBlock.hash {
			return false
		}
	}
	return true
}

func main() {
	//example main run

	myBC := CreateBlockChain(3)

	myBC.addBlock("A1", "A2", 4)
	myBC.addBlock("A2", "A3", 2)
	myBC.addBlock("A3", "A1", 5)

	fmt.Println(myBC.isValid())
}
